package routes

import (
	"ametory-pm/api/handlers"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetInboxRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	inboxHandler := handlers.NewInboxHandler(erpContext)
	inboxGroup := r.Group("/inbox")
	{
		inboxGroup.POST("/send", inboxHandler.SendMessageHandler)
		inboxGroup.GET("/inboxes", inboxHandler.GetInboxesHandler)
		inboxGroup.GET("/messages", inboxHandler.GetMessagesHandler)
		inboxGroup.GET("/message/:id", inboxHandler.GetMessagesDetailHandler)
		inboxGroup.GET("/count", inboxHandler.CountUnreadHandler)
	}

}
