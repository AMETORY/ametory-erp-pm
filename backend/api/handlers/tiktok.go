package handlers

import (
	"log"
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

type TiktokHandler struct {
	ctx *context.ERPContext
}

func NewTiktokHandler(ctx *context.ERPContext) *TiktokHandler {
	return &TiktokHandler{ctx: ctx}
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
