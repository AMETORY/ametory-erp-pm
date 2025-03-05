package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm/clause"
)

type ProjectHandler struct {
	ctx        *context.ERPContext
	pmService  *project_management.ProjectManagementService
	appService *app.AppService
}

func NewProjectHandler(ctx *context.ERPContext) *ProjectHandler {
	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	return &ProjectHandler{
		ctx:        ctx,
		pmService:  pmService,
		appService: appService,
	}
}

func (h *ProjectHandler) GetProjectsHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	projects, err := h.pmService.ProjectService.GetProjects(*c.Request, c.Query("search"), &memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": projects})
}
func (h *ProjectHandler) GetTemplatesHandler(c *gin.Context) {

	c.JSON(200, gin.H{"data": h.appService.CreateDefaultColumnsFromTemplate()})
}

func (h *ProjectHandler) GetProjectHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	id := c.Param("id")
	project, err := h.pmService.ProjectService.GetProjectByID(id, &memberID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": project})
}

func (h *ProjectHandler) CreateProjectHandler(c *gin.Context) {
	var input struct {
		Name        string     `gorm:"type:varchar(255)" json:"name,omitempty"`
		Description string     `json:"description,omitempty"`
		Deadline    *time.Time `json:"deadline,omitempty"`
		Status      string     `json:"status,omitempty"`
		Columns     []struct {
			Name  string `json:"name"`
			Icon  string `json:"icon"`
			Color string `json:"color"`
		} `json:"columns"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var member models.MemberModel
	h.ctx.DB.Where("user_id = ? and company_id = ?", c.MustGet("userID").(string), c.Request.Header.Get("ID-Company")).Find(&member)
	companyID := c.MustGet("companyID").(string)
	userID := c.MustGet("userID").(string)
	project := models.ProjectModel{
		Name:        input.Name,
		Description: input.Description,
		Deadline:    input.Deadline,
		Status:      input.Status,
		CompanyID:   &companyID,
		CreatedByID: &userID,
	}
	if err := h.pmService.ProjectService.CreateProject(&project); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for i, v := range input.Columns {
		if err := h.pmService.ProjectService.CreateColumn(&models.ColumnModel{
			Name:      v.Name,
			Icon:      &v.Icon,
			Color:     &v.Color,
			ProjectID: project.ID,
			Order:     i + 1,
		}); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	h.ctx.DB.Table("project_members").Create(map[string]interface{}{
		"project_model_id": project.ID,
		"member_model_id":  member.ID,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Project created successfully", "project": project})
}

func (h *ProjectHandler) UpdateProjectHandler(c *gin.Context) {
	id := c.Param("id")
	var project models.ProjectModel
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.pmService.ProjectService.UpdateProject(id, &project); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Project updated successfully", "project": project})
}

func (h *ProjectHandler) DeleteProjectHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := h.pmService.ProjectService.GetProjectByID(id, nil)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if err := h.pmService.ProjectService.DeleteProject(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Project deleted successfully"})
}

func (h *ProjectHandler) AddMemberHandler(c *gin.Context) {
	var input struct {
		MemberID string `json:"member_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	projectId := c.Param("id")
	err := h.pmService.ProjectService.AddMemberToProject(projectId, input.MemberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":    "Member added to project successfully",
		"project_id": projectId,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})
	c.JSON(200, gin.H{"message": "Member added to project successfully"})
}
func (h *ProjectHandler) AddNewColumnHandler(c *gin.Context) {
	projectId := c.Param("id")

	var input models.ColumnModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input.ProjectID = projectId

	err := h.pmService.ProjectService.CreateColumn(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	msg := gin.H{
		"message":    "add column to project successfully",
		"project_id": projectId,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})
	c.JSON(200, gin.H{"message": "add column to project successfully"})
}

func (h *ProjectHandler) UpdateColumnHandler(c *gin.Context) {
	projectId := c.Param("id")

	var input models.ColumnModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.pmService.ProjectService.UpdateColumn(input.ID, &input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":    "Column updated successfully",
		"project_id": projectId,
		"column_id":  input.ID,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})
	c.JSON(200, gin.H{"message": "Column updated successfully"})
}

func (h *ProjectHandler) RearrangeColumnsHandler(c *gin.Context) {
	var input struct {
		Columns []models.ColumnModel `json:"columns" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	projectId := c.Param("id")

	for i, v := range input.Columns {
		v.Order = i + 1
		h.ctx.DB.Omit(clause.Associations).Save(&v)
	}

	msg := gin.H{
		"message":    "Column rearrange successfully",
		"project_id": projectId,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})
	c.JSON(200, gin.H{"message": "Column rearrange successfully"})
}
func (h *ProjectHandler) GetMembersHandler(c *gin.Context) {
	projectId := c.Param("id")
	members, err := h.pmService.ProjectService.GetMembersByProjectID(projectId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": members})
}
