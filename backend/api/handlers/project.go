package handlers

import (
	"ametory-pm/services/app"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
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
	projects, err := h.pmService.ProjectService.GetProjects(*c.Request, c.Query("search"))
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
	id := c.Param("id")
	project, err := h.pmService.ProjectService.GetProjectByID(id)
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
	companyID := c.MustGet("companyID").(string)
	member := c.MustGet("member").(models.MemberModel)
	project := models.ProjectModel{
		Name:        input.Name,
		Description: input.Description,
		Deadline:    input.Deadline,
		Status:      input.Status,
		CompanyID:   &companyID,
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

	err := h.ctx.DB.Model(&project).Association("Members").Append(&member)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

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
	_, err := h.pmService.ProjectService.GetProjectByID(id)
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
