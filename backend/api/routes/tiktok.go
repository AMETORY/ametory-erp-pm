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
	}
	r.GET("/tiktok/webhook", handler.WebhookHandler)
}
