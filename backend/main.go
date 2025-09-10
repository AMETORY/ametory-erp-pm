package main

import (
	"ametory-pm/api/routes"
	"ametory-pm/config"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"ametory-pm/worker"
	ctx "context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/company"
	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/inventory"
	"github.com/AMETORY/ametory-erp-modules/message"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/cache"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/tag"
	"github.com/AMETORY/ametory-erp-modules/thirdparty"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/ai_generator"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/meta"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/gin-contrib/cors"
	"github.com/morkid/paginate"
	"github.com/robfig/cron"

	tiktok "tiktokshop/open/sdk_golang/service"

	mdl "ametory-pm/models"

	"github.com/gin-gonic/gin"
)

var BuildMachineID string

func main() {
	fmt.Println("START AMETORY ERP")

	currentID := getCurrentMachineID()
	if BuildMachineID != "" && currentID != "" {

		fmt.Println("Current Machine ID", currentID)
		fmt.Println("BuildMachineID", BuildMachineID)
		if currentID != BuildMachineID {
			panic("This binary is not allowed to run on this machine.")
		}
	}

	golangCtx := ctx.Background()
	t := time.Now()
	filename := t.Format("2006-01-02")
	logDir := "log"
	logPath := filepath.Join(logDir, filename+".log")

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("error creating directory: %v", err)
	}
	f, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
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
			config.App.Server.BaseURL,
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
	erpContext := context.NewERPContext(db, nil, &golangCtx, skipMigration)
	authService := auth.NewAuthService(erpContext)
	erpContext.AuthService = authService

	fileService := file.NewFileService(erpContext, cfg.Server.BaseURL)
	erpContext.FileService = fileService

	companyService := company.NewCompanyService(erpContext)
	erpContext.CompanyService = companyService

	contactService := contact.NewContactService(erpContext, companyService)
	erpContext.ContactService = contactService

	rbacSrv := auth.NewRBACService(erpContext)
	erpContext.RBACService = rbacSrv

	messageSrv := message.NewMessageService(erpContext)
	erpContext.MessageService = messageSrv

	csrService := customer_relationship.NewCustomerRelationshipService(erpContext)
	erpContext.CustomerRelationshipService = csrService

	inventoryService := inventory.NewInventoryService(erpContext)
	erpContext.InventoryService = inventoryService

	appService := app.NewAppService(erpContext, cfg, redisClient, websocket)
	erpContext.AppService = appService

	if os.Getenv("GEN_PERMISSIONS") != "" {
		for _, v := range appService.GenerateDefaultPermissions() {
			erpContext.DB.Create(&v)
		}
		var companies []models.CompanyModel
		erpContext.DB.Model(&models.CompanyModel{}).Find(&companies)
		for _, v := range companies {
			appService.GenerateDefaultRoles(v.ID)
		}

	}

	// fmt.Println("TIKTOK", cfg.Tiktok.AppKey, cfg.Tiktok.AppSecret, cfg.Tiktok.ServiceID)

	tiktokService := tiktok.NewTiktokService(erpContext, appService, csrService, cfg.Tiktok.AppKey, cfg.Tiktok.AppSecret, cfg.Tiktok.ServiceID)
	erpContext.AddThirdPartyService("Tiktok", tiktokService)

	shopeeService := services.NewShopeeService(erpContext, cfg.Shopee.APISecret, cfg.Shopee.PartnerID, cfg.Shopee.Host, cfg.Shopee.RedireclURL)
	erpContext.AddThirdPartyService("Shopee", shopeeService)

	lazadaService := services.NewLazadaService(erpContext, cfg.Lazada.APIKey, cfg.Lazada.APISecret, cfg.Lazada.Region, cfg.Lazada.CallbackURL)
	erpContext.AddThirdPartyService("Lazada", lazadaService)

	emailSender := thirdparty.NewSMTPSender(cfg.Email.Server, cfg.Email.Port, cfg.Email.Username, cfg.Email.Password, mail.Address{Name: cfg.Email.From, Address: cfg.Email.From})
	emailSender.SetTemplate("../templates/email/layout.html", "../templates/email/body.html")

	erpContext.EmailSender = emailSender

	rapidApiService := services.NewRapidAdpiService(erpContext)

	erpContext.AddThirdPartyService("RapidAPI", rapidApiService)
	erpContext.AddThirdPartyService("Cache", cache.NewCacheManager[paginate.Page]())

	tagSrv := tag.NewTagService(erpContext)
	erpContext.TagService = tagSrv

	aiGeneratorService := ai_generator.NewAiGeneratorService(erpContext.Ctx, erpContext.DB, skipMigration)

	erpContext.AddThirdPartyService("AiGenerator", aiGeneratorService)

	metaService := meta.NewMetaService(db, erpContext, cfg.Server.BaseURL, cfg.Facebook.BaseURL, "local")
	erpContext.AddThirdPartyService("Meta", metaService)

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

	pmService := project_management.NewProjectManagementService(erpContext)
	erpContext.ProjectManagementService = pmService

	broadcastSrv := app.NewBroadcastService(erpContext)
	erpContext.AddThirdPartyService("BROADCAST", broadcastSrv)

	flagIdleColum := flag.Bool("check-idle-column", false, "check idle column")
	flagTestEmail := flag.String("test-email", "", "test email")
	flag.Parse()
	if *flagIdleColum {
		fmt.Println("CHECK IDLE COLUMN")
		worker.CheckIdleColumn(erpContext)
		os.Exit(0)
	}
	if *flagTestEmail != "" {
		worker.TestEmail(*erpContext, *flagTestEmail)
		os.Exit(0)
	}

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
	routes.SetupAiGeminiRoutes(v1, erpContext)
	routes.NewWhatsappRoutes(v1, erpContext)
	routes.SetBroadcastRoutes(v1, erpContext)
	routes.SetupTagRoutes(v1, erpContext)
	routes.SetupProductRoutes(v1, erpContext)
	routes.SetupProductCategoryRoutes(v1, erpContext)
	routes.SetupTemplateRoutes(v1, erpContext)
	routes.SetupAnalyticRoutes(v1, erpContext)
	routes.NewTelegramRoutes(v1, erpContext)
	routes.SetupFacebookRoutes(v1, erpContext)
	routes.SetupLazadaRoutes(v1, erpContext)
	routes.SetupTiktokRoutes(v1, erpContext)
	routes.SetupWhatsappApiRoutes(v1, erpContext)

	// RUN WORKER
	go func() {
		worker.SendMail(erpContext)
	}()
	go func() {
		worker.ScheduledBroadcastWorker(erpContext)
	}()
	go func() {
		worker.BroadcastWorker(erpContext)
	}()
	go func() {
		worker.ScheduledMessageWorker(erpContext)
	}()
	go func() {
		worker.ImportContact(erpContext)
	}()
	go func() {
		worker.FormDownload(erpContext)
	}()

	go func() {
		c := cron.New()
		// c.AddFunc("0 * * * * *", func() { worker.CheckIdleColumn(erpContext) })
		c.AddFunc("@hourly", func() { worker.CheckIdleColumn(erpContext) })
		c.Start()
	}()

	go func() {
		// c := cron.New()
		// c.AddFunc("@every 5m", func() { worker.GetStoppedBroadcasts(erpContext, false) })
		worker.GetStoppedBroadcasts(erpContext)
		// c.Start()
	}()

	// Jalankan server di goroutine agar tidak memblokir
	srv := &http.Server{
		Addr:    ":" + config.App.Server.Port,
		Handler: r,
	}

	go func() {
		// Service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Tunggu sinyal interrupt untuk graceful shutdown.
	quit := make(chan os.Signal, 1)
	// Menerima sinyal SIGINT dan SIGTERM
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Di sini Anda bisa menambahkan perintah atau logika yang ingin dijalankan sebelum server mati.
	// Contoh: menutup koneksi database, menyelesaikan worker, dll.
	log.Println("Running cleanup tasks before shutdown...")
	// erpContext.DB.Close() // Contoh: menutup koneksi DB jika diperlukan

	// Set all PROCESSING broadcasts to STOPPED on startup
	log.Println("Updating all PROCESSING broadcasts to STOPPED status...")
	if err := erpContext.DB.Model(&mdl.BroadcastModel{}).Where("status IN (?)", []string{"PROCESSING", "RESTARTING"}).Debug().Updates(map[string]interface{}{"status": "STOPPED"}).Error; err != nil {
		log.Printf("Error updating broadcast statuses: %v\n", err)
	} else {
		log.Println("Successfully updated broadcast statuses.")
	}

	// Batas waktu 5 detik untuk server menyelesaikan request yang sedang berjalan.
	shutdownCtx, cancel := ctx.WithTimeout(ctx.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func getCurrentMachineID() string {
	switch runtime.GOOS {
	case "linux":
		// Linux pakai UUID dari DMI
		out, err := exec.Command("cat", "/sys/class/dmi/id/product_uuid").Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(out))

	case "darwin":
		// macOS pakai IOPlatformUUID
		out, err := exec.Command("ioreg", "-rd1", "-c", "IOPlatformExpertDevice").Output()
		if err != nil {
			return ""
		}
		lines := strings.Split(string(out), "\n")
		for _, line := range lines {
			if strings.Contains(line, "IOPlatformUUID") {
				parts := strings.Split(line, "\"")
				if len(parts) > 3 {
					return parts[3]
				}
			}
		}
	}

	return ""
}
