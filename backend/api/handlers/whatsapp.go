package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/cache"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/shared/objects"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"github.com/xuri/excelize/v2"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WhatsappHandler struct {
	erpContext                  *context.ERPContext
	waService                   *whatsmeow_client.WhatsmeowService
	appService                  *app.AppService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	geminiService               *google.GeminiService
	pmService                   *project_management.ProjectManagementService
	contactService              *contact.ContactService
	cacheService                *cache.CacheManager[paginate.Page]
}

// var eligibleKeyWords = []string{"Order", "order", "ORDER", "Orders", "orders", "ORDERS", "LOGIN", "login", "Login", "Menu", "MENU", "menu", "logout"}

func NewWhatsappHandler(erpContext *context.ERPContext) *WhatsappHandler {
	var waService *whatsmeow_client.WhatsmeowService
	waSrv, ok := erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
	if ok {
		waService = waSrv
	}

	var appService *app.AppService
	appSrv, ok := erpContext.AppService.(*app.AppService)
	if ok {
		appService = appSrv
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := erpContext.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	geminiService, ok := erpContext.ThirdPartyServices["GEMINI"].(*google.GeminiService)
	if !ok {
		panic("GeminiService is not found")
	}

	pmService, ok := erpContext.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}

	contactService, ok := erpContext.ContactService.(*contact.ContactService)
	if !ok {
		panic("ContactService is not instance of contact.ContactService")
	}

	cacheService, ok := erpContext.ThirdPartyServices["Cache"].(*cache.CacheManager[paginate.Page])
	if !ok {
		panic("CacheService is not instance of cache.CacheManager")
	}
	return &WhatsappHandler{
		erpContext:                  erpContext,
		waService:                   waService,
		appService:                  appService,
		customerRelationshipService: customerRelationshipService,
		geminiService:               geminiService,
		pmService:                   pmService,
		contactService:              contactService,
		cacheService:                cacheService,
	}
}

func parseTemplateID(msg string) *string {
	// Get UUID from string
	uuidRe := regexp.MustCompile(`@@\[([^\]]+)\]\(([^)]+)\)`)
	match := uuidRe.FindStringSubmatch(msg)
	if len(match) > 2 {
		msg = strings.ReplaceAll(msg, match[0], match[2])
		fmt.Println("MATCHED UUID", match[2])
		fmt.Println("MATCHED Msg", msg)
		return &msg
	}
	return nil

}

func parseMsgTemplate(contact mdl.ContactModel, member *models.MemberModel, msg string) string {
	re := regexp.MustCompile(`@\[([^\]]+)\]|\(\{\{([^}]+)\}\}\)`)

	// Replace
	result := re.ReplaceAllStringFunc(msg, func(s string) string {
		matches := re.FindStringSubmatch(s)
		re2 := regexp.MustCompile(`@\[([^\]]+)\]`)
		if re2.MatchString(s) {
			return ""
		}

		if matches[0] == "({{user}})" {
			return "*" + contact.Name + "*"
		}
		if matches[0] == "({{phone}})" {
			return "*" + *contact.Phone + "*"
		}
		if matches[0] == "({{agent}})" && member != nil {
			return "*" + member.User.FullName + "*"
		}

		if matches[0] == "({{product}})" {
			var customData map[string]string
			json.Unmarshal([]byte(contact.CustomData), &customData)
			return "*" + customData["product"] + "*"
		}
		return s // Kalau tidak ada datanya, biarkan
	})

	return result
}

