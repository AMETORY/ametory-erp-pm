package handlers

import (
	rapid_api_models "ametory-pm/models/rapid_api"
	"ametory-pm/objects"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/company"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommonHandler struct {
	ctx             *context.ERPContext
	companyService  *company.CompanyService
	appService      *app.AppService
	pmService       *project_management.ProjectManagementService
	rbacService     *auth.RBACService
	authService     *auth.AuthService
	fileService     *file.FileService
	rapidApiService *services.RapidApiService
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
	fileService, ok := ctx.FileService.(*file.FileService)
	if !ok {
		panic("FileService is not instance of file.FileService")
	}
	rapidApiService, ok := ctx.ThirdPartyServices["RapidAPI"].(*services.RapidApiService)
	if !ok {
		panic("RapidApiService is not instance of services.RapidApiService")
	}
	return &CommonHandler{
		ctx:             ctx,
		companyService:  companyService,
		appService:      appService,
		pmService:       pmService,
		rbacService:     rbacService,
		authService:     authService,
		fileService:     fileService,
		rapidApiService: rapidApiService,
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

func (h *CommonHandler) InvitedHandler(c *gin.Context) {
	members, err := h.pmService.MemberService.GetInvitedMembers(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": members})
}

func (h *CommonHandler) DeleteInvitedHandler(c *gin.Context) {
	invitationID := c.Param("id")
	err := h.pmService.MemberService.DeleteInvitation(invitationID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Invitation deleted successfully"})
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
	data.Email = input.Email

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

func (h *CommonHandler) UploadFileHandler(c *gin.Context) {
	h.ctx.Request = c.Request

	fileObject := models.FileModel{}
	refID, _ := c.GetPostForm("ref_id")
	refType, _ := c.GetPostForm("ref_type")
	skipSave := false
	skipSaveStr, _ := c.GetPostForm("skip_save")
	if skipSaveStr == "true" || skipSaveStr == "1" {
		skipSave = true
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fileByte, err := utils.FileHeaderToBytes(file)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filename := file.Filename

	fileObject.FileName = utils.FilenameTrimSpace(filename)
	fileObject.RefID = refID
	fileObject.RefType = refType
	fileObject.SkipSave = skipSave

	if err := h.fileService.UploadFile(fileByte, "local", "files", &fileObject); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "File uploaded successfully", "data": fileObject})
}

func (h *CommonHandler) CompanySettingHandler(c *gin.Context) {
	h.ctx.Request = c.Request
	data, err := h.companyService.GetCompanyByID(c.GetHeader("ID-Company"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Get company setting successfully", "data": data})
}
func (h *CommonHandler) UpdateCompanySettingHandler(c *gin.Context) {
	var input models.CompanyModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	h.ctx.Request = c.Request
	err = h.companyService.UpdateCompany(c.GetHeader("ID-Company"), &input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": " company setting update successfully"})
}

func (h *CommonHandler) GetRapidAPIPluginsHandler(c *gin.Context) {
	plugins, err := h.rapidApiService.GetPlugins()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "get plugins successfully", "data": plugins})
}

func (h *CommonHandler) AddRapidAdpiPluginHandler(c *gin.Context) {
	input := struct {
		ID   string `json:"id" binding:"required"`
		Key  string `json:"key" binding:"required"`
		Host string `json:"host" binding:"required"`
	}{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var companyPlugins rapid_api_models.CompanyRapidApiPlugin
	err = h.ctx.DB.Where("company_id = ? and rapid_api_plugin_id = ?", c.GetHeader("ID-Company"), input.ID).First(&companyPlugins).Error
	if err == nil {
		c.JSON(400, gin.H{"error": "plugin has added"})
		return
	}

	companyPlugin := rapid_api_models.CompanyRapidApiPlugin{
		CompanyID:        c.GetHeader("ID-Company"),
		RapidApiPluginID: input.ID,
		RapidApiKey:      input.Key,
		RapidApiHost:     input.Host,
	}

	if err := h.ctx.DB.Create(&companyPlugin).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Plugin added successfully"})

}

func (h *CommonHandler) GetCompanyPluginsHandler(c *gin.Context) {
	var companyPlugins []rapid_api_models.CompanyRapidApiPlugin
	err := h.ctx.DB.Preload("RapidApiPlugin.RapidApiEndpoints").Where("company_id = ?", c.GetHeader("ID-Company")).Find(&companyPlugins).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}
	for i, v := range companyPlugins {
		if len(v.RapidApiKey) > 30 {
			v.RapidApiKey = v.RapidApiKey[:len(v.RapidApiKey)-30] + "******************************"
		}
		companyPlugins[i] = v
	}
	c.JSON(200, gin.H{"message": "get company plugins successfully", "data": companyPlugins})
}

func (h *CommonHandler) DeleteCompanyPluginHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.ctx.DB.Where("company_id = ? and rapid_api_plugin_id = ?", c.GetHeader("ID-Company"), id).Delete(&rapid_api_models.CompanyRapidApiPlugin{}).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "delete plugin successfully"})
}
