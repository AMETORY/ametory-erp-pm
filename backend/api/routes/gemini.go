package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupGeminiRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	geminiHandler := handlers.NewGeminiHandler(erpContext)
	geminiGroup := r.Group("/gemini")
	geminiGroup.Use(middlewares.AuthMiddleware(erpContext, true), middlewares.RbacSuperAdminMiddleware(erpContext))
	{
		geminiGroup.POST("/generate", geminiHandler.GenerateContentHandler)
		geminiGroup.GET("/agent", geminiHandler.GetAgentHandler)
		geminiGroup.GET("/agent/:id", geminiHandler.GetAgentDetailHandler)
		geminiGroup.POST("/agent", geminiHandler.CreateAgentHandler)
		geminiGroup.PUT("/agent/:id", geminiHandler.UpdateAgentHandler)
		geminiGroup.DELETE("/agent/:id", geminiHandler.DeleteAgentHandler)
	}
}
