package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetBroadcastRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	broadcastHandler := handlers.NewBroadcastHandler(erpContext)
	broadcastGroup := r.Group("/broadcast")
	broadcastGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		broadcastGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:read"}), broadcastHandler.GetBroadcastsHandler)
		broadcastGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:read"}), broadcastHandler.GetBroadcastHandler)
		broadcastGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:create"}), broadcastHandler.CreateBroadcastHandler)
		broadcastGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:update"}), broadcastHandler.UpdateBroadcastHandler)
		broadcastGroup.PUT("/:id/send", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:send"}), broadcastHandler.SendBroadcastHandler)
		broadcastGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"whatsapp:broadcast:delete"}), broadcastHandler.DeleteBroadcastHandler)
	}

}
