package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func SetupProductCategoryRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	productCategoryHandler := handlers.NewProductCategoryHandler(erpContext)
	productCategoryGroup := r.Group("/product-category")
	productCategoryGroup.Use(middlewares.AuthMiddleware(erpContext, false))
	{
		productCategoryGroup.GET("/list", productCategoryHandler.ListProductCategoriesHandler)
		productCategoryGroup.GET("/:id", productCategoryHandler.GetProductCategoryHandler)
		productCategoryGroup.POST("/create", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:product_category:create"}), productCategoryHandler.CreateProductCategoryHandler)
		productCategoryGroup.PUT("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:product_category:update"}), productCategoryHandler.UpdateProductCategoryHandler)
		productCategoryGroup.DELETE("/:id", middlewares.RbacUserMiddleware(erpContext, []string{"inventory:product_category:delete"}), productCategoryHandler.DeleteProductCategoryHandler)
	}
}
