package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func NewTelegramRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewTelegramHandler(erpContext)
	group := r.Group("/telegram")
	group.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		group.POST("/webhook/connection/:connectionID", handler.SetUpWebHookHandler)
		group.PUT("/sessions/:session_id/clear", handler.TelegramClearSessionHandler)
		group.DELETE("/sessions/:session_id", handler.TelegramDeleteSessionHandler)
		group.GET("/sessions", handler.GetSessionsHandler)
		group.GET("/sessions/:session_id", handler.GetSessionDetailHandler)
		group.GET("/messages", handler.GetSessionMessagesHandler)
		group.POST("/:session_id/message", handler.SendMessage)
	}

	r.POST("/telegram/webhook", handler.WebhookHandler)

} // end func NewTelegramRoutes
