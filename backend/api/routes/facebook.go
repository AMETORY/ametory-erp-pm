package routes

import (
	"ametory-pm/api/handlers"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupFacebookRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewFacebookHandler(erpContext)
	facebook := r.Group("/facebook")
	{
		facebook.GET("/callback/:connectionID", handler.FacebookCallbackHandler)
	}

}
