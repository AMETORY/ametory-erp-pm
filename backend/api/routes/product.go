package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupProductRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	handler := handlers.NewProductHandler(erpContext)
	group := r.Group("/product")
	group.Use(middlewares.AuthMiddleware(erpContext, true))
	{
		group.GET("/list", handler.ListProductsHandler)
		group.GET("/:id", handler.GetProductHandler)
		group.POST("/create", handler.CreateProductHandler)
		group.PUT("/:id", handler.UpdateProductHandler)

		group.DELETE("/:id", handler.DeleteProductHandler)
		group.DELETE("/:id/image/:imageId", handler.DeleteImageProductHandler)
	}

}
