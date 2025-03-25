package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func NewWhatsappRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	waHandler := handlers.NewWhatsappHandler(erpContext)

	waGroup := r.Group("/whatsapp")
	waGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		waGroup.GET("/sessions", waHandler.GetSessionsHandler)
		waGroup.GET("/sessions/:session_id", waHandler.GetSessionDetailHandler)
		waGroup.PUT("/sessions/:session_id", waHandler.UpdateSessionHandler)
		waGroup.GET("/messages", waHandler.GetSessionMessagesHandler)
		waGroup.POST("/:session_id/message", waHandler.SendMessage)
		waGroup.GET("/devices", waHandler.GetDevice)
		waGroup.PUT("/update-webhook", waHandler.UpdateWebhook)
		waGroup.POST("/create-qr", waHandler.CreateQR)
		waGroup.GET("/get-qr/:session", waHandler.GetQR)
		waGroup.GET("/conversations/:jid", waHandler.GetConversationsHandler)
		waGroup.GET("/messages/:jid", waHandler.GetMessagesHandler)
		waGroup.DELETE("/delete-device/:jid", waHandler.DeleteDevice)

	}

	r.POST("/whatsapp-webhook", waHandler.WhatsappWebhookHandler)
	r.POST("/whatsapp/register", waHandler.WhatsappRegisterHandler)
	r.GET("/whatsapp/get-number/:code", waHandler.WhatsappGetNumberHandler)
}
