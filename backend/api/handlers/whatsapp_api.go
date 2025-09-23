package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"ametory-pm/services/app"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/shared/objects"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/ai_generator"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/meta"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type WhatsappApiHandler struct {
	erpContext                  *context.ERPContext
	contactService              *contact.ContactService
	appService                  *app.AppService
	metaService                 *meta.MetaService
	aiGeneratorService          *ai_generator.AiGeneratorService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
}

func NewWhatsappApiHandler(erpContext *context.ERPContext) *WhatsappApiHandler {
	contactService, ok := erpContext.ContactService.(*contact.ContactService)
	if !ok {
		panic("ContactService is not instance of contact.ContactService")
	}
	var appService *app.AppService
	appSrv, ok := erpContext.AppService.(*app.AppService)
	if ok {
		appService = appSrv
	}
	metaService, ok := erpContext.ThirdPartyServices["Meta"].(*meta.MetaService)
	if !ok {
		panic("MetaService is not instance of meta.MetaService")
	}
	aiGeneratorService, ok := erpContext.ThirdPartyServices["AiGenerator"].(*ai_generator.AiGeneratorService)
	if !ok {
		panic("aiGeneratorService is not instance of cache.CacheManager")
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := erpContext.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	return &WhatsappApiHandler{
		erpContext:                  erpContext,
		contactService:              contactService,
		appService:                  appService,
		metaService:                 metaService,
		aiGeneratorService:          aiGeneratorService,
		customerRelationshipService: customerRelationshipService,
	}
}

func (h *WhatsappApiHandler) runAutoPilot(phoneNumberID string, companyID *string, waMsg *models.WhatsappMessageModel) error {
	if waMsg == nil {
		return errors.New("message is nil")
	}
	now := time.Now()

	conn := connection.ConnectionModel{}
	err := h.erpContext.DB.Preload("AiAgent").First(&conn, "session = ?", phoneNumberID).Error
	if err != nil {
		return err
	}
	body := objects.MsgObject{
		Sender: waMsg.Sender,
	}
	// redisKey := "USER-DATA:" + conn.ID + ":" + body.Sender

	if conn.AiAgentID == nil {
		return nil
	}
	contact, err := h.contactService.GetContactByPhone(waMsg.Sender, *companyID)
	if err != nil {
		return err
	}
	var session mdl.WhatsappMessageSession
	err = h.erpContext.DB.Where("j_id = ?", phoneNumberID).First(&session).Error
	if err != nil {
		return err
	}
	autopilot := false

	// fmt.Println("AUTO RESPONSE TIME", conn.AutoResponseStartTime, conn.AutoResponseEndTime)
	if conn.AutoResponseStartTime != nil && conn.AutoResponseEndTime != nil {
		fmt.Println("AUTO RESPONSE TIME", *conn.AutoResponseStartTime, *conn.AutoResponseEndTime)
		autoResponseStartTime, err := time.ParseInLocation("2006-01-02 15:04", now.Format("2006-01-02")+" "+*conn.AutoResponseStartTime, time.Local)
		if err != nil {
			log.Println(err)
			return err
		}
		// fmt.Println("START TIME", now, autoResponseStartTime.Format("2006-01-02 15:04"))
		autoResponseEndTime, err := time.ParseInLocation("2006-01-02 15:04", now.Format("2006-01-02")+" "+*conn.AutoResponseEndTime, time.Local)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("BETWEEN", autoResponseStartTime.Format("2006-01-02 15:04"), "<", now.Format("2006-01-02 15:04"), ">", autoResponseEndTime.Format("2006-01-02 15:04"))
		if now.After(autoResponseStartTime) && now.Before(autoResponseEndTime) {
			autopilot = true
		}

	}

	if conn.AutoResponseStartTime == nil && conn.AutoResponseEndTime == nil {
		autopilot = true
	}

	fmt.Println("AUTO PILOT", autopilot)
	if session.IsHumanAgent {
		autopilot = false
	}
	// if isGroup {
	// 	autopilot = false

	// }
	fmt.Println("AUTO PILOT", autopilot)

	var replyResponse *models.WhatsappMessageModel
	if conn.AiAgent != nil && conn.IsAutoPilot && autopilot {
		generator, err := h.aiGeneratorService.GetGeneratorFromID(*conn.AiAgentID)
		if err != nil {
			return err
		}
		generator.SetApiKey(conn.AiAgent.ApiKey)
		var systemInstruction = fmt.Sprintf(`Sekarang Waktu menunjukkan: tgl: %s ,
jam: %s

%s
		
%s`, time.Now().Format("02-Jan-2006"), time.Now().Format("15:04:05"), conn.AiAgent.SystemInstruction, `

Tolong jawab dalam format JSON : 
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

`)
		// fmt.Println("SYSTEM INSTRUCTION", systemInstruction)
		// h.geminiService.SetUpSystemInstruction(systemInstruction)
		generator.SetSystemInstruction(systemInstruction)

		aiAgent, err := h.aiGeneratorService.GetAgent(*conn.AiAgentID)
		if err != nil {
			return err
		}

		isModel := true
		histories, _ := h.aiGeneratorService.GetHistories(&aiAgent.ID, conn.CompanyID, nil, &isModel, nil)
		var his []ai_generator.AiMessage = []ai_generator.AiMessage{}
		for _, v := range histories {

			his = append(his, ai_generator.AiMessage{
				Role:    "user",
				Content: v.Input,
			})

			his = append(his, ai_generator.AiMessage{
				Role:    "assistant",
				Content: v.Output,
			})

		}

		userHistories := []models.WhatsappMessageModel{}
		h.erpContext.DB.Model(&models.WhatsappMessageModel{}).Where("session = ?", body.Sender).Order("created_at desc").Limit(100).Find(&userHistories)

		userHistories = reverse(userHistories)

		for _, v := range userHistories {
			if v.IsFromMe {
				his = append(his, ai_generator.AiMessage{
					Role:    "user",
					Content: v.Message,
				})

			} else {
				his = append(his, ai_generator.AiMessage{
					Role:    "assistant",
					Content: v.Message,
				})

			}

		}

		output, err := generator.Generate(waMsg.Message, nil, his)
		if err != nil {
			fmt.Println("ERROR GENERATING CONTENT", err)
			return err
		}

		fmt.Println("OUTPUT", output)
		var response objects.AiResponse
		err = json.Unmarshal([]byte(output.Content), &response)
		if err != nil {
			fmt.Println("ERROR UNMARSHAL AI RESPONSE", err)
			return err
		}
		info := map[string]any{
			"Timestamp": time.Now().Format(time.RFC3339),
		}
		infoByte, err := json.Marshal(info)
		if err != nil {
			fmt.Println("ERROR MARSHAL INFO", err)
			return err
		}
		fmt.Println("SEND MESSAGE AUTO PILOT", body.JID, body.Sender, response.Response)

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

		history := models.AiAgentHistory{
			Input:       waMsg.Message,
			Output:      output.Content,
			AiAgentID:   &aiAgent.ID,
			SessionCode: &body.Sender,
		}
		// ADD HISTORY TO DB
		h.aiGeneratorService.CreateHistory(&history)
	}

	fmt.Println("SESSION ID #2", session.ID)
	if replyResponse != nil {
		if contact != nil {
			replyResponse.ContactID = &contact.ID
			replyResponse.CompanyID = contact.CompanyID
		}

		fmt.Println("SESSION ID #3", session.ID)
		fmt.Println("SESSION ID #3 response", replyResponse)

		err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(replyResponse)
		if err != nil {
			// log.Println(err)
			return err
		}

		replyResponse.SentAt = &now
		msg := gin.H{
			"message":    replyResponse.Message,
			"command":    "WHATSAPP_RECEIVED",
			"session_id": session.ID,
			"data":       replyResponse,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})

		now = time.Now()
		session.LastMessage = replyResponse.Message
		session.LastOnlineAt = &now
		h.erpContext.DB.Where("id = ?", session.ID).Model(&models.WhatsappMessageSession{}).Updates(&session)

	}

	return nil
}
func (h *WhatsappApiHandler) getMessageData(phoneNumberID string, waMsg *models.WhatsappMessageModel) error {
	if waMsg == nil {
		return errors.New("message is nil")
	}
	now := time.Now()
	waMsg.Session = fmt.Sprintf("%s@%s", waMsg.Sender, phoneNumberID)
	var session mdl.WhatsappMessageSession
	err := h.erpContext.DB.Where("session = ?", waMsg.Session).First(&session).Error
	if err != nil {
		return err
	}
	conn := connection.ConnectionModel{}
	err = h.erpContext.DB.First(&conn, "session = ?", phoneNumberID).Error
	if err != nil {
		return err
	}
	var contact models.ContactModel
	err = h.erpContext.DB.Where("phone = ?", waMsg.Sender).First(&contact).Error
	if err != nil {
		return err
	}

	session.LastMessage = waMsg.Message
	session.LastOnlineAt = &now

	h.erpContext.DB.Save(&session)

	// fmt.Println("GET MESSAGE")
	// utils.LogJson(waMsg)
	if err := h.erpContext.DB.Create(&waMsg).Error; err != nil {
		fmt.Println("ERROR CREATE WHATSAPP MESSAGE #2", err)
		return err
	}

	msg := gin.H{
		"message":    waMsg.Message,
		"command":    "WHATSAPP_RECEIVED",
		"session_id": session.ID,
		"data":       waMsg,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	return nil
}
func (h *WhatsappApiHandler) getContact(phoneNumber string, displayName string, companyID *string) (*models.ContactModel, error) {
	contact, err := h.contactService.GetContactByPhone(phoneNumber, *companyID)
	if err != nil {
		if err.Error() == "contact not found" {
			contact = &mdl.ContactModel{
				CompanyID:  companyID,
				Name:       displayName,
				Phone:      &phoneNumber,
				IsCustomer: true,
			}
			if err := h.erpContext.DB.Create(contact).Error; err != nil {
				fmt.Println("ERROR CREATE CONTACT", err)
				return nil, err
			}
		} else {
			fmt.Println("ERROR GET CONTACT BY PHONE NUMBER ID", err)
			return nil, err
		}
	}
	return contact, nil
}
func (h *WhatsappApiHandler) getSession(phoneNumberID string, phoneNumber string, displayName string, lastMessage string, companyID *string) (*objects.WhatsappApiSession, error) {
	conn := connection.ConnectionModel{}
	now := time.Now()
	err := h.erpContext.DB.First(&conn, "session = ?", phoneNumberID).Error
	if err != nil {
		return nil, err
	}
	var session mdl.WhatsappMessageSession
	sessionName := fmt.Sprintf("%s@%s", phoneNumber, phoneNumberID)
	var sessionContact *models.ContactModel
	err = h.erpContext.DB.Where("session = ?", sessionName).First(&session).Error
	if err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			var contactID *string
			contact, err := h.contactService.GetContactByPhone(phoneNumber, *conn.CompanyID)
			if err != nil {
				if err.Error() == "contact not found" {
					contact = &mdl.ContactModel{
						CompanyID:  companyID,
						Name:       displayName,
						Phone:      &phoneNumber,
						IsCustomer: true,
					}
					if err := h.erpContext.DB.Create(contact).Error; err != nil {
						fmt.Println("ERROR CREATE CONTACT", err)
						return nil, err
					}
				} else {
					fmt.Println("ERROR GET CONTACT BY PHONE NUMBER ID", err)
					return nil, err
				}
			}
			refType := "connection"
			sessionContact = contact

			session = mdl.WhatsappMessageSession{
				Session:      sessionName,
				CompanyID:    conn.CompanyID,
				JID:          phoneNumberID,
				SessionName:  phoneNumber,
				RefID:        &conn.ID,
				RefType:      &refType,
				LastOnlineAt: &now,
				LastMessage:  lastMessage,
				ContactID:    contactID,
				Contact:      sessionContact,
			}
			if err := h.erpContext.DB.Create(&session).Error; err != nil {
				fmt.Println("ERROR CREATE WHATSAPP MESSAGE SESSION", err)
				return nil, err
			}
		} else {
			fmt.Println("ERROR GET WHATSAPP MESSAGE SESSION BY SESSION", err)
			return nil, err
		}
	}

	var resp objects.WhatsappApiSession
	resp.PhoneNumberID = phoneNumberID
	resp.Session = conn.Session
	resp.AccessToken = conn.AccessToken
	if conn.CompanyID != nil {
		resp.CompanyID = *conn.CompanyID
	}
	return &resp, nil
}

