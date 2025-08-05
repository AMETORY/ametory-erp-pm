package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupAiGeminiRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	aiAgentHandler := handlers.NewAiAgentHandler(erpContext)
	aiAgentGroup := r.Group("/ai-agent")
	aiAgentGroup.Use(middlewares.AuthMiddleware(erpContext, true), middlewares.RbacSuperAdminMiddleware(erpContext))
	{
		aiAgentGroup.POST("/generate", aiAgentHandler.GenerateContentHandler)
		aiAgentGroup.GET("/agent", aiAgentHandler.GetAgentHandler)
		aiAgentGroup.GET("/agent/:id", aiAgentHandler.GetAgentDetailHandler)
		aiAgentGroup.GET("/agent/:id/histories", aiAgentHandler.GetAgentHistoriesHandler)
		aiAgentGroup.DELETE("/agent/:id/history/:historyId", aiAgentHandler.DeleteHistoryHandler)
		aiAgentGroup.PUT("/agent/:id/history/:historyId", aiAgentHandler.UpdateHistoryHandler)
		aiAgentGroup.PUT("/agent/:id/history/:historyId/toggle-model", aiAgentHandler.ToggleModelHistoryHandler)
		aiAgentGroup.POST("/agent", aiAgentHandler.CreateAgentHandler)
		aiAgentGroup.PUT("/agent/:id", aiAgentHandler.UpdateAgentHandler)
		aiAgentGroup.DELETE("/agent/:id", aiAgentHandler.DeleteAgentHandler)
	}
}
