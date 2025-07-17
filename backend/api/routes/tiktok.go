package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupTiktokRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewTiktokHandler(erpContext)
	group := r.Group("/tiktok")
	group.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		group.GET("/sessions", handler.GetSessionsHandler)
		group.GET("/sessions/:sessionId", handler.GetSessionDetailHandler)
		group.POST("/sessions/:sessionId/message", handler.SendMessageHandler)
		group.POST("/sessions/:sessionId/file", handler.SendFileHandler)
		group.GET("/sessions/:sessionId/messages", handler.GetSessionMessagesHandler)
	}
	r.GET("/tiktok/webhook", handler.WebhookHandler)
}
