package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetContactRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	contactHandler := handlers.NewContactHandler(erpContext)

	contactGroup := r.Group("/contact")
	contactGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		contactGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"contact:customer:read"}), contactHandler.GetContactsHandler)
		contactGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"contact:customer:read"}), contactHandler.GetContactHandler)
		contactGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"contact:customer:create"}), contactHandler.CreateContactHandler)
		contactGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"contact:customer:update"}), contactHandler.UpdateContactHandler)
		contactGroup.PUT("/:id/message", contactHandler.SendMessageContactHandler)
		contactGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"contact:customer:delete"}), contactHandler.GetContactsHandler)
	}
}
