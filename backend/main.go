package main

import (
	"ametory-pm/api/routes"
	"ametory-pm/config"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"ametory-pm/worker"
	ctx "context"
	"net/mail"
	"os"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/company"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/thirdparty"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := ctx.Background()
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	db, err := services.InitDB(cfg)
	if err != nil {
		panic(err)
	}
	redisClient := services.InitRedis()
	websocket := services.InitWS()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
			"http://localhost:3000/",
		},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "DELETE", "HEAD"},
		AllowHeaders: []string{
			"Origin",
			"Authorization",
			"Content-Length",
			"Content-Type",
			"Access-Control-Allow-Origin",
			"API-KEY",
			"Currency-Code",
			"Cache-Control",
			"X-Requested-With",
			"Content-Disposition",
			"Content-Description",
			"ID-Company",
			"ID-Distributor",
			"timezone",
		},
		ExposeHeaders: []string{"Content-Length", "Content-Disposition", "Content-Description"},
	}))

	skipMigration := true

	if os.Getenv("MIGRATION") != "" {
		skipMigration = false
	}
	erpContext := context.NewERPContext(db, nil, &ctx, skipMigration)
	authService := auth.NewAuthService(erpContext)
	erpContext.AuthService = authService

	companyService := company.NewCompanyService(erpContext)
	erpContext.CompanyService = companyService

	pmService := project_management.NewProjectManagementService(erpContext)
	erpContext.ProjectManagementService = pmService

	rbacSrv := auth.NewRBACService(erpContext)
	erpContext.RBACService = rbacSrv

	appService := app.NewAppService(erpContext, cfg, redisClient, websocket)
	erpContext.AppService = appService

	if os.Getenv("GEN_PERMISSIONS") != "" {
		for _, v := range appService.GenerateDefaultPermissions() {
			erpContext.DB.Create(&v)
		}
	}

	emailSender := thirdparty.NewSMTPSender(cfg.Email.Server, cfg.Email.Port, cfg.Email.Username, cfg.Email.Password, mail.Address{Name: cfg.Email.From, Address: cfg.Email.From})
	emailSender.SetTemplate("../templates/email/layout.html", "../templates/email/body.html")

	erpContext.EmailSender = emailSender

	routes.NewCommonRoutes(r, erpContext)
	v1 := r.Group("/api/v1")
	routes.SetupWSRoutes(v1, erpContext)
	routes.SetupAuthRoutes(v1, erpContext)
	routes.SetupProjectRoutes(v1, erpContext)

	// RUN WORKER
	go func() {
		worker.SendMail(erpContext)
	}()

	r.Run(":" + config.App.Server.Port)
}
