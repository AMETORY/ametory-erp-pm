package handlers

import (
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type InboxHandler struct {
	ctx            *context.ERPContext
	messageService *message.MessageService
}

func NewInboxHandler(ctx *context.ERPContext) *InboxHandler {
	messageService, ok := ctx.MessageService.(*message.MessageService)
	if !ok {
		panic("MessageService is not instance of message.MessageService")
	}

	return &InboxHandler{
		ctx:            ctx,
		messageService: messageService,
	}
}

func (h *InboxHandler) SendMessageHandler(c *gin.Context) {
	var input models.InboxMessageModel
	err := h.messageService.InboxService.SendMessage(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message sent"})
}

func (h *InboxHandler) GetInboxesHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	data, err := h.messageService.InboxService.GetInboxes(nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Inbox retrieved successfully"})
}

func (h *InboxHandler) GetMessagesHandler(c *gin.Context) {
	// memberID := c.MustGet("memberID").(string)
	var inboxID *string
	if c.Query("inboxID") != "" {
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
