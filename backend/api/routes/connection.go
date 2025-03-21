package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupConnectionRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	connectionHandler := handlers.NewConnectionHandler(erpContext)
	connectionGroup := r.Group("/connection")
	connectionGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		connectionGroup.GET("/list", connectionHandler.GetConnectionsHandler)
		connectionGroup.GET("/:id", connectionHandler.GetConnectionHandler)
		connectionGroup.POST("/create", connectionHandler.CreateConnectionHandler)
		connectionGroup.PUT("/:id", connectionHandler.UpdateConnectionHandler)
		connectionGroup.PUT("/:id/connect", connectionHandler.ConnectDeviceHandler)
		connectionGroup.PUT("/:id/get-qr/:session", connectionHandler.GetQRDeviceHandler)
		connectionGroup.DELETE("/:id", connectionHandler.DeleteConnectionHandler)
	}

}
