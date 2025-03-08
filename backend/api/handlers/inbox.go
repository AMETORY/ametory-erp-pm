package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type InboxHandler struct {
	ctx            *context.ERPContext
	messageService *message.MessageService
	appService     *app.AppService
}

func NewInboxHandler(ctx *context.ERPContext) *InboxHandler {
	messageService, ok := ctx.MessageService.(*message.MessageService)
	if !ok {
		panic("MessageService is not instance of message.MessageService")
	}

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}

	return &InboxHandler{
		ctx:            ctx,
		messageService: messageService,
		appService:     appService,
	}
}

func (h *InboxHandler) SendMessageHandler(c *gin.Context) {
	var input models.InboxMessageModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	memberID := c.MustGet("memberID").(string)

	input.SenderMemberID = &memberID
	err := h.messageService.InboxService.SendMessage(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, v := range input.Attachments {
		v.RefID = input.ID
		v.RefType = "inbox"
		h.ctx.DB.Save(&v)

	}

	if input.ParentInboxMessageID != nil {
		var parent models.InboxMessageModel
		h.ctx.DB.First(&parent, "id = ?", *input.ParentInboxMessageID)
		if parent.RecipientMemberID != &memberID {
			h.ctx.DB.Model(&parent).Where("id = ?", parent.ID).Update("read", false)
		}

		replies, err := parent.LoadRecursiveChildren(h.ctx.DB)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		lastReplies := replies[len(replies)-1]
		if lastReplies.RecipientMemberID != &memberID {
			h.ctx.DB.Model(&lastReplies).Where("id = ?", lastReplies.ID).Update("read", false)
		}
	} else {
		msg := gin.H{
			"message_id":   input.ID,
			"subject":      input.Subject,
			"inbox_id":     *input.InboxID,
			"message":      "Message sent",
			"sender_id":    c.MustGet("userID").(string),
			"recipient_id": *input.RecipientMemberID,
		}
		// utils.LogJson(msg)
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.Request.Header.Get("ID-Company"))
			fmt.Println(q.Request.URL.Path, url)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}

	c.JSON(200, gin.H{"message": "Message sent"})
}

func (h *InboxHandler) GetInboxesHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	userID := c.MustGet("userID").(string)
	data, err := h.messageService.InboxService.GetInboxes(&userID, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Inbox retrieved successfully"})
}

func (h *InboxHandler) CountUnreadHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	data, err := h.messageService.InboxService.CountUnread(nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Unread messages count retrieved successfully"})
}
func (h *InboxHandler) CountUnreadSentHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	data, err := h.messageService.InboxService.CountUnreadSendMessage(nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Unread messages count retrieved successfully"})
}

func (h *InboxHandler) DeleteMessageHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	userID := c.MustGet("userID").(string)
	messageID := c.Param("id")
	err := h.messageService.InboxService.DeleteMessage(messageID, &userID, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Message deleted"})
}

func (h *InboxHandler) SentMessageHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	data, err := h.messageService.InboxService.GetSentMessages(*c.Request, c.Query("search"), nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Messages retrieved successfully"})
}
func (h *InboxHandler) GetMessagesHandler(c *gin.Context) {
	// memberID := c.MustGet("memberID").(string)
	var inboxID *string
	if c.Query("inbox_id") != "" {
		_inboxID := c.Query("inbox_id")
		inboxID = &_inboxID
	}
	data, err := h.messageService.InboxService.GetMessageByInboxID(*c.Request, c.Query("search"), inboxID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Messages retrieved successfully"})
}

func (h *InboxHandler) GetMessagesDetailHandler(c *gin.Context) {
	messageID := c.Param("id")
	data, err := h.messageService.InboxService.GetInboxMessageDetail(messageID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	data.Read = true
	h.ctx.DB.Save(&data)

	for _, v := range data.Replies {
		v.Read = true
		h.ctx.DB.Save(&v)
	}

	c.JSON(200, gin.H{"data": data, "message": "Message retrieved successfully"})
}
