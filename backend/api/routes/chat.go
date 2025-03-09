package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetChatRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	chatHandler := handlers.NewChatHandler(erpContext)
	chatGroup := r.Group("/chat")
	chatGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		chatGroup.GET("/channels", chatHandler.GetChannelsHandler)
		chatGroup.GET("/channel/:id", chatHandler.GetChannelDetailHandler)
		chatGroup.GET("/channel/:id/messages", chatHandler.GetChannelMessageHandler)
		chatGroup.POST("/channel", chatHandler.CreateChannelHandler)
		chatGroup.POST("/channel/:id/message", chatHandler.CreateMessageHandler)
		chatGroup.GET("/channel/:id/message/:messageId", chatHandler.GetChatMessageDetailHandler)
		chatGroup.PUT("/channel/:id", chatHandler.UpdateChannelHandler)
		chatGroup.DELETE("/channel/:id", chatHandler.DeleteChannelHandler)
		chatGroup.DELETE("/channel/:id/message/:messageId", chatHandler.DeleteMessageHandler)
	}
}
