package routes

import (
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func NewCommonRoutes(r *gin.Engine, erpContext *context.ERPContext) {
	r.Static("/static", "../frontend/build/static")
	r.StaticFile("/", "../frontend/build/index.html")

	// Handle API routes
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from Golang!",
		})
	})

	// Rewrite semua path lainnya ke index.html (untuk SPA)
	r.NoRoute(func(c *gin.Context) {
		c.File("../frontend/build/index.html")
	})

}