func (h *WhatsappApiHandler) GetAllMessageTemplates(c *gin.Context) {
	wabaID := c.Params.ByName("waba_id")

	conn := connection.ConnectionModel{}
	err := h.erpContext.DB.First(&conn, "session_name = ?", wabaID).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.metaService.WhatsappApiService.SetAccessToken(&conn.AccessToken)
	templates, err := h.metaService.WhatsappApiService.GetAllMessageTemplates(wabaID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, templates)
}

func (h *WhatsappApiHandler) GetMessageTemplateByName(c *gin.Context) {
	wabaID := c.Params.ByName("waba_id")
	templateID := c.Params.ByName("template_id")
	conn := connection.ConnectionModel{}
	err := h.erpContext.DB.First(&conn, "session_name = ?", wabaID).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	h.metaService.WhatsappApiService.SetAccessToken(&conn.AccessToken)
	resp, err := h.metaService.WhatsappApiService.GetMessageTemplateByName(wabaID, templateID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, resp)
}

func (h *WhatsappApiHandler) WhatsappApiWebhookHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		resp, err := h.metaService.VerifyFacebook(c.Request, config.App.Facebook.FacebookVerifyToken)
		if err == nil {
			c.String(http.StatusOK, resp)
		} else {
			c.String(http.StatusUnauthorized, "Invalid verify token")
		}
		return
	}

	// var bodyMap map[string]interface{}
	// if err := c.BindJSON(&bodyMap); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
	// 	return
	// }

	// fmt.Println("BODY WEBHOOK")
	// utils.LogJson(bodyMap)

	var data objects.WhatsappApiWebhookRequest
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body #2"})
		return
	}
	// fmt.Println("DATA WEBHOOK")
	if config.App.Server.Debug {
		utils.LogJson(data)
	}
	// now := time.Now()
	h.metaService.WhatsappApiService.WhatsappApiWebhook(c.Request, data, "", h.getContact, h.getSession, h.getMessageData, h.runAutoPilot, nil)
	// fmt.Println("DATA WEBHOOK")
	// utils.LogJson(data)

	// var conn connection.ConnectionModel
	// for _, entry := range data.Entry {
	// 	for _, change := range entry.Changes {
	// 		if change.Field == "messages" && change.Value.MessagingProduct == "whatsapp" {
	// 			if len(change.Value.Contacts) > 0 {
	// 				phoneNumberID := change.Value.Metadata.PhoneNumberID
	// 				fmt.Println("PHONE NUMBER ID", phoneNumberID)
	// 				err := h.erpContext.DB.First(&conn, "session = ?", phoneNumberID).Error
	// 				if err != nil {
	// 					fmt.Println("ERROR GET CONNECTION BY PHONE NUMBER ID", err)
	// 					continue
	// 				}

	// 				var session mdl.WhatsappMessageSession
	// 				err = h.erpContext.DB.Where("j_id = ?", phoneNumberID).First(&session).Error
	// 				if err != nil {
	// 					fmt.Println("ERROR GET CONNECTION BY PHONE NUMBER ID", err)
	// 					continue
	// 				}

	// 				// GET CONNECTION BY PHONE NUMBER ID
	// 				contact, err := h.contactService.GetContactByPhone(change.Value.Contacts[0].WAID, *conn.CompanyID)
	// 				if err != nil {
	// 					if errors.Is(err, gorm.ErrRecordNotFound) {
	// 						contact = &mdl.ContactModel{
	// 							CompanyID: conn.CompanyID,
	// 							Name:      change.Value.Contacts[0].Profile.Name,
	// 							Phone:     &change.Value.Contacts[0].WAID,
	// 						}
	// 						if err := h.erpContext.DB.Create(contact).Error; err != nil {
	// 							fmt.Println("ERROR CREATE CONTACT", err)
	// 							continue
	// 						}
	// 					} else {
	// 						fmt.Println("ERROR GET CONTACT BY PHONE NUMBER ID", err)
	// 						continue
	// 					}
	// 				}

	// 				// GET CONTACT BY PHONE NUMBER ID
	// 				fmt.Println("GET CONTACT")
	// 				utils.SaveJson(contact)

	// 				// GET MESSAGE
	// 				for _, msg := range change.Value.Messages {
	// 					message := ""
	// 					// QUOTE MESSAGE

	// 					if msg.Type == "text" {
	// 						message = msg.Text.Body
	// 					}
	// 					if msg.Type == "image" && msg.Image != nil {
	// 						message = msg.Image.Caption
	// 						path, err := GetFacebookMediaObject(msg.Image.ID, conn.AccessToken)
	// 						if err != nil {
	// 							fmt.Println("ERROR", err)
	// 						}
	// 						fmt.Println("PATH", *path)
	// 					}
	// 					sessionWa := fmt.Sprintf("%s@%s", *contact.Phone, conn.Session)
	// 					waMsg := mdl.WhatsappMessageModel{
	// 						Message:   message,
	// 						MessageID: &msg.ID,
	// 						Sender:    msg.From,
	// 						JID:       phoneNumberID,
	// 						Contact:   contact,
	// 						SentAt:    &now,
	// 						Session:   sessionWa,
	// 						CompanyID: conn.CompanyID,
	// 					}

	// 					if msg.Context != nil {
	// 						waMsg.QuotedMessageID = &msg.Context.ID
	// 					}

	// 					// utils.LogJson(waMsg)

	// 					if err := h.erpContext.DB.Create(&waMsg).Error; err != nil {
	// 						fmt.Println("ERROR CREATE WHATSAPP MESSAGE #2", err)
	// 						continue
	// 					}

	// 					msg := gin.H{
	// 						"message":    message,
	// 						"command":    "WHATSAPP_RECEIVED",
	// 						"session_id": session.ID,
	// 						"data":       waMsg,
	// 					}
	// 					b, _ := json.Marshal(msg)
	// 					h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
	// 						url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
	// 						return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// 					})

	// 				}

	// 			}
	// 		}

	// 	}

	// }

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetFacebookMediaObject(mediaID, accessToken string) (*string, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v21.0/%s", mediaID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media object, status code: %d", resp.StatusCode)
	}

	var mediaObject whatsapp.FacebookMedia
	if err := json.NewDecoder(resp.Body).Decode(&mediaObject); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	req, err = http.NewRequest("GET", mediaObject.URL, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("URL", mediaObject.URL)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media content, status code: %d", resp.StatusCode)
	}

	extension := ""
	switch mediaObject.MimeType {
	case "image/jpeg":
		extension = ".jpg"
	case "image/png":
		extension = ".png"
	default:
		return nil, fmt.Errorf("unsupported media type: %s", mediaObject.MimeType)
	}

	fileName := fmt.Sprintf("%s%s", mediaObject.ID, extension)
	return downloadAndSaveMedia(mediaObject.URL, fileName)
}

func downloadAndSaveMedia(url, fileName string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media, status code: %d", resp.StatusCode)
	}

	assetsFolder := "./assets/images/"
	if _, err := os.Stat(assetsFolder); os.IsNotExist(err) {
		os.Mkdir(assetsFolder, os.ModePerm)
	}

	destination := filepath.Join(assetsFolder, fileName)
	if err := saveResponseBodyToFile(resp.Body, destination); err != nil {
		return nil, err
	}

	if config.App.Server.StorageProvider == "google" {
		// destination, err = HandleGoogleStorage(destination)
		// if err != nil {
		// 	return nil, err
		// }
	}

	return &destination, nil
}

func saveResponseBodyToFile(respBody io.Reader, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, respBody)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}
