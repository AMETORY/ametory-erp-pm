package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"log"
	"net/http"
	tiktok "tiktokshop/open/sdk_golang/service"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
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
	var requestData map[string]interface{}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
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

	c.JSON(http.StatusOK, gin.H{"data": resp})
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
