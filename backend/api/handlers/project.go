package handlers

import (
	prj "ametory-pm/models/project"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
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
	c.JSON(200, gin.H{"data": projects, "message": "Projects retrieved successfully"})
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

	var preference prj.ProjectPreferenceModel
	err = h.ctx.DB.First(&preference, "project_id = ?", project.ID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			preference.ProjectID = project.ID
			h.ctx.DB.Create(&preference)
		}
	}
	c.JSON(200, gin.H{"data": project, "preference": preference})
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

func (h *ProjectHandler) UpdateProjectPreferenceHandler(c *gin.Context) {
	memberID := c.MustGet("memberID").(string)
	id := c.Param("id")
	project, err := h.pmService.ProjectService.GetProjectByID(id, &memberID)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var preference prj.ProjectPreferenceModel
	if err := c.ShouldBindJSON(&preference); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(preference)
	if err := h.ctx.DB.Where("project_id = ?", project.ID).Model(&prj.ProjectPreferenceModel{}).Updates(map[string]any{
		"rapid_api_enabled":        preference.RapidApiEnabled.Bool,
		"contact_enabled":          preference.ContactEnabled.Bool,
		"custom_attribute_enabled": preference.CustomAttributeEnabled.Bool,
		"gemini_enabled":           preference.GeminiEnabled.Bool,
		"form_enabled":             preference.FormEnabled.Bool,
	}).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Project preference updated successfully"})

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

	msg := gin.H{
		"message":    "Project updated successfully",
		"project_id": id,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	h.pmService.ProjectService.AddActivity(id, c.MustGet("memberID").(string), nil, nil, "UPDATE_PROJECT", nil)
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
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), nil, nil, "ADD_MEMBER", nil)
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
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), nil, nil, "ADD_COLUMN", nil)
	c.JSON(200, gin.H{"message": "add column to project successfully"})
}

func (h *ProjectHandler) GetColumnByIDHandler(c *gin.Context) {
	projectId := c.Param("id")
	columnId := c.Param("columnId")

	column, err := h.pmService.ProjectService.GetColumnByID(columnId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Column not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	if column.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Column not found in project"})
		return
	}

	c.JSON(200, gin.H{"data": column, "message": "Column retrieved successfully"})
}

func (h *ProjectHandler) DeleteColumnActionHandler(c *gin.Context) {
	projectId := c.Param("id")
	columnId := c.Param("columnId")
	actionId := c.Param("actionId")

	column, err := h.pmService.ProjectService.GetColumnByID(columnId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, gin.H{"error": "Column not found"})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	if column.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Column not found in project"})
		return
	}

	err = h.pmService.ProjectService.DeleteColumnAction(actionId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":    "Column action deleted successfully",
		"project_id": projectId,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &columnId, nil, "DELETE_COLUMN_ACTION", nil)
	c.JSON(200, gin.H{"message": "Column action deleted successfully"})
}
func (h *ProjectHandler) EditColumnActionHandler(c *gin.Context) {
	projectId := c.Param("id")
	columnId := c.Param("columnId")

	var input models.ColumnAction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.pmService.ProjectService.UpdateColumnAction(input.ID, &input); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":    "Column action edited successfully",
		"project_id": projectId,
		"column_id":  columnId,
		"command":    "RELOAD",
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &columnId, nil, "EDIT_COLUMN_ACTION", nil)
	c.JSON(200, gin.H{"message": "Column action edited successfully"})
}

func (h *ProjectHandler) AddNewColumnActionHandler(c *gin.Context) {
	projectId := c.Param("id")
	columnId := c.Param("columnId")
	// fmt.Println("Column ID", columnId)
	var input models.ColumnAction
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	column, err := h.pmService.ProjectService.GetColumnByID(columnId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if column.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Column not found in project"})
		return
	}
	input.ColumnID = columnId

	err = h.pmService.ProjectService.CreateColumnAction(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// msg := gin.H{
	// 	"message":    "add column action to project successfully",
	// 	"project_id": projectId,
	// 	"column_id":  columnId,
	// 	"command":    "RELOAD",
	// 	"sender_id":  c.MustGet("userID").(string),
	// }
	// b, _ := json.Marshal(msg)
	// h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
	// 	url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
	// 	return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	// })
	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &columnId, nil, "ADD_COLUMN_ACTION", nil)
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
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &input.ID, nil, "UPDATE_COLUMN", nil)
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
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})
	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), nil, nil, "REARRANGE_COLUMN", nil)
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
func (h *ProjectHandler) CountCompletedTasksHandler(c *gin.Context) {
	projectId := c.Param("id")
	var count int64

	days, _ := strconv.Atoi(c.Query("days"))
	if days == 0 {
		days = 7
	}
	// Query to count tasks completed in the last days for the given project
	err := h.ctx.DB.Model(&models.TaskModel{}).
		Where("project_id = ? AND completed = ? AND completed_date >= ?", projectId, true, time.Now().AddDate(0, 0, -days)).
		Count(&count).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count, "message": "Completed tasks counted successfully"})
}

