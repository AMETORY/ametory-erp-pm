package handlers

import (
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/company"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommonHandler struct {
	ctx            *context.ERPContext
	companyService *company.CompanyService
	appService     *app.AppService
	pmService      *project_management.ProjectManagementService
	rbacService    *auth.RBACService
	authService    *auth.AuthService
}

func NewCommonHandler(ctx *context.ERPContext) *CommonHandler {
	companyService, ok := ctx.CompanyService.(*company.CompanyService)
	if !ok {
		panic("CompanyService is not instance of company.CompanyService")
	}
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}
	rbacService, ok := ctx.RBACService.(*auth.RBACService)
	if !ok {
		panic("RBACService is not instance of auth.RBACService")
	}
	authService, ok := ctx.AuthService.(*auth.AuthService)
	if !ok {
		panic("AuthService is not instance of auth.AuthService")
	}
	return &CommonHandler{
		ctx:            ctx,
		companyService: companyService,
		appService:     appService,
		pmService:      pmService,
		rbacService:    rbacService,
		authService:    authService,
	}
}

func (h *CommonHandler) GetMembersHandler(c *gin.Context) {
	members, err := h.pmService.MemberService.GetMembers(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": members})
}

func (h *CommonHandler) GetRolesHandler(c *gin.Context) {
	roles, err := h.rbacService.GetAllRoles(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	items := roles.Items.(*[]models.RoleModel)
	newItems := make([]models.RoleModel, 0)
	for _, v := range *items {
		if !v.IsSuperAdmin {
			v.Permissions = nil
			newItems = append(newItems, v)
		}
	}
	roles.Items = &newItems
	c.JSON(200, gin.H{"data": roles})
}

func (h *CommonHandler) InviteMemberHandler(c *gin.Context) {
	var input struct {
		FullName  string  `json:"full_name"`
		RoleID    *string `json:"role_id"`
		Email     string  `json:"email"`
		ProjectID *string `json:"project_id"`
	}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var data models.MemberInvitationModel
	data.FullName = input.FullName
	data.RoleID = input.RoleID
	data.ProjectID = input.ProjectID

	var user models.UserModel
	var link = ""
	var password = ""

	err = h.ctx.DB.Where("email = ?", input.Email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Create new user if not exists
		username := utils.CreateUsernameFromFullName(input.FullName)
		// fmt.Println("username", username)
		password = utils.RandString(8, false)
		u, err := h.authService.Register(input.FullName, username, input.Email, password, "")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user = *u

	} else {
		user.Password = ""
	}

	var company models.CompanyModel
	h.ctx.DB.Where("id = ?", c.GetHeader("ID-Company")).First(&company)

	err = h.ctx.DB.Model(&user).Association("Companies").Append(&company)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.UserID = user.ID
	if user.VerificationToken != "" {
		data.Token = user.VerificationToken
	}

	data.InviterID = c.MustGet("userID").(string)
	data.CompanyID = &company.ID
	token, err := h.pmService.MemberService.InviteMember(&data)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	link = fmt.Sprintf("%s/invitation/verify/%s", h.appService.Config.Server.FrontendURL, token)
	notif := fmt.Sprintf("Anda telah diundang untuk bergabung di perusahaan %s ", company.Name)
	if input.ProjectID != nil {
		var project models.ProjectModel
		h.ctx.DB.Where("id = ?", *input.ProjectID).First(&project)
		notif += fmt.Sprintf("dalam proyek %s", project.Name)
	}
	var emailData objects.EmailData = objects.EmailData{
		FullName: user.FullName,
		Email:    user.Email,
		Subject:  "Selamat datang di Ametory Project Manager",
		Notif:    notif,
		Link:     link,
		Password: password,
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

	c.JSON(200, gin.H{"message": "Member invited successfully", "token": token})
}

func (h *CommonHandler) AcceptMemberInvitationHandler(c *gin.Context) {
	token := c.Param("token")

	var invitation models.MemberInvitationModel
	h.ctx.DB.Where("token = ?", token).First(&invitation)

	err := h.pmService.MemberService.AcceptMemberInvitation(token, invitation.UserID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user models.UserModel
	h.ctx.DB.Where("id = ?", invitation.UserID).First(&user)
	now := time.Now()
	if user.VerifiedAt == nil {
		user.VerifiedAt = &now
		user.VerificationToken = ""
		user.VerificationTokenExpiredAt = nil
		h.ctx.DB.Save(&user)
	}
	c.JSON(200, gin.H{"message": "Member invitation accepted successfully"})
}
