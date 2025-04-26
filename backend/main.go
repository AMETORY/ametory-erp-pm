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
	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/tag"
	"github.com/AMETORY/ametory-erp-modules/thirdparty"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
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

	fileService := file.NewFileService(erpContext, cfg.Server.BaseURL)
	erpContext.FileService = fileService

	companyService := company.NewCompanyService(erpContext)
	erpContext.CompanyService = companyService

	contactService := contact.NewContactService(erpContext, companyService)
	erpContext.ContactService = contactService

	pmService := project_management.NewProjectManagementService(erpContext)
	erpContext.ProjectManagementService = pmService

	rbacSrv := auth.NewRBACService(erpContext)
	erpContext.RBACService = rbacSrv

	messageSrv := message.NewMessageService(erpContext)
	erpContext.MessageService = messageSrv

	csrService := customer_relationship.NewCustomerRelationshipService(erpContext)
	erpContext.CustomerRelationshipService = csrService

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

	rapidApiService := services.NewRapidAdpiService(erpContext)

	erpContext.AddThirdPartyService("RapidAPI", rapidApiService)

	tagSrv := tag.NewTagService(erpContext)
	erpContext.TagService = tagSrv

	// fmt.Println(erpContext.ThirdPartyServices)

	// GEMINI
	geminiSrv := google.NewGeminiService(erpContext, cfg.Google.GeminiAPIKey)
	geminiSrv.SetupModel(
		1,
		40,
		0.95,
		8192,
		"application/json",
		"gemini-2.0-flash-exp",
	)
	erpContext.AddThirdPartyService("GEMINI", geminiSrv)

	// WA
	waSrv := whatsmeow_client.NewWhatsmeowService(cfg.Whatsapp.BaseURL, cfg.Whatsapp.MockNumber, cfg.Whatsapp.IsMock, "")
	erpContext.AddThirdPartyService("WA", waSrv)

	broadcastSrv := app.NewBroadcastService(erpContext)
	erpContext.AddThirdPartyService("BROADCAST", broadcastSrv)

	routes.NewCommonRoutes(r, erpContext)
	v1 := r.Group("/api/v1")
	routes.SetupWSRoutes(v1, erpContext)
	routes.SetupAuthRoutes(v1, erpContext)
	routes.SetupProjectRoutes(v1, erpContext)
	routes.SetInboxRoutes(v1, erpContext)
	routes.SetChatRoutes(v1, erpContext)
	routes.SetFormRoutes(v1, erpContext)
	routes.SetContactRoutes(v1, erpContext)
	routes.SetupConnectionRoutes(v1, erpContext)
	routes.SetupGeminiRoutes(v1, erpContext)
	routes.NewWhatsappRoutes(v1, erpContext)
	routes.SetBroadcastRoutes(v1, erpContext)
	routes.SetupTagRoutes(v1, erpContext)

	// RUN WORKER
	go func() {
		worker.SendMail(erpContext)
	}()

	r.Run(":" + config.App.Server.Port)
}
