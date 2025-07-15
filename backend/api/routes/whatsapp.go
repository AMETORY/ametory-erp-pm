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
		// http://localhost:8081/api/v1/whatsapp/sessions?session_id=f220706b-5f07-423b-8f3a-4793e4937fd5&page=1&size=20&connection_session=6283899133519%3A65%40s.whatsapp.net
		waGroup.GET("/sessions", waHandler.GetSessionsHandler)
		waGroup.GET("/sessions/:session_id", waHandler.GetSessionDetailHandler)
		waGroup.DELETE("/sessions/:session_id", waHandler.DeleteSessionHandler)
		waGroup.PUT("/sessions/:session_id/clear", waHandler.ClearSessionHandler)
		waGroup.PUT("/sessions/:session_id", waHandler.UpdateSessionHandler)
		waGroup.GET("/messages", waHandler.GetSessionMessagesHandler)
		waGroup.PUT("/messages/:messageId/read", waHandler.MarkAsReadHandler)
		waGroup.POST("/:session_id/message", waHandler.SendMessage)
		waGroup.POST("/:session_id/read-all", waHandler.ReadAllMessage)
		waGroup.GET("/devices", waHandler.GetDevice)
		waGroup.PUT("/update-webhook", waHandler.UpdateWebhook)
		waGroup.POST("/create-qr", waHandler.CreateQR)
		waGroup.POST("/export", waHandler.ExportHandler)
		waGroup.GET("/get-qr/:session", waHandler.GetQR)
		waGroup.GET("/conversations/:jid", waHandler.GetConversationsHandler)
		waGroup.GET("/messages/:jid", waHandler.GetMessagesHandler)
		waGroup.DELETE("/delete-device/:jid", waHandler.DeleteDevice)

	}

	r.POST("/whatsapp-webhook", waHandler.WhatsappWebhookHandler)
	r.POST("/whatsapp/register", waHandler.WhatsappRegisterHandler)
	r.GET("/whatsapp/get-number/:code", waHandler.WhatsappGetNumberHandler)
}
