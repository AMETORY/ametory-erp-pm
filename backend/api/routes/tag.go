package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupTagRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewTagHandler(erpContext)

	tagGroup := r.Group("/tag")
	tagGroup.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		tagGroup.GET("/list", middlewares.RbacUserMiddleware(erpContext, []string{"tag:read"}), handler.GetTagsHandler)
		tagGroup.GET("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"tag:read"}), handler.GetTagDetailHandler)
		tagGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"tag:create"}), handler.CreateTagHandler)
		tagGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"tag:update"}), handler.UpdateTagHandler)
		tagGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"tag:delete"}), handler.DeleteTagHandler)
	}
}
