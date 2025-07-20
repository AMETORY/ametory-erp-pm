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
		connectionGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:read"}), connectionHandler.GetConnectionsHandler)
		connectionGroup.GET("/auth-url/shopee", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.GetShopeeAuthURLHandler)
		connectionGroup.GET("/auth-url/lazada", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.GetLazadaAuthURLHandler)
		connectionGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:read"}), connectionHandler.GetConnectionHandler)
		connectionGroup.POST("/create", connectionHandler.CreateConnectionHandler)
		connectionGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:update"}), connectionHandler.UpdateConnectionHandler)
		connectionGroup.PUT("/:id/authorize", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.AuthorizeConnectionHandler)
		connectionGroup.PUT("/:id/sync-contact", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.SyncContactConnectionHandler)
		connectionGroup.PUT("/:id/connect", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.ConnectDeviceHandler)
		connectionGroup.PUT("/:id/get-qr/:session", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:create"}), connectionHandler.GetQRDeviceHandler)
		connectionGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"connection:connection:delete"}), connectionHandler.DeleteConnectionHandler)
	}

}
