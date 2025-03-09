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
	r.Static("/assets/files", "../backend/assets/files")
	r.StaticFile("/android-chrome-192x192.png", "../frontend/build/android-chrome-192x192.png")
	r.StaticFile("/android-chrome-512x512.png", "../frontend/build/android-chrome-512x512.png")
	r.StaticFile("/apple-touch-icon.png", "../frontend/build/apple-touch-icon.png")
	r.StaticFile("/asset-manifest.json", "../frontend/build/asset-manifest.json")
	r.StaticFile("/favicon-16x16.png", "../frontend/build/favicon-16x16.png")
	r.StaticFile("/favicon-32x32.png", "../frontend/build/favicon-32x32.png")
	r.StaticFile("/ss.png", "../frontend/build/ss.png")
	r.StaticFile("/favicon.ico", "../frontend/build/favicon.ico")
	r.StaticFile("/site.webmanifest", "../frontend/build/site.webmanifest")
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
	r.POST("/api/v1/file/upload", middlewares.AuthMiddleware(erpContext, false), commonHander.UploadFileHandler)
	r.GET("/api/v1/members", middlewares.AuthMiddleware(erpContext, false), commonHander.GetMembersHandler)
	r.GET("/api/v1/roles", middlewares.AuthMiddleware(erpContext, false), commonHander.GetRolesHandler)
	r.GET("/api/v1/accept-invitation/:token", commonHander.AcceptMemberInvitationHandler)
	r.POST("/api/v1/invite-member", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.InviteMemberHandler)
	r.GET("/api/v1/invited", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.InvitedHandler)
	r.DELETE("/api/v1/invited/:id", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.DeleteInvitedHandler)
}
