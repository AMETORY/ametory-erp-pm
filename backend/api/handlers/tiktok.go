package handlers

import (
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	customer_service_v202309 "tiktokshop/open/sdk_golang/models/customer_service/v202309"
	tiktok "tiktokshop/open/sdk_golang/service"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type TiktokHandler struct {
	ctx           *context.ERPContext
	tiktokService *tiktok.TiktokService
	appService    *app.AppService
}

func NewTiktokHandler(ctx *context.ERPContext) *TiktokHandler {
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	tiktokService, ok := ctx.ThirdPartyServices["Tiktok"].(*tiktok.TiktokService)
	if !ok {
		panic("ThirdPartyServices is not instance of tiktok.TiktokService")
	}
	return &TiktokHandler{
		ctx:           ctx,
		appService:    appService,
		tiktokService: tiktokService,
	}
}

func (h *TiktokHandler) WebhookHandler(c *gin.Context) {
	var requestData map[string]any
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println("INCOMING TIKTOK WEBHOOK")
	utils.LogJson(requestData)

	shopId, ok := requestData["shop_id"].(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid shop_id"})
		return
	}

	var connection connection.ConnectionModel
	err := h.ctx.DB.Where("username = ?", shopId).First(&connection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if connection.Type != "tiktok" || connection.Status != "ACTIVE" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not active"})
		return
	}

	webhookType, ok := requestData["type"].(int)
	if ok {
		// NEW CONVERSATION
		if webhookType == 13 {
			msgNotif := gin.H{
				"command": "TIKTOK_NEW_CONVERSATION",
			}
			msgNotifStr, _ := json.Marshal(msgNotif)
			h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
				url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *connection.CompanyID)
				return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
			})

		}
		if webhookType == 14 {
			data := requestData["data"].(map[string]any)
			dataByte, _ := json.Marshal(data)

			var msgData models.TiktokMessage
			if err := json.Unmarshal(dataByte, &msgData); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			var responseMessage customer_service_v202309.CustomerService202309GetConversationMessagesResponseDataMessages
			responseMessage.Id = &msgData.MessageID
			msgContent := msgData.Content

			responseMessage.CreateTime = msgData.CreateTime
			responseMessage.Type = &msgData.Type
			responseMessage.Content = &msgData.Content
			responseMessage.Sender = &customer_service_v202309.CustomerService202309GetConversationMessagesResponseDataMessagesSender{
				Nickname: &msgData.Sender.ImUserID,
				Role:     &msgData.Sender.Role,
			}

			msgNotif := gin.H{
				"message":    msgContent,
				"command":    "TIKTOK_RECEIVED",
				"session_id": msgData.ConversationID,
				"data":       responseMessage,
			}
			msgNotifStr, _ := json.Marshal(msgNotif)
			h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
				url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *connection.CompanyID)
				return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
			})
		}
	}

	// Use requestData for further processing

	log.Println("Received TikTok webhook data:", requestData)

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *TiktokHandler) GetSessionsHandler(c *gin.Context) {
	connectionID := c.Query("connection_id")
	nextPageToken := c.Query("next_page_token")

	connection, err := h.appService.ConnectionService.GetConnection(connectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if connection.Type != "tiktok" || connection.Status != "ACTIVE" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not active"})
		return
	}

	if connection.AuthData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not authorized"})
		return
	}

	authData := map[string]any{}
	if err := json.Unmarshal([]byte(*connection.AuthData), &authData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.tiktokService.CustomerService202309GetConversationsGet(authData["access_token"].(string), connection.Password, nextPageToken, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")
	for _, v := range resp.Conversations {
		var sessionName string
		var participant *customer_service_v202309.CustomerService202309GetConversationsResponseDataConversationsParticipants
		for _, v := range v.Participants {
			if *v.Role == "BUYER" {
				participant = &v
				sessionName = *participant.Nickname
				break
			}
		}

		var partData *json.RawMessage
		refType := "connection"
		if participant != nil {
			b, _ := json.Marshal(participant)
			p := json.RawMessage(b)
			partData = &p

		}
		var createTime, lastMsgTime time.Time
		if v.CreateTime != nil {
			createTime = time.Unix(*v.CreateTime, 0)
		}
		if v.LatestMessage != nil {
			lastMsgTime = time.Unix(*v.LatestMessage.CreateTime, 0)
		}

		var lastMsg json.RawMessage
		if v.LatestMessage != nil && v.LatestMessage.Content != nil {
			b := json.RawMessage([]byte(*v.LatestMessage.Content))
			lastMsg = b
		}

		dataSession := models.TiktokMessageSession{
			RefID:       &connectionID,
			RefType:     &refType,
			Session:     *v.Id,
			SessionName: sessionName,
			CompanyID:   &companyID,
			Participant: partData,
			CreateTime:  &createTime,
			LastMsgTime: &lastMsgTime,
			LastMessage: &lastMsg,
		}
		err := h.ctx.DB.Where("session = ?", *v.Id).First(&dataSession).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			h.ctx.DB.Create(&dataSession)
		} else {
			h.ctx.DB.Model(&dataSession).Where("session = ?", *v.Id).Updates(map[string]interface{}{
				"session_name":  sessionName,
				"participant":   partData,
				"last_msg_time": &lastMsgTime,
				"last_message":  &lastMsg,
			})
		}

	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (h *TiktokHandler) SendFileHandler(c *gin.Context) {
	conversationID := c.Param("sessionId")
	connectionId := c.Query("connection_id")

	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session Id is required"})
		return
	}
	if connectionId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Connection Id is required"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save file to temporary directory
	filePath := fmt.Sprintf("%s/%s%s", os.TempDir(), uuid.New(), filepath.Ext(file.Filename))
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileData, err := os.Open(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer fileData.Close()
	defer os.Remove(filePath)

	connection, err := h.appService.ConnectionService.GetConnection(connectionId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if connection.Type != "tiktok" || connection.Status != "ACTIVE" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not active"})
		return
	}

	if connection.AuthData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not authorized"})
		return
	}

	authData := map[string]any{}
	if err := json.Unmarshal([]byte(*connection.AuthData), &authData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.tiktokService.CustomerService202309UploadBuyerMessagesImagePost(authData["access_token"].(string), connection.Password, conversationID, fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
func (h *TiktokHandler) SendMessageHandler(c *gin.Context) {
	conversationID := c.Param("sessionId")

	if conversationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Session Id is required"})
		return
	}

	var input map[string]any
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	connection, err := h.appService.ConnectionService.GetConnection(input["connection_id"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if connection.Type != "tiktok" || connection.Status != "ACTIVE" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not active"})
		return
	}

	if connection.AuthData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not authorized"})
		return
	}

	authData := map[string]any{}
	if err := json.Unmarshal([]byte(*connection.AuthData), &authData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.tiktokService.CustomerService202309SendMessagePost(authData["access_token"].(string), connection.Password, conversationID, input["type"].(string), input["content"].(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responseMessage customer_service_v202309.CustomerService202309GetConversationMessagesResponseDataMessages
	responseMessage.Id = resp.MessageId
	msgType := input["type"].(string)
	msgContent := input["content"].(string)
	nickname := input["nickname"].(string)
	role := input["role"].(string)

	now := time.Now().Unix()
	responseMessage.CreateTime = &now
	responseMessage.Type = &msgType
	responseMessage.Content = &msgContent
	responseMessage.Sender = &customer_service_v202309.CustomerService202309GetConversationMessagesResponseDataMessagesSender{
		Nickname: &nickname,
		Role:     &role,
	}

	msgNotif := gin.H{
		"message":    msgContent,
		"command":    "TIKTOK_RECEIVED",
		"session_id": conversationID,
		"data":       responseMessage,
	}
	msgNotifStr, _ := json.Marshal(msgNotif)
	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *connection.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(http.StatusOK, gin.H{"data": responseMessage})
}
func (h *TiktokHandler) GetSessionDetailHandler(c *gin.Context) {
	conversationID := c.Param("sessionId")

	dataSession := models.TiktokMessageSession{}

	err := h.ctx.DB.Where("session = ?", conversationID).First(&dataSession).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Session not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": dataSession})
}
func (h *TiktokHandler) GetSessionMessagesHandler(c *gin.Context) {
	connectionID := c.Query("connection_id")
	nextPageToken := c.Query("next_page_token")
	conversationID := c.Param("sessionId")

	connection, err := h.appService.ConnectionService.GetConnection(connectionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if connection.Type != "tiktok" || connection.Status != "ACTIVE" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not active"})
		return
	}

	if connection.AuthData == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Connection is not authorized"})
		return
	}

	authData := map[string]any{}
	if err := json.Unmarshal([]byte(*connection.AuthData), &authData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.tiktokService.CustomerService202309GetConversationMessagesGet(authData["access_token"].(string), connection.Password, conversationID, nextPageToken, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})

}

// func (h *TiktokHandler) handleTiktokMessage(c *gin.Context, conn *models.Connection) error {
// 	data, err := io.ReadAll(c.Request.Body)
// 	if err != nil {
// 		return err
// 	}

// 	var tiktokMessage models.TiktokMessage
// 	err = json.Unmarshal(data, &tiktokMessage)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = h.ctx.CustomerRelationshipService.CreateTiktokMessage(conn, tiktokMessage)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
