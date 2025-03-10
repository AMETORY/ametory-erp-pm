package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetFormRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	formHandler := handlers.NewFormHandler(erpContext)
	formTemplateGroup := r.Group("/form-template")
	formTemplateGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		formTemplateGroup.GET("/list", formHandler.GetFormTemplatesHandler)
		formTemplateGroup.GET("/:id", formHandler.GetFormTemplateHandler)
		formTemplateGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form_template:create"}), formHandler.CreateFormTemplateHandler)
		formTemplateGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form_template:update"}), formHandler.UpdateFormTemplateHandler)
		formTemplateGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form_template:delete"}), formHandler.DeleteFormTemplateHandler)
	}

	formGroup := r.Group("/form")
	formGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		formGroup.GET("/list", formHandler.GetFormsHandler)
		formGroup.GET("/:id", formHandler.GetFormHandler)
		formGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form:create"}), formHandler.CreateFormHandler)
		formGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form:update"}), formHandler.UpdateFormHandler)
		formGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"customer_relationship:form:delete"}), formHandler.DeleteFormHandler)
	}
}
