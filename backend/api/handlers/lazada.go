package handlers

import (
	"ametory-pm/services/app"
	"fmt"
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type LazadaHandler struct {
	ctx        *context.ERPContext
	appService *app.AppService
}

func NewLazadaHandler(ctx *context.ERPContext) *LazadaHandler {
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	return &LazadaHandler{
		ctx:        ctx,
		appService: appService,
	}
}

func (h *LazadaHandler) GetSessionsHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Get sessions"})
}

func (h *LazadaHandler) GetSessionDetailHandler(c *gin.Context) {
	sessionID := c.Param("sessionId")
	fmt.Println(sessionID)
	c.JSON(http.StatusOK, gin.H{"message": "Get session detail"})
}

func (h *LazadaHandler) SendMessageHandler(c *gin.Context) {
	sessionID := c.Param("sessionId")
	message := c.PostForm("message")
	if message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Message is empty"})
		return
	}
	fmt.Println(sessionID, message)
	c.JSON(http.StatusOK, gin.H{"message": "Send message successfully"})
}

func (h *LazadaHandler) SendFileHandler(c *gin.Context) {
	sessionID := c.Param("sessionId")
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is empty"})
		return
	}
	fmt.Println(file.Filename, sessionID)

	c.JSON(http.StatusOK, gin.H{"message": "Send file successfully"})
}

func (h *LazadaHandler) GetSessionMessagesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Get messages successfully"})
}

func (h *LazadaHandler) WebhookHandler(c *gin.Context) {
	var input map[string]any
	err := c.ShouldBind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	utils.LogJson(input)

	c.JSON(http.StatusOK, gin.H{"message": "Handle webhook successfully"})
}
