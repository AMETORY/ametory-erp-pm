package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupWhatsappApiRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewWhatsappApiHandler(erpContext)
	group := r.Group("/whatsapp-api")
	group.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		group.GET("/message-templates/:waba_id", handler.GetAllMessageTemplates)
		group.GET("/message-templates/:waba_id/:template_id", handler.GetMessageTemplateByName)
	}

	r.GET("/whatsapp-api/webhook", handler.WhatsappApiWebhookHandler)
	r.POST("/whatsapp-api/webhook", handler.WhatsappApiWebhookHandler)

}
