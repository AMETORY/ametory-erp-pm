package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupTemplateRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewTemplateHandler(erpContext)

	templateGroup := r.Group("/template")
	templateGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		templateGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:read"}), handler.GetTemplatesHandler)
		templateGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:read"}), handler.GetTemplateDetailHandler)
		templateGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:create"}), handler.CreateTemplateHandler)
		templateGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.UpdateTemplateHandler)
		templateGroup.PUT("/:id/add-message", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.AddMessageHandler)
		templateGroup.PUT("/:id/add-image/:msgId", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.AddImageTemplateHandler)
		templateGroup.DELETE("/:id/delete-message/:msgId", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.DeleteMessageTemplate)
		templateGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:delete"}), handler.DeleteTemplateHandler)
		templateGroup.POST("/:id/interactive/:messageId", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.CreateInteractiveTemplateHandler)
		templateGroup.PUT("/:id/interactive/:interactiveId/update", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.UpdateInteractiveTemplateHandler)
		templateGroup.GET("/:id/interactive/:messageId", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:whatsapp_template:update"}), handler.GetInteractiveTemplateHandler)

	}
}
