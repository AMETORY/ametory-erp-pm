package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship/telegram"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type TelegramHandler struct {
	erpContext                  *context.ERPContext
	appService                  *app.AppService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	geminiService               *google.GeminiService
	pmService                   *project_management.ProjectManagementService
	contactSrv                  *contact.ContactService
}

func NewTelegramHandler(erpContext *context.ERPContext) *TelegramHandler {

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
	contactSrv, ok := erpContext.ContactService.(*contact.ContactService)
	if !ok {
		panic("ContactService is not instance of contact.ContactService")
	}
	return &TelegramHandler{
		erpContext:                  erpContext,
		appService:                  appService,
		customerRelationshipService: customerRelationshipService,
		geminiService:               geminiService,
		pmService:                   pmService,
		contactSrv:                  contactSrv,
	}
}

func (h *TelegramHandler) GetSessionsHandler(c *gin.Context) {
	sessionId := c.Query("session_id")
	sessionName := ""
	var session *models.TelegramMessageSession
	if sessionId != "" {
		err := h.erpContext.DB.First(&session, "id = ?", sessionId).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		sessionName = session.Session
	}
	fmt.Println("sessionName", sessionName)
	sessions, err := h.customerRelationshipService.TelegramService.GetSessionMessageBySessionName(sessionName, *c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	telegramSessions := sessions.Items.(*[]models.TelegramMessageSession)
	newTelegramSessions := []models.TelegramMessageSession{}
	for _, v := range *telegramSessions {
		if v.Session != "" {
			var conn connection.ConnectionModel
			err = h.erpContext.DB.Select("id, session_name, name, color").First(&conn, "id = ?", v.Session).Error
			if err == nil {
				v.Ref = &conn
				v.RefID = &conn.ID
			}
		}

		profile, err := v.Contact.GetProfilePicture(h.erpContext.DB)
		if err == nil {
			v.Contact.ProfilePicture = profile

		}

		newTelegramSessions = append(newTelegramSessions, v)
	}
	sessions.Items = newTelegramSessions
	c.JSON(200, gin.H{"data": sessions})
}

func (h *TelegramHandler) GetSessionMessagesHandler(c *gin.Context) {
	sessionId := c.Query("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session *models.TelegramMessageSession
	err := h.erpContext.DB.First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	messages, err := h.customerRelationshipService.TelegramService.GetMessageSessionChatBySessionName(session.ID, session.ContactID, *c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messages.Items = reverseTelegram(*messages.Items.(*[]models.TelegramMessage))

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": messages})
}

func (h *TelegramHandler) GetSessionDetailHandler(c *gin.Context) {
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session models.TelegramMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var connection connection.ConnectionModel
	err = h.erpContext.DB.First(&connection, "id = ?", session.Session).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	session.Contact.ProfilePicture, _ = session.Contact.GetProfilePicture(h.erpContext.DB)

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": session, "connection": connection})
}

func (h *TelegramHandler) TelegramClearSessionHandler(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	var session models.TelegramMessageSession
	err := h.erpContext.DB.First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = h.erpContext.DB.Unscoped().Where("session = ?", session.Session).Delete(&models.TelegramMessage{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session.LastMessage = ""
	session.LastOnlineAt = nil

	if err := h.erpContext.DB.Save(&session).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"command":    "TELEGRAM_CLEAR_MESSAGE",
		"session_id": sessionId,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(http.StatusOK, gin.H{"message": "Session cleared successfully"})
}

func (h *TelegramHandler) TelegramDeleteSessionHandler(c *gin.Context) {
	sessionId := c.Param("session_id")
	if sessionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session_id is required"})
		return
	}

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session models.TelegramMessageSession
	err := h.erpContext.DB.First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	err = h.erpContext.DB.Unscoped().Where("session = ?", session.Session).Delete(&models.TelegramMessage{}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.erpContext.DB.Unscoped().Delete(&models.TelegramMessageSession{}, "id = ?", sessionId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Session deleted successfully"})
}

func (h *TelegramHandler) SetUpWebHookHandler(c *gin.Context) {
	connectionID := c.Param("connectionID")
	if connectionID == "" {
		log.Println("connectionID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "connectionID is empty"})
		return
	}
	var connection connection.ConnectionModel
	err := h.erpContext.DB.First(&connection, "id = ?", connectionID).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	url := fmt.Sprintf("%s/api/v1/telegram/webhook?connection_id=%s", config.App.Server.BaseURL, connectionID)
	// fmt.Println("URL", url)
	log.Println("SET WEBHOOK", url)
	h.customerRelationshipService.TelegramService.SetToken(&connection.SessionName, &connection.AccessToken)
	err = h.customerRelationshipService.TelegramService.SetWebhook(url)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "ok"})
}

func (h *TelegramHandler) SendMessage(c *gin.Context) {
	input := struct {
		Message     string             `json:"message"`
		FileURL     string             `json:"file_url"`
		FileCaption string             `json:"file_caption"`
		Files       []models.FileModel `json:"files"`
	}{}
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	sessionID := c.Param("session_id")
	if sessionID == "" {
		log.Println("connectionID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "connectionID is empty"})
		return
	}
	var session models.TelegramMessageSession
	err := h.erpContext.DB.Preload("Contact").First(&session, "id = ?", sessionID).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var connection connection.ConnectionModel
	err = h.erpContext.DB.First(&connection, "id = ?", session.Session).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.customerRelationshipService.TelegramService.SetToken(&connection.SessionName, &connection.AccessToken)

	chatID, err := strconv.Atoi(*session.Contact.TelegramID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mimeType := ""
	if input.FileURL != "" {
		mimeType = utils.GetMimeType(input.FileURL)
	}
	userID := c.MustGet("userID").(string)
	memberID := c.MustGet("memberID").(string)

	dataReply := models.TelegramMessage{
		ContactID:                session.ContactID,
		Message:                  input.Message,
		CompanyID:                session.CompanyID,
		Session:                  session.Session,
		IsFromMe:                 true,
		TelegramMessageSessionID: &session.ID,
		UserID:                   &userID,
		MemberID:                 &memberID,
		MediaURL:                 input.FileURL,
		MimeType:                 mimeType,
	}
	if input.FileCaption != "" {
		dataReply.Message = input.FileCaption
	}

	telegramMsg := telegram.TelegramMsg{
		ChatID:      int64(chatID),
		Message:     input.Message,
		FileURL:     input.FileURL,
		FileCaption: input.FileCaption,
		MimeType:    mimeType,
		Data:        &dataReply,
		Save:        true,
	}

	h.customerRelationshipService.TelegramService.SetInput(&telegramMsg)
	_, err = customer_relationship.SendCustomerServiceMessage(h.customerRelationshipService.TelegramService)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msgNotif := gin.H{
		"message":    input.Message,
		"command":    "TELEGRAM_RECEIVED",
		"session_id": session.ID,
		"data":       dataReply,
	}
	msgNotifStr, _ := json.Marshal(msgNotif)
	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	c.JSON(200, gin.H{"message": "ok"})
}

func (h *TelegramHandler) WebhookHandler(c *gin.Context) {
	now := time.Now()
	connectionID := c.Query("connection_id")
	if connectionID == "" {
		log.Println("connectionID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "connectionID is empty"})
		return
	}
	var connection connection.ConnectionModel
	err := h.erpContext.DB.First(&connection, "id = ?", connectionID).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reqBody, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Telegram request body: %s", string(reqBody))

	tgResponse := models.TGResponse{}
	if err := json.Unmarshal(reqBody, &tgResponse); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// utils.LogJson(tgResponse)

	fromID := strconv.Itoa(int(tgResponse.Message.From.ID))
	if fromID == "" {
		log.Println("fromID is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "fromID is empty"})
		return
	}

	h.customerRelationshipService.TelegramService.SetToken(&connection.SessionName, &connection.AccessToken)

	var contact models.ContactModel
	err = h.erpContext.DB.First(&contact, "telegram_id = ?", fromID).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			connType := "telegram"
			contact = models.ContactModel{
				TelegramID:     &fromID,
				Name:           fmt.Sprintf("%s %s", tgResponse.Message.From.FirstName, tgResponse.Message.From.LastName),
				ConnectionType: &connType,
				CompanyID:      connection.CompanyID,
			}
			if err := h.contactSrv.CreateContact(&contact); err != nil {
				log.Println(err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	session, err := h.customerRelationshipService.TelegramService.CheckSession(&tgResponse, &contact, connectionID, *connection.CompanyID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	msgID := strconv.Itoa(int(tgResponse.Message.MessageID))
	telegramData := models.TelegramMessage{
		ContactID:                session.ContactID,
		Message:                  tgResponse.Message.Text,
		CompanyID:                session.CompanyID,
		MessageID:                &msgID,
		Session:                  session.Session,
		TelegramMessageSessionID: &session.ID,
	}
	if len(tgResponse.Message.Photos) > 0 {
		photo := tgResponse.Message.Photos[len(tgResponse.Message.Photos)-1]
		telegramData.Message = tgResponse.Message.Caption
		mediaUrl, mimeType, err := h.saveFile(connection, photo.FileID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		telegramData.MediaURL = mediaUrl
		telegramData.MimeType = mimeType

	}

	if tgResponse.Message.Video != nil {
		telegramData.Message = tgResponse.Message.Caption
		mediaUrl, mimeType, err := h.saveFile(connection, tgResponse.Message.Video.FileID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		telegramData.MediaURL = mediaUrl
		telegramData.MimeType = mimeType
	}
	if tgResponse.Message.Voice != nil {
		telegramData.Message = tgResponse.Message.Caption
		mediaUrl, mimeType, err := h.saveFile(connection, tgResponse.Message.Voice.FileID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		telegramData.MediaURL = mediaUrl
		telegramData.MimeType = mimeType
	}
	if tgResponse.Message.Document != nil {
		telegramData.Message = tgResponse.Message.Caption
		mediaUrl, mimeType, err := h.saveFile(connection, tgResponse.Message.Document.FileID)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		telegramData.MediaURL = mediaUrl
		telegramData.MimeType = mimeType
	}

	if tgResponse.Message.ReplyToMessage != nil {
		telegramData.QuotedMessage = &tgResponse.Message.ReplyToMessage.Text
		msgID := fmt.Sprintf("%d", tgResponse.Message.ReplyToMessage.MessageID)
		telegramData.QuotedMessageID = &msgID
	}

	err = h.customerRelationshipService.TelegramService.SaveMessage(&telegramData)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// GET PROFILE PICTURE
	profilePic, _ := contact.GetProfilePicture(h.erpContext.DB)
	if profilePic == nil {
		resp, err := h.customerRelationshipService.TelegramService.GetUserProfilePhotos(tgResponse.Message.Chat.ID)
		if err == nil {
			utils.LogJson(resp)
			fileUrl := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", connection.AccessToken, resp["result"].(map[string]any)["file_path"].(string))
			path, mimeType, err := saveFileContenFromUrl(fileUrl)
			if err == nil {
				fileName := filepath.Base(path)
				mediaURL := fmt.Sprintf("%s/%s", h.appService.Config.Server.BaseURL, path)
				h.erpContext.DB.Create(&models.FileModel{
					FileName: fileName,
					Path:     path,
					URL:      mediaURL,
					MimeType: mimeType,
					RefID:    contact.ID,
					RefType:  "contact",
				})
			}

		}

	}

	telegramData.Contact = &contact
	telegramData.SentAt = &now
	msg := gin.H{
		"message":    tgResponse.Message.Text,
		"command":    "TELEGRAM_RECEIVED",
		"session_id": session.ID,
		"data":       telegramData,
	}

	fmt.Println(msg)
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
		fmt.Println("URL", url)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(200, gin.H{"message": "ok"})
}

func (h *TelegramHandler) saveFile(connection connection.ConnectionModel, fileID string) (string, string, error) {
	resp, err := h.customerRelationshipService.TelegramService.GetFile(fileID)
	if err != nil {
		return "", "", err

	}
	fileUrl := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", connection.AccessToken, resp["result"].(map[string]any)["file_path"].(string))
	path, mimeType, err := saveFileContenFromUrl(fileUrl)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	mediaURL := fmt.Sprintf("%s/%s", h.appService.Config.Server.BaseURL, path)
	return mediaURL, mimeType, nil
}
func reverseTelegram(messages []models.TelegramMessage) []models.TelegramMessage {
	for i, j := 0, len(messages)-1; i < len(messages)/2; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}