func (h *ProjectHandler) CountUpdatedTasksHandler(c *gin.Context) {
	projectId := c.Param("id")
	var count int64

	days, _ := strconv.Atoi(c.Query("days"))
	if days == 0 {
		days = 7
	}
	// Query to count tasks updated in the last days for the given project
	err := h.ctx.DB.Model(&models.TaskModel{}).
		Where("project_id = ? AND updated_at >= ?", projectId, time.Now().AddDate(0, 0, -days)).
		Count(&count).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count, "message": "Updated tasks counted successfully"})
}

func (h *ProjectHandler) CountCreatedTasksHandler(c *gin.Context) {
	projectId := c.Param("id")
	var count int64

	days, _ := strconv.Atoi(c.Query("days"))
	if days == 0 {
		days = 7
	}
	// Query to count tasks created in the last days for the given project
	err := h.ctx.DB.Model(&models.TaskModel{}).
		Where("project_id = ? AND created_at >= ?", projectId, time.Now().AddDate(0, 0, -days)).
		Count(&count).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count, "message": "Created tasks counted successfully"})
}

func (h *ProjectHandler) CountDueTasksHandler(c *gin.Context) {
	projectId := c.Param("id")
	var count int64

	days, _ := strconv.Atoi(c.Query("days"))
	if days == 0 {
		days = 7
	}
	// Query to count tasks whose due date is in the last days for the given project
	err := h.ctx.DB.Model(&models.TaskModel{}).
		Where("project_id = ? AND end_date <= ?", projectId, time.Now().AddDate(0, 0, days)).
		Count(&count).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count, "message": "Due tasks counted successfully"})
}

func (h *ProjectHandler) CountNextDueTasksHandler(c *gin.Context) {
	projectId := c.Param("id")
	var count int64

	days, _ := strconv.Atoi(c.Query("days"))
	if days == 0 {
		days = 7
	}
	// Query to count tasks whose due date is in the next days for the given project
	err := h.ctx.DB.Model(&models.TaskModel{}).
		Where("project_id = ? AND end_date > ? AND end_date <= ?", projectId, time.Now(), time.Now().AddDate(0, 0, days)).
		Count(&count).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"count": count, "message": "Next due tasks counted successfully"})
}

func (h *ProjectHandler) CountColumnTasksHandler(c *gin.Context) {
	projectId := c.Param("id")

	var columns []models.ColumnModel

	// Query to retrieve columns in the given project
	err := h.ctx.DB.Model(&models.ColumnModel{}).
		Where("project_id = ?", projectId).
		Find(&columns).Error

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for i, column := range columns {
		var count int64

		// Query to count tasks in the given column
		err := h.ctx.DB.Model(&models.TaskModel{}).
			Where("project_id = ? AND column_id = ?", projectId, column.ID).
			Count(&count).Error

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		column.CountTasks = count
		columns[i] = column
	}

	c.JSON(200, gin.H{"data": columns, "message": "Tasks in the columns counted successfully"})

}

func (h *ProjectHandler) GetRecentActivities(c *gin.Context) {
	projectId := c.Param("id")
	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit == 0 {
		limit = 10
	}
	activities, err := h.pmService.ProjectService.GetRecentActivities(projectId, limit)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": activities, "message": "Recent activities retrieved successfully"})
}

func (h *ProjectHandler) CountTasksByPriorityHandler(c *gin.Context) {
	projectId := c.Param("id")
	priorities := []string{"LOW", "MEDIUM", "HIGH", "URGENT"}
	var result []map[string]any

	// Query to count tasks based on priority for the given project
	for _, priority := range priorities {
		var count int64
		err := h.ctx.DB.Model(&models.TaskModel{}).
			Where("project_id = ? AND priority = ?", projectId, priority).
			Count(&count).Error

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		result = append(result, map[string]any{
			"value": priority,
			"label": strings.ToTitle(priority),
			"color": getPriorityColor(priority),
			"count": count,
		})
	}

	c.JSON(200, gin.H{"data": result, "message": "Tasks counted by priority successfully"})
}

func getPriorityColor(priority string) string {
	switch priority {
	case "LOW":
		return "#8BC34A"
	case "MEDIUM":
		return "#F7DC6F"
	case "HIGH":
		return "#FFC107"
	case "URGENT":
		return "#F44336"
	default:
		return "#000000"
	}
}

func (h *ProjectHandler) CountTasksBySeverityHandler(c *gin.Context) {
	projectId := c.Param("id")
	severities := []string{"LOW", "MEDIUM", "HIGH", "CRITICAL"}
	var result []map[string]any

	// Query to count tasks based on severity for the given project
	for _, severity := range severities {
		var count int64
		err := h.ctx.DB.Model(&models.TaskModel{}).
			Where("project_id = ? AND severity = ?", projectId, severity).
			Count(&count).Error

		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		result = append(result, map[string]any{
			"value": severity,
			"label": strings.ToTitle(severity),
			"color": getSeverityColor(severity),
			"count": count,
		})
	}

	c.JSON(200, gin.H{"data": result, "message": "Tasks counted by severity successfully"})
}

func getSeverityColor(severity string) string {
	switch severity {
	case "LOW":
		return "#8BC34A"
	case "MEDIUM":
		return "#F7DC6F"
	case "HIGH":
		return "#FFC107"
	case "CRITICAL":
		return "#F44336"
	default:
		return "#000000"
	}
}