func (h *WhatsappHandler) ReadAllMessage(c *gin.Context) {
	sessionId := c.Params.ByName("session_id")
	var session *models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	err = h.customerRelationshipService.WhatsappService.ReadAllMessages(session.Session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Read all message success"})
}

func (h *WhatsappHandler) SendMessage(c *gin.Context) {
	var input struct {
		Message   string                       `json:"message"`
		Files     []models.FileModel           `json:"files"`
		Products  []models.ProductModel        `json:"products"`
		IsCaption bool                         `json:"is_caption"`
		RefMsg    *models.WhatsappMessageModel `json:"ref_msg"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message is required"})
		return
	}
	sessionId := c.Params.ByName("session_id")
	var session *models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	splitJID := strings.Split(session.JID, "@")
	splitSep := strings.Split(splitJID[0], ":")
	var isGroup bool
	var ok bool

	info := make(map[string]any)
	var infoStr string = "{}"
	var responseDuration *float64
	var lastCustMsg models.WhatsappMessageModel
	var refID *string
	err = h.customerRelationshipService.WhatsappService.GetWhatsappLastCustomerMessages(session.JID, session.Session, &lastCustMsg)
	if err == nil {
		json.Unmarshal([]byte(lastCustMsg.Info), &info)

		info["Timestamp"] = time.Now().Format("2006-01-02T15:04:05-07:00")

		b, _ := json.Marshal(info)
		infoStr = string(b)

		isGroup, ok = lastCustMsg.MessageInfo["IsGroup"].(bool)
		if ok {
			if lastCustMsg.IsFromMe {
				responseDuration = lastCustMsg.ResponseTime
				refID = lastCustMsg.RefID
			} else {
				var dur time.Duration = time.Since(*lastCustMsg.CreatedAt)
				duration := dur.Seconds()
				responseDuration = &duration
				lastCustMsg.IsReplied = true
				h.erpContext.DB.Save(&lastCustMsg)
				refID = &lastCustMsg.ID

			}
		}

		// fmt.Println("ERROR 2", dur)
	}

	// fmt.Println("LAST MESSAGE")
	// utils.LogJson(lastCustMsg)

	userID := c.MustGet("userID").(string)
	memberID := c.MustGet("memberID").(string)
	member := c.MustGet("member").(models.MemberModel)

	parsedMessage := parseMsgTemplate(*session.Contact, &member, input.Message)
	now := time.Now()
	templateID := parseTemplateID(input.Message)
	var connection connection.ConnectionModel
	err = h.erpContext.DB.Select("id, session, company_id").First(&connection, "session ilike ? and company_id = ?", splitSep[0]+"%", session.CompanyID).Error
	if err == nil {
		fmt.Println("MESSAGE CONNECTION", session.JID)
		fmt.Println("ACTIVE CONNECTION", connection.Session)
		if session.JID != connection.Session {
			// UPDATE SESSION JID
			session.JID = connection.Session
			err = h.erpContext.DB.Save(&session).Error
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}
	}

	// thumbnail, attachments := getThumbnail(input.Files)
	// var mediaURL, mimeType string
	// if thumbnail != nil {
	// 	mediaURL = thumbnail.URL
	// 	mimeType = thumbnail.MimeType
	// }

	var waDataReply models.WhatsappMessageModel = models.WhatsappMessageModel{
		Sender:   splitSep[0],
		Receiver: *session.Contact.Phone,
		Message:  parsedMessage,

		MessageInfo:  info,
		Info:         infoStr,
		Session:      session.Session,
		JID:          session.JID,
		IsFromMe:     true,
		IsRead:       false,
		IsGroup:      isGroup,
		ContactID:    session.ContactID,
		CompanyID:    session.CompanyID,
		ResponseTime: responseDuration,
		UserID:       &userID,
		MemberID:     &memberID,
		RefID:        refID,
	}

	if input.RefMsg != nil {
		waDataReply.QuotedMessageID = &input.RefMsg.ID
		waDataReply.QuotedMessage = &input.RefMsg.Message
	}
	to := waDataReply.Receiver
	if waDataReply.IsGroup {
		to = waDataReply.Session
	}
	fmt.Println("WA DATA", fmt.Sprintf("%s@%s", splitSep[0], splitJID[1]))
	utils.LogJson(waDataReply)
	if templateID == nil {
		h.customerRelationshipService.WhatsappService.SetMsgData(h.waService, &waDataReply, to, input.Files, input.Products, true, input.RefMsg)
		resp, err := customer_relationship.SendCustomerServiceMessage(h.customerRelationshipService.WhatsappService)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		msgNotif := gin.H{
			"message":    input.Message,
			"command":    "WHATSAPP_RECEIVED",
			"session_id": session.ID,
			"data":       waDataReply,
		}
		fmt.Println("RESPONSE SEND MESSAGE", reflect.TypeOf(resp))
		utils.LogJson(resp)
		// msgID, ok := resp.(*models.WhatsappMessageModel).MessageInfo["ID"].(string)
		// if ok {
		// 	waDataReply.MessageID = &msgID
		// 	h.erpContext.DB.Save(&waDataReply)
		// }

		msgNotifStr, _ := json.Marshal(msgNotif)
		h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	} else {
		template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(*templateID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, msg := range template.Messages {
			waDataReply.Message = parseMsgTemplate(*session.Contact, &member, msg.Body)

			h.customerRelationshipService.WhatsappService.SetMsgData(h.waService, &waDataReply, to, msg.Files, msg.Products, true, nil)
			_, err := customer_relationship.SendCustomerServiceMessage(h.customerRelationshipService.WhatsappService)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			msgNotif := gin.H{
				"message":    input.Message,
				"command":    "WHATSAPP_RECEIVED",
				"session_id": session.ID,
				"data":       waDataReply,
			}
			msgNotifStr, _ := json.Marshal(msgNotif)
			h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
				url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
				return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
			})
		}
	}

	session.LastMessage = waDataReply.Message
	session.LastOnlineAt = &now
	err = h.erpContext.DB.Save(&session).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 	if templateID != nil {
	// 		member := c.MustGet("member").(models.MemberModel)
	// 		template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(*templateID)
	// 		if err != nil {
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		for _, msg := range template.Messages {
	// 			now := time.Now()
	// 			waDataReply.ID = utils.Uuid()
	// 			waDataReply.Message = parseMsgTemplate(*session.Contact, &member, msg.Body)
	// 			waDataReply.CreatedAt = &now
	// 			err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 			if err != nil {
	// 				log.Println(err)
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 				return
	// 			}
	// 			session.LastMessage = waDataReply.Message
	// 			session.LastOnlineAt = &now
	// 			err = h.erpContext.DB.Save(&session).Error
	// 			if err != nil {
	// 				log.Println(err)
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 				return
	// 			}

	// 			msgNotif := gin.H{
	// 				"message":    waDataReply.Message,
	// 				"command":    "WHATSAPP_RECEIVED",
	// 				"session_id": session.ID,
	// 				"data":       waDataReply,
	// 			}
	// 			msgNotifStr, _ := json.Marshal(msgNotif)
	// 			h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 				url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 				return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 			})
	// 			to := waDataReply.Receiver
	// 			if waDataReply.IsGroup {
	// 				to = waDataReply.Session
	// 			}

	// 			thumbnail, restFiles := models.GetThumbnail(msg.Files)

	// 			var fileType, fileUrl string
	// 			if thumbnail != nil {
	// 				fileType = "image"
	// 				fileUrl = thumbnail.URL
	// 			}

	// 			msgData := whatsmeow_client.WaMessage{
	// 				JID:      waDataReply.JID,
	// 				Text:     waDataReply.Message,
	// 				To:       to,
	// 				IsGroup:  waDataReply.IsGroup,
	// 				FileType: fileType,
	// 				FileUrl:  fileUrl,
	// 			}

	// 			h.waService.SetChatData(msgData)
	// 			_, err := objects.SendChatMessage(h.waService)
	// 			if err != nil {
	// 				log.Println(err)
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 				return
	// 			}

	// 			// h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 			// 	JID:     waDataReply.JID,
	// 			// 	Text:    waDataReply.Message,
	// 			// 	To:      to,
	// 			// 	IsGroup: waDataReply.IsGroup,
	// 			// })
	// 			for _, v := range restFiles {
	// 				msgFileData := whatsmeow_client.WaMessage{
	// 					JID:     waDataReply.JID,
	// 					To:      to,
	// 					IsGroup: waDataReply.IsGroup,
	// 					FileUrl: v.URL,
	// 				}
	// 				if strings.Contains(v.MimeType, "image") && v.URL != "" {
	// 					msgFileData.FileType = "image"
	// 				} else {
	// 					msgFileData.FileType = "file"
	// 				}
	// 				h.waService.SetChatData(msgFileData)
	// 				_, err := objects.SendChatMessage(h.waService)
	// 				if err != nil {
	// 					log.Println(err)
	// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 					return
	// 				}

	// 				// msgFile := mdl.WhatsappMessageModel{
	// 				// 	MediaURL: v.URL,
	// 				// 	MimeType: v.MimeType,
	// 				// }
	// 				waDataReply.ID = utils.Uuid()
	// 				waDataReply.MediaURL = v.URL
	// 				waDataReply.MimeType = v.MimeType
	// 				err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 				if err != nil {
	// 					log.Println(err)
	// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 					return
	// 				}
	// 				msgNotif := gin.H{
	// 					"message":    waDataReply.Message,
	// 					"command":    "WHATSAPP_RECEIVED",
	// 					"session_id": session.ID,
	// 					"data":       waDataReply,
	// 				}
	// 				msgNotifStr, _ := json.Marshal(msgNotif)
	// 				h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 					url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 					return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 				})
	// 			}

	// 			// for _, v := range msg.Files {

	// 			// 	now := time.Now()
	// 			// 	waDataReply.ID = utils.Uuid()
	// 			// 	waDataReply.Message = ""
	// 			// 	waDataReply.CreatedAt = &now
	// 			// 	waDataReply.MediaURL = v.URL
	// 			// 	waDataReply.MimeType = v.MimeType
	// 			// 	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 			// 	if err != nil {
	// 			// 		log.Println(err)
	// 			// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			// 		return
	// 			// 	}
	// 			// 	session.LastMessage = waDataReply.Message
	// 			// 	session.LastOnlineAt = &now
	// 			// 	err = h.erpContext.DB.Save(&session).Error
	// 			// 	if err != nil {
	// 			// 		log.Println(err)
	// 			// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			// 		return
	// 			// 	}

	// 			// 	msgNotif := gin.H{
	// 			// 		"message":    waDataReply.Message,
	// 			// 		"command":    "WHATSAPP_RECEIVED",
	// 			// 		"session_id": session.ID,
	// 			// 		"data":       waDataReply,
	// 			// 	}
	// 			// 	msgNotifStr, _ := json.Marshal(msgNotif)
	// 			// 	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 			// 		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 			// 		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 			// 	})

	// 			// 	if strings.Contains(v.MimeType, "image") && v.URL != "" {
	// 			// 		resp, _ := h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 			// 			JID:      waDataReply.JID,
	// 			// 			Text:     "",
	// 			// 			To:       to,
	// 			// 			IsGroup:  waDataReply.IsGroup,
	// 			// 			FileType: "image",
	// 			// 			FileUrl:  v.URL,
	// 			// 		})
	// 			// 		fmt.Println("RESPONSE", resp)
	// 			// 	} else {
	// 			// 		resp, _ := h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 			// 			JID:      waDataReply.JID,
	// 			// 			Text:     "",
	// 			// 			To:       to,
	// 			// 			IsGroup:  waDataReply.IsGroup,
	// 			// 			FileType: "document",
	// 			// 			FileUrl:  v.URL,
	// 			// 		})
	// 			// 		fmt.Println("RESPONSE", resp)
	// 			// 	}

	// 			// }

	// 			for _, v := range msg.Products {
	// 				desc := ""
	// 				var images []models.FileModel
	// 				h.erpContext.DB.Where("ref_id = ? and ref_type = ?", v.ID, "product").Find(&images)

	// 				if v.Description != nil {
	// 					desc = *v.Description
	// 				}
	// 				dataMsg := fmt.Sprintf(`*%s*
	// _%s_

	// %s
	// 				`, v.DisplayName, utils.FormatRupiah(v.Price), desc)
	// 				productMsg := whatsmeow_client.WaMessage{
	// 					JID:      waDataReply.JID,
	// 					Text:     dataMsg,
	// 					To:       to,
	// 					IsGroup:  waDataReply.IsGroup,
	// 					FileType: "image",
	// 				}

	// 				now := time.Now()
	// 				waDataReply.ID = utils.Uuid()
	// 				waDataReply.Message = dataMsg
	// 				waDataReply.CreatedAt = &now
	// 				if len(images) > 0 {
	// 					waDataReply.MediaURL = images[0].URL
	// 					waDataReply.MimeType = images[0].MimeType
	// 				}

	// 				err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 				if err != nil {
	// 					log.Println(err)
	// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 					return
	// 				}
	// 				session.LastMessage = waDataReply.Message
	// 				session.LastOnlineAt = &now
	// 				err = h.erpContext.DB.Save(&session).Error
	// 				if err != nil {
	// 					log.Println(err)
	// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 					return
	// 				}

	// 				msgNotif := gin.H{
	// 					"message":    waDataReply.Message,
	// 					"command":    "WHATSAPP_RECEIVED",
	// 					"session_id": session.ID,
	// 					"data":       waDataReply,
	// 				}
	// 				msgNotifStr, _ := json.Marshal(msgNotif)
	// 				h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 					url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 					return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 				})

	// 				if len(images) > 0 {
	// 					productMsg.FileType = "image"
	// 					productMsg.FileUrl = images[0].URL
	// 					waDataReply.MediaURL = images[0].URL
	// 					waDataReply.MimeType = images[0].MimeType
	// 				}

	// 				h.waService.SendMessage(productMsg)
	// 			}
	// 			time.Sleep(time.Second * 1)

	// 		}

	// 		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	// 		return
	// 	}

	// 	waDataReply.ID = utils.Uuid()
	// 	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 	if err != nil {
	// 		log.Println(err)
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	session.LastMessage = input.Message
	// 	session.LastOnlineAt = &now

	// 	err = h.erpContext.DB.Save(&session).Error
	// 	if err != nil {
	// 		log.Println(err)
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	msgNotif := gin.H{
	// 		"message":    input.Message,
	// 		"command":    "WHATSAPP_RECEIVED",
	// 		"session_id": session.ID,
	// 		"data":       waDataReply,
	// 	}
	// 	msgNotifStr, _ := json.Marshal(msgNotif)
	// 	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 	})
	// 	to := waDataReply.Receiver
	// 	if waDataReply.IsGroup {
	// 		to = waDataReply.Session
	// 	}

	// 	var fileType, fileUrl string
	// 	if thumbnail != nil {
	// 		fileType = "image"
	// 		fileUrl = thumbnail.URL
	// 	}
	// 	resp, err := h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 		JID:      waDataReply.JID,
	// 		Text:     waDataReply.Message,
	// 		To:       to,
	// 		IsGroup:  waDataReply.IsGroup,
	// 		FileType: fileType,
	// 		FileUrl:  fileUrl,
	// 	})

	// 	respData, ok := resp.(map[string]any)["data"].(map[string]any)
	// 	if ok {
	// 		// utils.LogJson(respData["ID"])
	// 		msgID, ok2 := respData["ID"].(string)
	// 		if ok2 {
	// 			waDataReply.MessageID = &msgID
	// 			h.erpContext.DB.Save(&waDataReply)
	// 		}

	// 	}

	// 	if err != nil {
	// 		log.Println(err)
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	msg, err := h.customerRelationshipService.WhatsappService.GetWhatsappLastMessages(session.JID, session.Session)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	for _, v := range attachments {

	// 		now := time.Now()
	// 		waDataReply.ID = utils.Uuid()
	// 		waDataReply.Message = ""
	// 		waDataReply.CreatedAt = &now
	// 		waDataReply.MediaURL = v.URL
	// 		waDataReply.MimeType = v.MimeType
	// 		err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 		if err != nil {
	// 			log.Println(err)
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		session.LastMessage = waDataReply.Message
	// 		session.LastOnlineAt = &now
	// 		err = h.erpContext.DB.Save(&session).Error
	// 		if err != nil {
	// 			log.Println(err)
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 		msgNotif := gin.H{
	// 			"message":    waDataReply.Message,
	// 			"command":    "WHATSAPP_RECEIVED",
	// 			"session_id": session.ID,
	// 			"data":       waDataReply,
	// 		}
	// 		msgNotifStr, _ := json.Marshal(msgNotif)
	// 		h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 		})

	// 		if strings.Contains(v.MimeType, "image") && v.URL != "" {
	// 			resp, _ := h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 				JID:      waDataReply.JID,
	// 				Text:     "",
	// 				To:       to,
	// 				IsGroup:  waDataReply.IsGroup,
	// 				FileType: "image",
	// 				FileUrl:  v.URL,
	// 			})
	// 			fmt.Println("RESPONSE", resp)
	// 		} else {
	// 			resp, _ := h.waService.SendMessage(whatsmeow_client.WaMessage{
	// 				JID:      waDataReply.JID,
	// 				Text:     "",
	// 				To:       to,
	// 				IsGroup:  waDataReply.IsGroup,
	// 				FileType: "document",
	// 				FileUrl:  v.URL,
	// 			})
	// 			fmt.Println("RESPONSE", resp)
	// 		}

	// 		time.Sleep(time.Millisecond * 500)
	// 	}

	// 	for _, v := range input.Products {
	// 		desc := ""
	// 		var images []models.FileModel
	// 		h.erpContext.DB.Where("ref_id = ? and ref_type = ?", v.ID, "product").Find(&images)

	// 		if v.Description != nil {
	// 			desc = *v.Description
	// 		}
	// 		dataMsg := fmt.Sprintf(`*%s*
	// _%s_

	// %s
	// 		`, v.DisplayName, utils.FormatRupiah(v.Price), desc)
	// 		productMsg := whatsmeow_client.WaMessage{
	// 			JID:      waDataReply.JID,
	// 			Text:     dataMsg,
	// 			To:       to,
	// 			IsGroup:  waDataReply.IsGroup,
	// 			FileType: "image",
	// 		}

	// 		now := time.Now()
	// 		waDataReply.ID = utils.Uuid()
	// 		waDataReply.Message = dataMsg
	// 		waDataReply.CreatedAt = &now
	// 		if len(images) > 0 {
	// 			waDataReply.MediaURL = images[0].URL
	// 			waDataReply.MimeType = images[0].MimeType
	// 		}

	// 		err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	// 		if err != nil {
	// 			log.Println(err)
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		session.LastMessage = waDataReply.Message
	// 		session.LastOnlineAt = &now
	// 		err = h.erpContext.DB.Save(&session).Error
	// 		if err != nil {
	// 			log.Println(err)
	// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 		msgNotif := gin.H{
	// 			"message":    waDataReply.Message,
	// 			"command":    "WHATSAPP_RECEIVED",
	// 			"session_id": session.ID,
	// 			"data":       waDataReply,
	// 		}
	// 		msgNotifStr, _ := json.Marshal(msgNotif)
	// 		h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
	// 			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
	// 			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 		})

	// 		if len(images) > 0 {
	// 			productMsg.FileType = "image"
	// 			productMsg.FileUrl = images[0].URL
	// 			waDataReply.MediaURL = images[0].URL
	// 			waDataReply.MimeType = images[0].MimeType
	// 		}

	// 		h.waService.SendMessage(productMsg)
	// 	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": waDataReply})
}

func (h *WhatsappHandler) GetDevice(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	device, err := h.waService.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(device, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *WhatsappHandler) DeleteDevice(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	err := h.waService.DeviceDelete(jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
func (h *WhatsappHandler) GetQR(c *gin.Context) {
	session := c.Params.ByName("session")
	if session == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	respBody, err := h.waService.GetQR(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": respBody})
}
func (h *WhatsappHandler) CreateQR(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	input := struct {
		Webhook   string `json:"webhook"`
		HeaderKey string `json:"header_key"`
		Session   string `json:"session"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respBody, err := h.waService.CreateQR(input.Session, input.Webhook, input.HeaderKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// respData, ok := resp["data"].(map[string]interface{})
	// if ok {

	// }

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": resp})
}
func (h *WhatsappHandler) ExportHandler(c *gin.Context) {
	var input struct {
		StartDate time.Time `json:"start_date"`
		EndDate   time.Time `json:"end_date"`
		MemberIDs []string  `json:"member_ids"`
		TagIDs    []string  `json:"tag_ids"`
		Sessions  []string  `json:"sessions"`
	}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var sessions []models.WhatsappMessageSession
	// err = h.erpContext.DB.Preload("Contact").Where("session_name IN (?)", input.ConnectionNames).Find(&sessions).Error
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// Assuming `sessions` is already populated with the data to be exported
	file := excelize.NewFile()
	sheet1 := file.GetSheetName(0)

	headerStyle, err := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 14,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#DCE6F1"}, // Soft blue
			Pattern: 1,
		},
	})
	row := 1
	headers := []string{"Tgl", "Nama", "Nomor WA", "Percakapan", "Tag"}
	colWidth := []float64{15, 20, 20, 75, 30}
	for i, header := range headers {
		file.SetCellValue(sheet1, fmt.Sprintf("%s%d", utils.NumToAlphabet(i+1), row), header)
		file.SetColWidth(sheet1, utils.NumToAlphabet(i+1), utils.NumToAlphabet(i+1), colWidth[i])
		// Apply styles: bold font, bigger font, center align, and soft blue background

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		file.SetCellStyle(sheet1, fmt.Sprintf("A%d", row), fmt.Sprintf("%s%d", utils.NumToAlphabet(len(headers)), row), headerStyle)

	}

	companyID := c.GetHeader("ID-Company")

	var messages []models.WhatsappMessageModel
	db := h.erpContext.DB.Preload("Member.User").Preload("Contact.Tags").Model(&models.WhatsappMessageModel{}).Where("whatsapp_messages.company_id = ?", companyID)
	db = db.Joins("JOIN contacts on contacts.id = whatsapp_messages.contact_id")
	if len(input.TagIDs) > 0 {
		db = db.Joins("LEFT JOIN contact_tags on contact_tags.contact_model_id = contacts.id")
		db = db.Where("contact_tags.tag_model_id IN (?)", input.TagIDs)
	}

	if len(input.Sessions) > 0 {
		db = db.Where("whatsapp_messages.j_id IN (?)", input.Sessions)
	}
	db = db.Where("whatsapp_messages.created_at between ? and ?", input.StartDate, input.EndDate)
	db = db.Where("whatsapp_messages.is_from_me = ?", false)
	db = db.Where("whatsapp_messages.message != ''")
	db = db.Order("created_at ASC")
	db.Find(&messages)

	row++

	for _, msg := range messages {
		var tags = []string{}
		if msg.Contact == nil {
			continue
		}
		for _, tag := range msg.Contact.Tags {
			tags = append(tags, tag.Name)
		}
		cells := []string{msg.CreatedAt.Format("02-01-2006 15:04"), msg.Contact.Name, *msg.Contact.Phone, msg.Message, strings.Join(tags, ", ")}

		for i, c := range cells {
			file.SetCellValue(sheet1, fmt.Sprintf("%s%d", utils.NumToAlphabet(i+1), row), c)

		}
		row++
	}

	// for i, message := range messages {

	// }

	var buf bytes.Buffer
	if err := file.Write(&buf); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to write XLSX file"})
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=sessions.xlsx")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", buf.Bytes())

}
func (h *WhatsappHandler) UpdateWebhook(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	input := struct {
		Webhook   string `json:"webhook"`
		HeaderKey string `json:"header_key"`
		Session   string `json:"session"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.waService.UpdateWebhook(input.Session, input.Webhook, input.HeaderKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *WhatsappHandler) GetMessagesHandler(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	messages, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessages(*c.Request, c.DefaultQuery("search", ""), jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := messages.Items.(*[]models.WhatsappMessageModel)
	messages.Items = reverse(*items)

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": messages})
}
func (h *WhatsappHandler) GetConversationsHandler(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	conversations, err := h.customerRelationshipService.WhatsappService.GetMessageSession(jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, v := range conversations {
		if v.IsGroup {
			groupInfo, err := h.waService.GetGroupInfo(jid, v.Session)

			if err == nil {
				groupName, ok := groupInfo["data"].(map[string]interface{})["Name"].(string)
				if ok {
					v.MessageInfo["GroupName"] = groupName
				}
			}
		}
		conversations[i] = v
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": conversations})
}

func reverse(messages []models.WhatsappMessageModel) []models.WhatsappMessageModel {
	for i, j := 0, len(messages)-1; i < len(messages)/2; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}

func (h *WhatsappHandler) WhatsappGetNumberHandler(c *gin.Context) {
	code, err := services.REDIS.Get(*h.erpContext.Ctx, "verify:"+c.Param("code")).Result()

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid code"})
		return
	}
	data := map[string]interface{}{}
	err = json.Unmarshal([]byte(code), &data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid code"})
		return
	}

	c.JSON(200, gin.H{"message": "Success", "data": data["phone_number"]})
}
func (h *WhatsappHandler) WhatsappRegisterHandler(c *gin.Context) {
	var input struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Code    string `json:"code"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dataStr, err := services.REDIS.Get(*h.erpContext.Ctx, "verify:"+input.Code).Result()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{}

	err = json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	phone := data["phone_number"].(string)
	companyID, ok := data["company_id"].(string)
	if !ok {
		companyID = ""
	}

	if phone != input.Phone {
		c.JSON(400, gin.H{"error": "invalid code"})
		return
	}

	phoneNumber := utils.ParsePhoneNumber(input.Phone, "ID")

	var member models.ContactModel = models.ContactModel{
		Name:       input.Name,
		Email:      input.Email,
		Phone:      &phoneNumber,
		Address:    input.Address,
		CompanyID:  &companyID,
		IsCustomer: true,
	}
	member.ID = utils.Uuid()
	err = h.erpContext.DB.Create(&member).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// sendWAMessage(h.erpContext, data["jid"].(string), phoneNumber, "Registration has been completed")

	waData := whatsmeow_client.WaMessage{
		JID:     data["jid"].(string),
		Text:    "Registration has been completed",
		To:      phoneNumber,
		IsGroup: false,
	}
	h.waService.SetChatData(waData)
	_, err = objects.SendChatMessage(h.waService)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Registration has been completed"})
}
func (h *WhatsappHandler) WhatsappWebhookHandler(c *gin.Context) {

	var body whatsapp.MsgObject

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("RECEIPT MESSAGE")
	utils.LogJson(body)
	// GET CONNECTION
	conn, err := h.appService.ConnectionService.GetConnectionBySession(body.SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msgType, ok := body.Info["Type"].(string)
	if ok && msgType == "read" {
		msgIDs, ok := body.Info["MessageIDs"].([]any)
		if ok {
			for _, v := range msgIDs {
				var msg models.WhatsappMessageModel
				h.erpContext.DB.Where("message_id = ?", v.(string)).First(&msg)
				msg.IsRead = true
				h.erpContext.DB.Save(&msg)
				var whatsappSession *models.WhatsappMessageSession
				err = h.erpContext.DB.First(&whatsappSession, "session = ?", body.Info["Chat"].(string)).Error
				if err == nil {
					msgNotif := gin.H{
						"command":     "WHATSAPP_MESSAGE_READ",
						"session_id":  whatsappSession.ID,
						"message_ids": msgIDs,
					}
					b, _ := json.Marshal(msgNotif)
					h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
						url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
						return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
					})
				}

				// h.appService.WhatsappService.MarkMessageAsRead(v)
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}

	convMsg := ""
	var quotedMsg, quotedMsgID *string
	if body.Message.Conversation != nil {
		convMsg = *body.Message.Conversation
	}

	var sessionAuth *models.ContactModel
	if conn.SessionAuth {
		// fmt.Println("KENAPA KESINI", conn.SessionAuth)
		// CHECK IS PHONE NUMBER REGISTERED
		if !h.appService.ConnectionService.IsPhoneNumberRegistered(body.Sender) {
			randomStr := utils.RandString(8, false)

			data := map[string]interface{}{
				"phone_number": body.Sender,
				"jid":          body.JID,
				"company_id":   conn.CompanyID,
			}

			b, _ := json.Marshal(data)
			services.REDIS.Set(*h.erpContext.Ctx, fmt.Sprintf("verify:%s", randomStr), string(b), 30*time.Minute)
			replyData := whatsmeow_client.WaMessage{
				JID: body.JID,
				Text: fmt.Sprintf(`Selamat datang di %s
Anda belum terdaftar di sistem kami, silakan lakukan pendaftaran terlebih dahulu dengan mengikuti tautan berikut:

%s

*TERIMA KASIH*`, conn.Name, config.App.Server.FrontendURL+"/member/register/"+randomStr),
				To:      body.Sender,
				IsGroup: false,
			}

			h.waService.SendMessage(replyData)
			return
		}
		// CHECK SESSION AUTH

		session, err := h.appService.ConnectionService.GetActiveSession(body.Sender)
		if err != nil {
		}
		sessionAuth = session

		if convMsg == "LOGIN" && session == nil {
			err := doLogin(h.erpContext, h.waService, body.JID, body.Sender, conn)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			return
		}

		if session == nil {
			// sendWAMessage(h.erpContext, body.JID, body.Sender, "Anda belum Login, silakan ketik *LOGIN* lalu kirim untuk melakukan login")

			waData := whatsmeow_client.WaMessage{
				JID:     body.JID,
				Text:    "Anda belum Login, silakan ketik *LOGIN* lalu kirim untuk melakukan login",
				To:      body.Sender,
				IsGroup: false,
			}
			h.waService.SetChatData(waData)
			_, err = objects.SendChatMessage(h.waService)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			return
		}

	} else {
		// CHECK IS PHONE NUMBER REGISTERED
		var contact models.ContactModel
		err := h.erpContext.DB.Where("phone = ? AND company_id = ?", body.Sender, *conn.CompanyID).First(&contact).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			contact.Phone = &body.Sender
			pushName, ok := body.Info["PushName"].(string)
			if ok {
				contact.Name = pushName
				contact.CompanyID = conn.CompanyID
				contact.IsCustomer = true
			}
			if body.Info["IsGroup"].(bool) {
				contact.Name = body.GroupInfo.Name
			}
			err := h.erpContext.DB.Create(&contact).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}

		sessionAuth = &contact
	}

	profilePic, _ := sessionAuth.GetProfilePicture(h.erpContext.DB)
	if profilePic == nil && body.ProfilePic != "" {
		resp, err := http.Get(body.ProfilePic)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		byteValue, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		path := filepath.Join("assets", "files", sessionAuth.ID+".jpg")
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err := os.WriteFile(path, byteValue, 0644); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mediaURLSaved := config.App.Server.BaseURL + "/" + path

		h.erpContext.DB.Create(&models.FileModel{
			FileName: sessionAuth.Name,
			Path:     path,
			URL:      mediaURLSaved,
			RefID:    sessionAuth.ID,
			RefType:  "contact",
		})

	}
	// fmt.Println("session", sessionAuth)

	if body.Sender == "status" {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}

	if body.Info["Category"].(string) == "peer" {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}

	var fileUrl, mimeType string
	if body.Message.Conversation != nil {
		convMsg = *body.Message.Conversation
	}
	if body.Message.ImageMessage != nil {
		convMsg = body.Message.ImageMessage.Caption
		mimeType = body.Message.ImageMessage.Mimetype
	}
	if body.Message.VideoMessage != nil {
		convMsg = body.Message.VideoMessage.Caption
		mimeType = body.Message.VideoMessage.Mimetype
	}

	if body.Message.DocumentMessage != nil {
		convMsg = body.Message.DocumentMessage.Caption
		mimeType = body.Message.DocumentMessage.Mimetype
	}
	if body.Message.ExtendedTextMessage != nil {
		convMsg = body.Message.ExtendedTextMessage.Text
		quotedMsg = &body.Message.ExtendedTextMessage.ContextInfo.QuotedMessage.Conversation
		quotedMsgID = &body.Message.ExtendedTextMessage.ContextInfo.StanzaID
	}
	var mediaURLSaved string
	if body.MediaPath != "" {
		mediaURL := config.App.Whatsapp.BaseURL + body.MediaPath

		resp, err := http.Get(mediaURL)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()
		byteValue, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		path := filepath.Join("assets", body.MediaPath)
		os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err := os.WriteFile(path, byteValue, 0644); err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		mediaURLSaved = config.App.Server.FrontendURL + "/" + path
		fileUrl = mediaURLSaved

	}

	fmt.Println("PROFILE PICTURE", body.ProfilePic)

	infoByte, err := json.Marshal(body.Info)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msgID := body.Info["ID"].(string)
	var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
		Sender:          body.Sender,
		Message:         convMsg,
		MimeType:        mimeType,
		MediaURL:        fileUrl,
		Info:            string(infoByte),
		Session:         body.SessionID,
		JID:             body.JID,
		IsFromMe:        body.Info["IsFromMe"].(bool),
		IsGroup:         body.Info["IsGroup"].(bool),
		MessageID:       &msgID,
		QuotedMessage:   quotedMsg,
		QuotedMessageID: quotedMsgID,
	}

	if sessionAuth != nil {
		waData.ContactID = &sessionAuth.ID
		waData.CompanyID = sessionAuth.CompanyID
	}
	now := time.Now()
	var whatsappSession *models.WhatsappMessageSession
	err = h.erpContext.DB.First(&whatsappSession, "session = ? AND session_name = ?", body.SessionID, body.SessionName).Error // FIXING SESSION UNIQUE
	if errors.Is(err, gorm.ErrRecordNotFound) {
		refType := "connection"
		// CREATE NEW SESSION
		sessionData := models.WhatsappMessageSession{
			JID:          body.JID,
			Session:      body.SessionID,
			SessionName:  body.SessionName,
			LastOnlineAt: &now,
			LastMessage:  convMsg,
			RefID:        &conn.ID,
			RefType:      &refType,
			IsGroup:      body.Info["IsGroup"].(bool),
		}
		if sessionAuth != nil {
			sessionData.CompanyID = sessionAuth.CompanyID
			sessionData.ContactID = &sessionAuth.ID
		}
		fmt.Println("CREATE SESSION")

		h.erpContext.DB.Create(&sessionData)
		whatsappSession = &sessionData
		if conn.NewSessionColumnID != nil {
			// CREATE NEW TASK
			senderName := sessionAuth.Name
			if senderName == "" {
				senderName = body.Sender
			}
			refType := "whatsapp_session"
			task := models.TaskModel{
				Name:      fmt.Sprintf("%s - %s", senderName, *sessionAuth.Phone),
				ProjectID: *conn.ProjectID,
				ColumnID:  conn.NewSessionColumnID,
				// StartDate:      &formResponse.CreatedAt,
				// EndDate:        &formResponse.CreatedAt,
				Description: convMsg,
				RefID:       &sessionData.ID,
				RefType:     &refType,
			}
			err = h.pmService.TaskService.CreateTask(&task)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}

			msg := gin.H{
				"message":   "Task created successfully",
				"column_id": conn.NewSessionColumnID,
			}
			b, _ := json.Marshal(msg)
			h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
				url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
				return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
			})
		}

		waData.IsNew = true
		// h.erpContext.DB.Create(&waData)
		whatsappSession.ID = sessionData.ID
		fmt.Println("SESSION ID", whatsappSession.ID)
	} else {
		// var lastOnlineAt time.Time
		// if whatsappSession.LastOnlineAt != nil {
		// 	lastOnlineAt = *whatsappSession.LastOnlineAt
		// } else {
		// }
		lastOnlineAt := time.Now()
		whatsappSession.LastOnlineAt = &lastOnlineAt
		whatsappSession.LastMessage = convMsg
		h.erpContext.DB.Save(&whatsappSession)

		if conn.IdleColumnID != nil && !lastOnlineAt.IsZero() {
			// CREATE NEW TASK
			if (now.Sub(lastOnlineAt).Hours() / 24) > conn.IdleDuration {
				senderName := sessionAuth.Name
				if senderName == "" {
					senderName = body.Sender
				}
				refType := "whatsapp_session"
				task := models.TaskModel{
					Name:      fmt.Sprintf("%s - %s", senderName, *sessionAuth.Phone),
					ProjectID: *conn.ProjectID,
					ColumnID:  conn.IdleColumnID,
					// StartDate:      &formResponse.CreatedAt,
					// EndDate:        &formResponse.CreatedAt,
					Description: convMsg,
					RefID:       &whatsappSession.ID,
					RefType:     &refType,
				}
				err = h.pmService.TaskService.CreateTask(&task)
				if err != nil {
					c.JSON(500, gin.H{"error": err.Error()})
					return
				}

				msg := gin.H{
					"message":   "Task created successfully",
					"column_id": conn.IdleColumnID,
				}
				b, _ := json.Marshal(msg)
				h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
					url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
					return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
				})
			}
		}
	}
	if ok && msgType == "reaction" {
		if body.Message.ReactionMessage != nil {
			// sender
			sender, ok := body.Info["Chat"].(string)
			if ok {
				splitSender := strings.SplitN(sender, "@", 2)
				if len(splitSender) > 1 {
					sender = splitSender[0]

					contact, err := h.contactService.GetContactByPhone(sender, *conn.CompanyID)
					if err == nil {

						msg, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageByMessageID(body.Message.ReactionMessage.Key.ID)
						if err == nil {
							reactionData := models.WhatsappMessageReaction{
								BaseModel: shared.BaseModel{
									ID: utils.Uuid(),
								},
								ContactID: &contact.ID,
								Reaction:  body.Message.ReactionMessage.Text,
							}
							h.customerRelationshipService.WhatsappService.AddMessageReaction(*msg, reactionData)
							reactions, _ := h.customerRelationshipService.WhatsappService.GetWhatsappMessageReactions(msg.ID)
							msg := gin.H{
								"message":    "",
								"command":    "WHATSAPP_MESSAGE_REACTIONS",
								"session_id": whatsappSession.ID,
								"message_id": msg.ID,
								"data":       reactions,
							}
							b, _ := json.Marshal(msg)
							h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
								url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
								return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
							})
						}
					}
				}
			}

			// fmt.Println("RECEIPT REACTION MESSAGE", body.Message.ReactionMessage
		}

		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}
	whatsappSession.IsGroup = waData.IsGroup
	fmt.Println("UPDATE SESSION")
	err = h.erpContext.DB.Omit(clause.Associations).Save(&whatsappSession).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
	if err != nil {
		// log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	waData.Contact = sessionAuth
	waData.SentAt = &now
	msg := gin.H{
		"message":    convMsg,
		"command":    "WHATSAPP_RECEIVED",
		"session_id": whatsappSession.ID,
		"data":       waData,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	autopilot := false

	// fmt.Println("AUTO RESPONSE TIME", conn.AutoResponseStartTime, conn.AutoResponseEndTime)
	if conn.AutoResponseStartTime != nil && conn.AutoResponseEndTime != nil {
		fmt.Println("AUTO RESPONSE TIME", *conn.AutoResponseStartTime, *conn.AutoResponseEndTime)
		autoResponseStartTime, err := time.ParseInLocation("2006-01-02 15:04", now.Format("2006-01-02")+" "+*conn.AutoResponseStartTime, time.Local)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// fmt.Println("START TIME", now, autoResponseStartTime.Format("2006-01-02 15:04"))
		autoResponseEndTime, err := time.ParseInLocation("2006-01-02 15:04", now.Format("2006-01-02")+" "+*conn.AutoResponseEndTime, time.Local)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("BETWEEN", autoResponseStartTime.Format("2006-01-02 15:04"), "<", now.Format("2006-01-02 15:04"), ">", autoResponseEndTime.Format("2006-01-02 15:04"))
		if now.After(autoResponseStartTime) && now.Before(autoResponseEndTime) {
			autopilot = true
		}

	}

	fmt.Println("AUTO PILOT", autopilot)
	if whatsappSession.IsHumanAgent {
		autopilot = false
	}

	var replyResponse *models.WhatsappMessageModel

	if conn.GeminiAgent == nil && conn.IsAutoPilot && autopilot && conn.AutoResponseMessage != "" {
		// sendWAMessage(h.erpContext, body.JID, body.Sender, conn.AutoResponseMessage)
		waData := whatsmeow_client.WaMessage{
			JID:     body.JID,
			Text:    conn.AutoResponseMessage,
			To:      body.Sender,
			IsGroup: false,
		}
		h.waService.SetChatData(waData)
		_, err = objects.SendChatMessage(h.waService)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		replyResponse = &models.WhatsappMessageModel{
			Receiver:    body.Sender,
			Message:     conn.AutoResponseMessage,
			MimeType:    body.MimeType,
			Session:     body.SessionID,
			JID:         body.JID,
			IsFromMe:    true,
			Info:        string(infoByte),
			IsGroup:     body.Info["IsGroup"].(bool),
			IsAutoPilot: true,
		}

	}
	if conn.GeminiAgent != nil && conn.IsAutoPilot && autopilot {
		h.geminiService.SetupModel(conn.GeminiAgent.SetTemperature, conn.GeminiAgent.SetTopK, conn.GeminiAgent.SetTopP, conn.GeminiAgent.SetMaxOutputTokens, conn.GeminiAgent.ResponseMimetype, conn.GeminiAgent.Model)
		h.geminiService.SetUpSystemInstruction(fmt.Sprintf(`%s
		
%s`, conn.GeminiAgent.SystemInstruction, `

Tolong jawab dalam format : 
{
  "response": string,
  "type": string,
  "command": string,
  "params": object
}

Keterangan:
response: jawaban bila tipe nya pertanyaan
type: command atau question
command: jika tipe command
params: jika tipe command dibutuhkan parameter

`))

		var histories []models.GeminiHistoryModel
		err = h.erpContext.DB.Model(&models.GeminiHistoryModel{}).Find(&histories, "agent_id = ? and is_model = ?", conn.GeminiAgent.ID, true).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Agent histories is not found"})
			return
		}
		chatHistories := []map[string]any{}
		for _, v := range histories {
			chatHistories = append(chatHistories, map[string]any{
				"role":    "user",
				"content": v.Input,
			})
			chatHistories = append(chatHistories, map[string]any{
				"role":    "model",
				"content": v.Output,
			})
		}

		userHistories := []models.WhatsappMessageModel{}
		h.erpContext.DB.Model(&models.WhatsappMessageModel{}).Where("session = ?", body.SessionID).Order("created_at desc").Limit(100).Find(&userHistories)

		userHistories = reverse(userHistories)

		for _, v := range userHistories {
			if v.IsFromMe {
				chatHistories = append(chatHistories, map[string]any{
					"role":    "user",
					"content": v.Message,
				})
			} else {
				chatHistories = append(chatHistories, map[string]any{
					"role":    "model",
					"content": v.Message,
				})
			}

		}

		// utils.LogJson(chatHistories)
		output, err := h.geminiService.GenerateContent(*h.erpContext.Ctx, convMsg, chatHistories, "", "")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var response geminiResponse
		err = json.Unmarshal([]byte(output), &response)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		info := map[string]interface{}{
			"Timestamp": time.Now().Format(time.RFC3339),
		}
		infoByte, err := json.Marshal(info)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// sendWAMessage(h.erpContext, body.JID, body.Sender, response.Response)
		waData := whatsmeow_client.WaMessage{
			JID:     body.JID,
			Text:    response.Response,
			To:      body.Sender,
			IsGroup: false,
		}
		h.waService.SetChatData(waData)
		_, err = objects.SendChatMessage(h.waService)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		replyResponse = &models.WhatsappMessageModel{
			Receiver:    body.Sender,
			Message:     response.Response,
			MimeType:    body.MimeType,
			Session:     body.SessionID,
			JID:         body.JID,
			IsFromMe:    true,
			Info:        string(infoByte),
			IsGroup:     body.Info["IsGroup"].(bool),
			IsAutoPilot: true,
		}
	}

	fmt.Println("SESSION ID #2", whatsappSession.ID)
	if replyResponse != nil {
		if sessionAuth != nil {
			replyResponse.ContactID = &sessionAuth.ID
			replyResponse.CompanyID = sessionAuth.CompanyID
		}

		fmt.Println("SESSION ID #3", whatsappSession.ID)
		fmt.Println("SESSION ID #3 response", replyResponse)

		err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(replyResponse)
		if err != nil {
			// log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		replyResponse.SentAt = &now
		msg := gin.H{
			"message":    replyResponse.Message,
			"command":    "WHATSAPP_RECEIVED",
			"session_id": whatsappSession.ID,
			"data":       replyResponse,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})

		now = time.Now()
		whatsappSession.LastMessage = replyResponse.Message
		whatsappSession.LastOnlineAt = &now
		h.erpContext.DB.Where("id = ?", whatsappSession.ID).Model(&models.WhatsappMessageSession{}).Updates(&whatsappSession)
	}

}

func (h *WhatsappHandler) GetSessionsHandler(c *gin.Context) {
	session := c.Params.ByName("session")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	sessions, err := h.customerRelationshipService.WhatsappService.GetSessionMessageBySessionName(session, *c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	waSessions := sessions.Items.(*[]models.WhatsappMessageSession)
	newWaSessions := []models.WhatsappMessageSession{}
	for _, v := range *waSessions {
		fmt.Println("SESSION", v)
		var totalUnread int64
		h.erpContext.DB.Model(&models.WhatsappMessageModel{}).Where("session = ? AND is_read = ? and is_from_me = ?", v.Session, false, false).Count(&totalUnread)
		v.CountUnread = int(totalUnread)
		if v.RefType != nil {
			var refType = "connection"
			if *v.RefType == refType {
				var conn connection.ConnectionModel
				err = h.erpContext.DB.Select("id, session_name, name, color").First(&conn, "id = ?", v.RefID).Error
				if err == nil {
					v.Ref = &conn
				} else {
					// GET EXIST CONNECTION
					err := h.erpContext.DB.Select("id, session_name, name, color").First(&conn, "session_name = ?", v.SessionName).Error
					if err == nil {
						fmt.Println("NEW CONNECTION", conn)
						v.SessionName = conn.SessionName
						resp, err := h.waService.GetJIDBySessionName(conn.SessionName)
						if err == nil {
							v.JID = resp["jid"].(string)
						}
						// v.JID = conn.
						v.RefID = &conn.ID
						v.Ref = &conn
						h.erpContext.DB.Save(&v)
					}

				}
			}

		}
		newWaSessions = append(newWaSessions, v)
	}

	sessions.Items = &newWaSessions

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": sessions})
}

func (h *WhatsappHandler) UpdateSessionHandler(c *gin.Context) {
	var input models.WhatsappMessageSession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	var session models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = h.erpContext.DB.Model(&models.WhatsappMessageSession{}).Where("id = ?", sessionId).Updates(map[string]any{
		"is_human_agent": input.IsHumanAgent,
		"session":        input.Session,
		"session_name":   input.SessionName,
	}).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Session updated successfully"})
}

func (h *WhatsappHandler) DeleteSessionHandler(c *gin.Context) {
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")
	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	var session models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = h.erpContext.DB.Unscoped().Where("session = ?", session.Session).Delete(&models.WhatsappMessageModel{}).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := h.erpContext.DB.Unscoped().Delete(&models.WhatsappMessageSession{}, "id = ?", sessionId).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}
func (h *WhatsappHandler) ClearSessionHandler(c *gin.Context) {
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")
	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	var session models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = h.erpContext.DB.Unscoped().Where("session = ?", session.Session).Delete(&models.WhatsappMessageModel{}).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	session.LastMessage = ""
	session.LastOnlineAt = nil

	if err := h.erpContext.DB.Save(&session).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"command":    "WHATSAPP_CLEAR_MESSAGE",
		"session_id": sessionId,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(http.StatusOK, gin.H{"message": "Session Cleared successfully"})
}
func (h *WhatsappHandler) GetSessionDetailHandler(c *gin.Context) {
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session models.WhatsappMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	session.Contact.ProfilePicture, _ = session.Contact.GetProfilePicture(h.erpContext.DB)

	var connection connection.ConnectionModel
	err = h.erpContext.DB.First(&connection, "id = ?", session.RefID).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": session, "connection": connection})
}
func (h *WhatsappHandler) MarkAsReadHandler(c *gin.Context) {
	messageId := c.Param("messageId")
	sessionId := c.Query("session_id")

	msg, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessage(messageId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if msg.MessageID != nil && !msg.IsRead {
		err = h.waService.MarkAsRead(msg.JID, []string{*msg.MessageID}, msg.Sender)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
	}
	err = h.customerRelationshipService.WhatsappService.MarkMessageAsRead(messageId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	msgNotif := gin.H{
		"message":    "ok",
		"command":    "UPDATE_SESSION",
		"session_id": sessionId,
	}
	msgNotifStr, _ := json.Marshal(msgNotif)
	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.GetHeader("ID-Company"))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
func (h *WhatsappHandler) GetSessionMessagesHandler(c *gin.Context) {
	sessionId := c.Query("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session *models.WhatsappMessageSession
	err := h.erpContext.DB.First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	messages, err := h.cacheService.Remember("wa-msg-"+sessionId, 10*time.Second, func() (paginate.Page, error) {
		return h.customerRelationshipService.WhatsappService.GetMessageSessionChatBySessionName(session.Session, "", session.ContactID, *c.Request)
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messages.Items = reverse(*messages.Items.(*[]models.WhatsappMessageModel))

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": messages})
}

func doLogin(erpContext *context.ERPContext, waService *whatsmeow_client.WhatsmeowService, jid, sender string, conn *connection.ConnectionModel) error {
	var member *models.ContactModel
	err := erpContext.DB.Find(&member, "phone = ?", sender).Error
	if err != nil {
		return err
	}
	b, err := json.Marshal(member)
	if err != nil {
		return err
	}
	services.REDIS.Set(*erpContext.Ctx, fmt.Sprintf("session:%s", *member.Phone), string(b), 7*24*time.Hour)
	msgContent := fmt.Sprintf(`
Hallo %s
Selamat datang di %s

Session anda akan berlaku selama 7 hari:

*ADMIN*
			`, member.Name, conn.Name)
	// sendWAMessage(erpContext, jid, sender, msgContent)
	waData := whatsmeow_client.WaMessage{
		JID:     jid,
		Text:    msgContent,
		To:      sender,
		IsGroup: false,
	}
	waService.SetChatData(waData)
	_, err = objects.SendChatMessage(waService)
	if err != nil {
		return err
	}
	return nil
}

// func (h *WhatsappHandler) WebhookHandler(c *gin.Context) {
// 	setting, err := h.appService.GetSetting()
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var session *mdl.Member

// 	// b, err := io.ReadAll(c.Request.Body)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }

// 	// log.Println(string(b))

// 	var msg objects.MsgObject
// 	if err := c.ShouldBindJSON(&msg); err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	session, _ = h.appService.MemberService.GetActiveSession(msg.Sender)
// 	if !h.appService.MemberService.IsPhoneNumberRegistered(msg.Sender) {
// 		randomStr := utils.RandString(8)

// 		data := map[string]interface{}{
// 			"phone_number": msg.Sender,
// 			"jid":          msg.JID,
// 		}

// 		b, _ := json.Marshal(data)
// 		services.REDIS.Set(fmt.Sprintf("verify:%s", randomStr), string(b), 30*time.Minute)
// 		replyData := whatsmeow_client.WaMessage{
// 			JID: msg.JID,
// 			Text: fmt.Sprintf(`Selamat datang di Sistem Informasi WA Dinkes Kota Sukabumi
// Anda belum terdaftar di sistem kami, silakan lakukan pendaftaran terlebih dahulu dengan mengikuti tautan berikut:

// %s

// *TERIMA KASIH*`, config.App.Server.FrontendURL+"/member/register/"+randomStr),
// 			To:      msg.Sender,
// 			IsGroup: false,
// 		}

// 		h.waService.SendMessage(replyData)
// 		return
// 	}
// 	if msg.Message.Conversation != nil {
// 		conv := *msg.Message.Conversation
// 		splitConv := strings.Split(conv, " ")
// 		fmt.Println("splitConv", splitConv)
// 		if splitConv[0] != "" {
// 			for _, keyword := range eligibleKeyWords {
// 				if strings.Contains(strings.ToLower(*msg.Message.Conversation), strings.ToLower(keyword)) {
// 					infoByte, err := json.Marshal(msg.Info)
// 					if err != nil {
// 						log.Println(err)
// 						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 						return
// 					}
// 					var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
// 						Sender:   msg.Sender,
// 						Message:  conv,
// 						MimeType: msg.MimeType,
// 						Info:     string(infoByte),
// 						Session:  msg.SessionID,
// 						JID:      msg.JID,
// 						IsFromMe: msg.Info["IsFromMe"].(bool),
// 						IsGroup:  msg.Info["IsGroup"].(bool),
// 					}

// 					err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
// 					if err != nil {
// 						log.Println(err)
// 						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 						return
// 					}
// 					waData.MessageInfo = msg.Info
// 					msgWs := gin.H{
// 						"type":    "NEW_MESSAGE",
// 						"session": msg.SessionID,
// 						"data":    waData,
// 					}

// 					msgWsByte, _ := json.Marshal(msgWs)
// 					services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 						return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 					})
// 					h.erpContext.TempData = msg
// 					cmd.TextCommand(h.erpContext, session, msg.JID, msg.Sender, splitConv...)
// 					return
// 				}
// 			}
// 		}

// 	}

// 	if session == nil {
// 		utils.SendWAMessage(h.erpContext, msg.JID, msg.Sender, "Anda belum Login, silakan ketik *LOGIN* lalu kirim untuk melakukan login")
// 		return
// 	}

// 	if msg.Sender == "status" {
// 		c.JSON(http.StatusOK, gin.H{"message": "ok"})
// 		return
// 	}

// 	b, _ := json.MarshalIndent(msg, "", "  ")
// 	log.Println("Received message:", string(b))
// 	var msgContent, fileUrl, mimeType string
// 	if msg.Message.Conversation != nil {
// 		msgContent = *msg.Message.Conversation
// 	}
// 	if msg.Message.ImageMessage != nil {
// 		msgContent = msg.Message.ImageMessage.Caption
// 		mimeType = msg.Message.ImageMessage.Mimetype
// 	}
// 	if msg.Message.VideoMessage != nil {
// 		msgContent = msg.Message.VideoMessage.Caption
// 		mimeType = msg.Message.VideoMessage.Mimetype
// 	}

// 	if msg.Message.DocumentMessage != nil {
// 		msgContent = msg.Message.DocumentMessage.Caption
// 		mimeType = msg.Message.DocumentMessage.Mimetype
// 	}
// 	var mediaURLSaved string
// 	if msg.MediaPath != "" {
// 		mediaURL := config.App.Whatsapp.BaseURL + msg.MediaPath

// 		resp, err := http.Get(mediaURL)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer resp.Body.Close()
// 		byteValue, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		path := filepath.Join("assets", msg.MediaPath)
// 		os.MkdirAll(filepath.Dir(path), os.ModePerm)
// 		if err := os.WriteFile(path, byteValue, 0644); err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		mediaURLSaved = config.App.Server.BaseURL + "/" + path
// 		fileUrl = mediaURLSaved

// 	}

// 	infoByte, err := json.Marshal(msg.Info)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
// 		Sender:   msg.Sender,
// 		Message:  msgContent,
// 		MediaURL: mediaURLSaved,
// 		MimeType: msg.MimeType,
// 		Info:     string(infoByte),
// 		Session:  msg.SessionID,
// 		JID:      msg.JID,
// 		IsFromMe: msg.Info["IsFromMe"].(bool),
// 		IsGroup:  msg.Info["IsGroup"].(bool),
// 	}

// 	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	waData.MessageInfo = msg.Info
// 	msgWs := gin.H{
// 		"type":    "NEW_MESSAGE",
// 		"session": msg.SessionID,
// 		"data":    waData,
// 	}

// 	msgWsByte, _ := json.Marshal(msgWs)
// 	services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 		return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 	})
// 	splitJID := strings.Split(msg.JID, "@")
// 	splitSep := strings.Split(splitJID[0], ":")
// 	// utils.LogJson(msg.Message.ExtendedTextMessage)
// 	if msg.Message.ExtendedTextMessage != nil {

// 		contains := false
// 		for _, v := range msg.Message.ExtendedTextMessage.ContextInfo.MentionedJID {
// 			if v == splitSep[0]+"@"+splitJID[1] {
// 				contains = true
// 				break
// 			}
// 		}
// 		if contains {
// 			msgContent = msg.Message.ExtendedTextMessage.Text
// 		}

// 	}
// 	if waData.IsGroup && msg.Message.ExtendedTextMessage == nil {
// 		msgContent = ""
// 	}
// 	if msgContent != "" {

// 		if h.erpContext.ThirdPartyServices["GEMINI"] != nil && setting.AutoPilot {
// 			histories, err := services.REDIS.LRange(config.App.Redis.Key+"-"+msg.Sender, 0, 20).Result()
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			// fmt.Println("histories", histories)

// 			for i, j := 0, len(histories)-1; i < j; i, j = i+1, j-1 {
// 				histories[i], histories[j] = histories[j], histories[i]
// 			}

// 			userHistories := []map[string]interface{}{}
// 			for _, v := range histories {
// 				var history objects.History
// 				if err := json.Unmarshal([]byte(v), &history); err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}
// 				userHistories = append(userHistories, map[string]interface{}{
// 					"role":    "user",
// 					"content": history.Input,
// 				})
// 				userHistories = append(userHistories, map[string]interface{}{
// 					"role":    "model",
// 					"content": history.Output,
// 				})
// 			}

// 			// SEND TO AI
// 			log.Println("SEND MESSAGE TO AI", msgContent)
// 			resp, err := h.erpContext.ThirdPartyServices["GEMINI"].(*google.GeminiService).GenerateContent(*h.erpContext.Ctx, msgContent, userHistories, fileUrl, mimeType)
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}
// 			log.Println("Generated content:", resp)

// 			parseData := map[string]interface{}{}
// 			if err := json.Unmarshal([]byte(resp), &parseData); err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			responseMsg := ""
// 			if parseData["type"] == "answer" {
// 				responseMsg = parseData["conversation"].(string)
// 			}
// 			if parseData["type"] == "command" {
// 				log.Println("Received command:", resp)
// 				msgContent, ok := parseData["message"].(string)
// 				if ok {
// 					responseMsg = msgContent
// 				}
// 				respCommand, err := command.GeminiResponse(parseData)
// 				if err == nil {
// 					responseMsg = respCommand
// 				}
// 			}
// 			// SEND MESSAGE
// 			if h.erpContext.ThirdPartyServices["WA"] != nil {
// 				to := msg.Sender
// 				if waData.IsGroup {
// 					to = msg.SessionID
// 				}
// 				replyData := whatsmeow_client.WaMessage{
// 					JID:     msg.JID,
// 					Text:    responseMsg,
// 					To:      to,
// 					IsGroup: waData.IsGroup,
// 				}

// 				utils.LogJson(replyData)
// 				_, err := h.erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
// 				if err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}

// 				info := make(map[string]interface{})
// 				json.Unmarshal(infoByte, &info)

// 				info["Timestamp"] = time.Now().Format("2006-01-02T15:04:05-07:00")

// 				i, _ := json.Marshal(info)
// 				infoByte = i

// 				var waDataReply models.WhatsappMessageModel = models.WhatsappMessageModel{
// 					Sender:   splitSep[0],
// 					Receiver: msg.Sender,
// 					Message:  responseMsg,
// 					// MediaURL: mediaURLSaved,
// 					// MimeType: msg.MimeType,
// 					MessageInfo: info,
// 					Info:        string(infoByte),
// 					Session:     msg.SessionID,
// 					JID:         msg.JID,
// 					IsFromMe:    true,
// 					IsGroup:     msg.Info["IsGroup"].(bool),
// 				}
// 				err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
// 				if err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}

// 				msgWs := gin.H{
// 					"type":    "NEW_MESSAGE",
// 					"session": msg.SessionID,
// 					"data":    waDataReply,
// 				}

// 				msgWsByte, _ := json.Marshal(msgWs)
// 				services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 					return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 				})
// 			}

// 			// ADD HISTORY DATA
// 			historyData := objects.History{
// 				Sender: msg.Sender,
// 				Input:  msgContent,
// 				Output: resp,
// 			}

// 			b, err := json.Marshal(historyData)
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}
// 			err = services.REDIS.LPush(config.App.Redis.Key+"-"+msg.Sender, string(b)).Err()
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

//			}
//		}
//		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": msg})
//	}

// func sendWAMessage(erpContext *context.ERPContext, jid, to, message string) (any, error) {
// 	replyData := whatsmeow_client.WaMessage{
// 		JID:     jid,
// 		Text:    message,
// 		To:      to,
// 		IsGroup: false,
// 	}
// 	// utils.LogJson(replyData)
// 	return erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
// }
// func sendWAFileMessage(erpContext *context.ERPContext, jid, to, message, fileType, fileUrl string) (any, error) {
// 	replyData := whatsmeow_client.WaMessage{
// 		JID:      jid,
// 		Text:     message,
// 		To:       to,
// 		IsGroup:  false,
// 		FileType: fileType,
// 		FileUrl:  fileUrl,
// 	}
// 	// utils.LogJson(replyData)
// 	return erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
// }

type geminiResponse struct {
	Response string         `json:"response"`
	Type     string         `json:"type"`
	Command  string         `json:"command"`
	Params   map[string]any `json:"params"`
}
