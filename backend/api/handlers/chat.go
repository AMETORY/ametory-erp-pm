package handlers

import (
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	ctx            *context.ERPContext
	messageService *message.MessageService
}

func NewChatHandler(ctx *context.ERPContext) *ChatHandler {

	messageService, ok := ctx.MessageService.(*message.MessageService)
	if !ok {
		panic("MessageService is not instance of message.MessageService")
	}
	return &ChatHandler{ctx: ctx, messageService: messageService}
}

func (h *ChatHandler) GetChannelsHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	channels, err := h.messageService.ChatService.GetChannelByParticipantMemberID(memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": channels})
}

func (h *ChatHandler) GetChannelMessageHandler(c *gin.Context) {
	channelID := c.Param("id")
	messages, err := h.messageService.ChatService.GetChatMessageByChannelID(channelID, c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": messages})
}

func (h *ChatHandler) CreateChannelHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	var input models.ChatChannelModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.messageService.ChatService.CreateChannel(&input, nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Channel created", "channel_id": input.ID})
}

func (h *ChatHandler) CreateMessageHandler(c *gin.Context) {
	var input models.ChatMessageModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.messageService.ChatService.CreateMessage(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message created", "message_id": input.ID})
}

func (h *ChatHandler) GetChatMessageDetailHandler(c *gin.Context) {
	messageID := c.Param("messageId")
	data, err := h.messageService.ChatService.GetDetailMessage(messageID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": data, "message": "Message retrieved successfully"})
}

func (h *ChatHandler) UpdateChannelHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	var input models.ChatMessageModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	messageID := c.Param("messageId")
	err := h.messageService.ChatService.UpdateMessage(messageID, &input, nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Channel updated"})
}

func (h *ChatHandler) DeleteChannelHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	channelID := c.Param("id")
	err := h.messageService.ChatService.DeleteChannel(channelID, nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Channel deleted"})
}

func (h *ChatHandler) DeleteMessageHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	messageID := c.Param("messageId")
	err := h.messageService.ChatService.DeleteMessage(messageID, nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Message deleted"})
}
