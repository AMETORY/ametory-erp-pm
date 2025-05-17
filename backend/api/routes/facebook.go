package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupFacebookRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewFacebookHandler(erpContext)
	facebook := r.Group("/facebook")
	facebook.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		// facebook.GET("/callback/:connectionID", handler.FacebookCallbackHandler)

		facebook.POST("/instagram/:sessionId/message", handler.SendInstagramMessageHandler)
		facebook.GET("/instagram/sessions", handler.GetInstagramSessionsHandler)
		facebook.GET("/instagram/sessions/:session_id", handler.GetInstagramSessionDetailHandler)
		facebook.GET("/instagram/messages", handler.GetSessionMessagesHandler)
	}

	r.GET("/facebook/callback", handler.FacebookCallbackHandler)
	r.GET("/facebook/instagram/callback", handler.InstagramCallbackHandler)
	r.GET("/facebook/webhook", handler.FacebookWebhookHandler)
	r.POST("/facebook/webhook", handler.FacebookWebhookHandler)
}
