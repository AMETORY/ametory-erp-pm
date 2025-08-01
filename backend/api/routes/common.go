package routes

import (
	"ametory-pm/api/handlers"
	"ametory-pm/api/middlewares"
	"net/http"
	"strings"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func NewCommonRoutes(r *gin.Engine, erpContext *context.ERPContext) {
	r.Use(BlockGitMiddleware())
	r.Static("/static", "../frontend/build/static")
	r.Static("/assets/files", "../backend/assets/files")
	r.Static("/assets/static", "../backend/assets/static")
	r.StaticFile("/android-chrome-192x192.png", "../frontend/build/android-chrome-192x192.png")
	r.StaticFile("/android-chrome-512x512.png", "../frontend/build/android-chrome-512x512.png")
	r.StaticFile("/apple-touch-icon.png", "../frontend/build/apple-touch-icon.png")
	r.StaticFile("/asset-manifest.json", "../frontend/build/asset-manifest.json")
	r.StaticFile("/favicon-16x16.png", "../frontend/build/favicon-16x16.png")
	r.StaticFile("/favicon-32x32.png", "../frontend/build/favicon-32x32.png")
	r.StaticFile("/ss.png", "../frontend/build/ss.png")
	r.StaticFile("/logo.png", "../frontend/build/logo.png")
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
	r.DELETE("/api/v1/file/:id", middlewares.AuthMiddleware(erpContext, false), commonHander.DeleteFileHandler)
	r.GET("/api/v1/members", middlewares.AuthMiddleware(erpContext, false), commonHander.GetMembersHandler)
	r.GET("/api/v1/roles", middlewares.AuthMiddleware(erpContext, false), commonHander.GetRolesHandler)
	r.POST("/api/v1/company", middlewares.AuthMiddleware(erpContext, false), commonHander.CreateCompanyHandler)
	r.GET("/api/v1/accept-invitation/:token", commonHander.AcceptMemberInvitationHandler)
	r.POST("/api/v1/invite-member", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.InviteMemberHandler)
	r.PUT("/api/v1/members/:id", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:update"}), commonHander.UpdateMemberHandler)
	r.GET("/api/v1/invited", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.InvitedHandler)
	r.DELETE("/api/v1/invited/:id", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacUserMiddleware(erpContext, []string{"project_management:member:invite"}), commonHander.DeleteInvitedHandler)
	r.GET("/api/v1/setting", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.CompanySettingHandler)
	r.PUT("/api/v1/setting", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.UpdateCompanySettingHandler)
	r.GET("/api/v1/rapid-api-plugins", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.GetRapidAPIPluginsHandler)
	r.POST("/api/v1/add-rapid-api-plugin", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.AddRapidAdpiPluginHandler)
	r.GET("/api/v1/company-rapid-api-plugins", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.GetCompanyPluginsHandler)
	r.DELETE("/api/v1/company-rapid-api-plugin/:id", middlewares.AuthMiddleware(erpContext, false), middlewares.RbacSuperAdminMiddleware(erpContext), commonHander.DeleteCompanyPluginHandler)

}

func BlockGitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Cek jika path URL dimulai dengan "/.git"
		if strings.HasPrefix(c.Request.URL.Path, "/.git") {
			// Hentikan request chain dan kirim status 404 Not Found
			// 404 lebih baik dari 403 karena tidak mengkonfirmasi keberadaan resource
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		// Jika tidak, lanjutkan ke handler berikutnya
		c.Next()
	}
}
