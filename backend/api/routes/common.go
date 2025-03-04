package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"
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

	commonHander := handlers.NewCommonHandler(erpContext)
	r.GET("/api/v1/members", middlewares.AuthMiddleware(erpContext, false), commonHander.GetMembersHandler)
	r.GET("/api/v1/roles", middlewares.AuthMiddleware(erpContext, false), commonHander.GetRolesHandler)
	r.GET("/api/v1/accept-invitation/:token", commonHander.AcceptMemberInvitationHandler)
	r.POST("/api/v1/invite-member", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.InviteMemberHandler)

}
