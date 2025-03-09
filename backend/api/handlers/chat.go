package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type ChatHandler struct {
	ctx            *context.ERPContext
	messageService *message.MessageService
	appService     *app.AppService
}

func NewChatHandler(ctx *context.ERPContext) *ChatHandler {

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}

	messageService, ok := ctx.MessageService.(*message.MessageService)
	if !ok {
		panic("MessageService is not instance of message.MessageService")
	}
	return &ChatHandler{ctx: ctx, messageService: messageService, appService: appService}
}

func (h *ChatHandler) GetChannelsHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	channels, err := h.messageService.ChatService.GetChannelByParticipantMemberID(memberID, c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": channels})
}
func (h *ChatHandler) GetChannelDetailHandler(c *gin.Context) {
	channelID := c.Param("id")
	channel, err := h.messageService.ChatService.GetChannelDetail(channelID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": channel})
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
	input.CreatedByMemberID = &memberID
	err := h.messageService.ChatService.CreateChannel(&input, nil, &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if input.Avatar != nil {
		input.Avatar.RefID = input.ID
		input.Avatar.RefType = "chat"
		h.ctx.DB.Save(input.Avatar)
	}
	c.JSON(200, gin.H{"message": "Channel created", "channel_id": input.ID})
}

func (h *ChatHandler) CreateMessageHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	channelID := c.Param("id")
	var input models.ChatMessageModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	input.ChatChannelID = &channelID
	input.SenderMemberID = &memberID
	input.ChatData = "{}"
	input.Date = &now
	err := h.messageService.ChatService.CreateMessage(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for i, v := range input.Files {
		v.RefID = input.ID
		v.RefType = "chat"
		h.ctx.DB.Save(&v)
		input.Files[i] = v
	}

	var channel models.ChatChannelModel
	h.ctx.DB.Preload("Avatar").Where("id = ?", channelID).Find(&channel)

	member := c.MustGet("member").(models.MemberModel)
	user := c.MustGet("user").(models.UserModel)
	member.User = user
	input.SenderMember = &member
	channelURL := ""
	if channel.Avatar != nil {
		channelURL = channel.Avatar.URL
	}
	msg := gin.H{
		"message":        "Message created",
		"channel_id":     channelID,
		"channel_name":   channel.Name,
		"channel_avatar": channelURL,
		"command":        "RECEIVE_MESSAGE",
		"data":           input,
		"sender_id":      c.MustGet("userID").(string),
		"sender_name":    member.User.FullName,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
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

func (h *ChatHandler) AddChannelParticipant(c *gin.Context) {
	channelID := c.Param("id")
	var input []string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, v := range input {
		err := h.messageService.ChatService.AddParticipant(channelID, nil, &v)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		msg := gin.H{
			"message":      "Participant added to channel successfully",
			"channel_id":   channelID,
			"command":      "CHANNEL_RELOAD",
			"data":         input,
			"sender_id":    c.MustGet("userID").(string),
			"recipient_id": v,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}

	c.JSON(200, gin.H{"message": "Participant added to channel successfully"})
}

func (h *ChatHandler) DeleteChannelParticipant(c *gin.Context) {
	channelID := c.Param("id")
	var input []string
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	for _, v := range input {
		err := h.messageService.ChatService.DeleteParticipant(channelID, nil, &v)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		msg := gin.H{
			"message":      "Participant added to channel successfully",
			"channel_id":   channelID,
			"command":      "CHANNEL_RELOAD",
			"data":         input,
			"sender_id":    c.MustGet("userID").(string),
			"recipient_id": v,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}

	c.JSON(200, gin.H{"message": "Participant removed from channel successfully"})
}
