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

	}

	r.GET("/whatsapp-api/webhook", handler.WhatsappApiWebhookHandler)
	r.POST("/whatsapp-api/webhook", handler.WhatsappApiWebhookHandler)

}
