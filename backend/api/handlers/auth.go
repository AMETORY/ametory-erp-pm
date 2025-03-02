package handlers

import (
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/company"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	ctx            *context.ERPContext
	authService    *auth.AuthService
	companyService *company.CompanyService
	rbacService    *auth.RBACService
	appService     *app.AppService
}

func NewAuthHandler(ctx *context.ERPContext) *AuthHandler {
	authService, ok := ctx.AuthService.(*auth.AuthService)
	if !ok {
		panic("AuthService is not instance of auth.AuthService")
	}

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	companyService, ok := ctx.CompanyService.(*company.CompanyService)
	if !ok {
		panic("CompanyService is not instance of company.CompanyService")
	}
	rbacService, ok := ctx.RBACService.(*auth.RBACService)
	if !ok {
		panic("RBACService is not instance of auth.RBACService")
	}
	return &AuthHandler{
		ctx:            ctx,
		authService:    authService,
		appService:     appService,
		companyService: companyService,
		rbacService:    rbacService,
	}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	user, err := h.authService.Login(input.Email, input.Password, true)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(user.ID, time.Now().AddDate(0, 0, h.appService.Config.Server.TokenExpiredDay).Unix(), h.appService.Config.Server.SecretKey)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": token})
}

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var input struct {
		FullName    string `json:"full_name" binding:"required"`
		Email       string `json:"email" binding:"required"`
		CompanyName string `json:"company_name" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Password    string `json:"password"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	username := utils.CreateUsernameFromFullName(input.FullName)

	if input.Password == "" {
		input.Password = utils.RandString(8, false)
	}

	user, err := h.authService.Register(input.FullName, username, input.Email, input.Password, input.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var company models.CompanyModel
	company.Name = input.CompanyName
	err = h.companyService.CreateCompany(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.ctx.DB.Model(&user).Association("Companies").Append(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generate default roles
	var superAdmin *models.RoleModel
	roles := h.appService.GenerateDefaultRoles(company.ID)

	for _, v := range roles {
		if v.IsSuperAdmin {
			user.Roles = append(user.Roles, v)
			h.ctx.DB.Save(&user)
			superAdmin = &v
		}
	}

	// CREATE MEMBER
	if superAdmin != nil {
		member := models.MemberModel{
			UserID:    user.ID,
			CompanyID: &company.ID,
			RoleID:    &superAdmin.ID,
		}
		err = h.ctx.DB.Create(&member).Error
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var emailData objects.EmailData = objects.EmailData{
		FullName: user.FullName,
		Email:    user.Email,
		Subject:  "Selamat datang di Ametory Project Manager",
		Notif:    "Silakan verifikasi akun Anda, dengan mengklik link berikut",
		Link:     fmt.Sprintf("%s/verify/%s", h.appService.Config.Server.FrontendURL, user.VerificationToken),
		Password: input.Password,
	}

	b, err := json.Marshal(emailData)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	// fmt.Println("SEND MAIL", string(b))
	err = h.appService.Redis.Publish(*h.ctx.Ctx, "SEND:MAIL", string(b)).Err()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User registered successfully, please check your email to verify your account"})

}

func (h *AuthHandler) ForgotPasswordHandler(c *gin.Context) {
	var input struct {
		EmailOrPhoneNumber string `json:"email_or_phone_number"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"message": "Invalid input", "error": err.Error()})
		return
	}
	authSrv, ok := h.ctx.AuthService.(*auth.AuthService)
	if !ok {
		c.JSON(500, gin.H{"message": "Auth service is not available"})
		return
	}
	user, err := authSrv.GetUserByEmailOrPhone(input.EmailOrPhoneNumber)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	newPassword := utils.RandString(8, false)

	hashedPassword, err := models.HashPassword(newPassword)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}

	user.Password = hashedPassword

	if err := h.ctx.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	var emailData objects.EmailData = objects.EmailData{
		FullName: user.FullName,
		Email:    user.Email,
		Subject:  "Permintaan Penggatian PASSWORD",
		Notif:    "Berikut ini adalah PASSWORD baru Anda",
		Link:     "",
		Password: newPassword,
	}

	b, err := json.Marshal(emailData)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	fmt.Println("SEND MAIL", string(b))
	err = h.appService.Redis.Publish(*h.ctx.Ctx, "SEND:MAIL", string(b)).Err()
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "New password sent"})
}

func (h *AuthHandler) ChangePasswordHandler(c *gin.Context) {
	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.authService.GetUserByID(c.MustGet("userID").(string))
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	err = h.authService.ChangePassword(user.ID, input.OldPassword, input.NewPassword)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Password changed successfully"})
}

func (h *AuthHandler) VerificationEmailHandler(c *gin.Context) {
	token := c.Param("token")
	if h.authService == nil {
		c.JSON(500, gin.H{"message": "Auth service is not available"})
		return
	}
	err := h.authService.VerificationEmail(token)
	if err != nil {
		c.JSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Email verified"})
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	user := c.MustGet("user").(models.UserModel)
	h.ctx.DB.Preload("Companies").Find(&user)
	c.JSON(200, gin.H{"user": user})
}
