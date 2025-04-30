package handlers

import (
	prj "ametory-pm/models/project"
	rapid_api_models "ametory-pm/models/rapid_api"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type TaskHandler struct {
	ctx             *context.ERPContext
	pmService       *project_management.ProjectManagementService
	appService      *app.AppService
	rapidApiService *services.RapidApiService
}

func NewTaskHandler(ctx *context.ERPContext) *TaskHandler {
	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	rapidApiService, ok := ctx.ThirdPartyServices["RapidAPI"].(*services.RapidApiService)
	if !ok {
		panic("RapidApiService is not instance of services.RapidApiService")
	}
	return &TaskHandler{
		ctx:             ctx,
		pmService:       pmService,
		appService:      appService,
		rapidApiService: rapidApiService,
	}
}

func (h *TaskHandler) GetTaskDetailHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")

	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if task.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Task not found in project"})
		return
	}

	var preference prj.ProjectPreferenceModel
	err = h.ctx.DB.First(&preference, "project_id = ?", projectId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			preference.ProjectID = projectId
			h.ctx.DB.Create(&preference)
		}
	}
	c.JSON(200, gin.H{"data": task, "message": "Task retrieved successfully", "preference": preference})
}
func (h *TaskHandler) GetTasksHandler(c *gin.Context) {
	projectId := c.Param("id")

	tasks, err := h.pmService.TaskService.GetTasks(*c.Request, c.Query("search"), &projectId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": tasks, "message": "Tasks retrieved successfully"})
}
func (h *TaskHandler) MoveTaskHandler(c *gin.Context) {
	var input struct {
		ColumnID       string `json:"column_id"`
		SourceColumnID string `json:"source_column_id"`
		OrderNumber    int    `json:"order_number"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	projectId := c.Param("id")
	taskId := c.Param("taskId")
	_, err := h.pmService.ProjectService.GetProjectByID(projectId, nil)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	task.ColumnID = &input.ColumnID
	task.OrderNumber = input.OrderNumber
	err = h.ctx.DB.Omit("Assignee").Save(&task).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":          "Task moved successfully",
		"column_id":        input.ColumnID,
		"source_column_id": input.SourceColumnID,
		"sender_id":        c.MustGet("userID").(string),
		"project_id":       projectId,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &input.ColumnID, &taskId, "MOVE_TASK", nil)
	c.JSON(200, gin.H{"message": "Task moved successfully"})
}

func (h *TaskHandler) RearrangeTaskHandler(c *gin.Context) {
	projectId := c.Param("id")
	var input models.ColumnModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.pmService.ProjectService.GetProjectByID(projectId, nil)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	for i, v := range input.Tasks {
		v.OrderNumber = i + 1
		err = h.ctx.DB.Save(&v).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	msg := gin.H{
		"message":   "Task rearrange successfully",
		"column_id": input.ID,
		"sender_id": c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), &input.ID, nil, "REARRANGE_TASK", nil)
	c.JSON(200, gin.H{"message": "Task rearrange successfully"})

}
func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	projectId := c.Param("id")
	var input models.TaskModel
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	memberID := c.MustGet("member").(models.MemberModel).ID
	input.CreatedByID = &memberID
	input.ProjectID = projectId
	input.StartDate = &now
	input.Status = "ACTIVE"
	input.Priority = "LOW"
	input.Severity = "LOW"
	totalTask, err := h.pmService.TaskService.CountTasksInColumn(*input.ColumnID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	input.OrderNumber = int(totalTask)
	err = h.pmService.TaskService.CreateTask(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":    "Task created successfully",
		"column_id":  input.ColumnID,
		"sender_id":  c.MustGet("userID").(string),
		"project_id": input.ProjectID,
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), input.ColumnID, &input.ID, "CREATE_TASK", nil)
	c.JSON(200, gin.H{"message": "Task created successfully", "task_id": input.ID})
}

func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")
	var input models.TaskModel
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.pmService.ProjectService.GetProjectByID(projectId, nil)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	lastAttributeID := task.TaskAttributeID
	if task.AssigneeID != input.AssigneeID {
		msg := gin.H{
			"task_id":      taskId,
			"message":      "Assignee changed successfully",
			"command":      "RELOAD_TASK",
			"column_id":    task.ColumnID,
			"project_id":   task.ProjectID,
			"sender_id":    c.MustGet("userID").(string),
			"recipient_id": input.AssigneeID,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})

	}

	if input.Completed {
		now := time.Now()
		input.CompletedDate = &now
		input.Percentage = 100
	}
	err = h.ctx.DB.Updates(&input).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if input.TaskAttributeID != lastAttributeID && input.TaskAttributeID != nil {
		var taskAttribute models.TaskAttributeModel
		h.ctx.DB.Find(&taskAttribute, "id = ?", input.TaskAttributeID)
		input.TaskAttibuteData = taskAttribute.Data
		err = h.ctx.DB.Save(&input).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	// fmt.Println("ATTRIBUTE", *input.TaskAttribute)
	if input.TaskAttribute != nil {
		b, _ := json.Marshal(*input.TaskAttribute)
		attrStr := string(b)
		input.TaskAttibuteData = &attrStr
		err = h.ctx.DB.Updates(&input).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	h.ctx.DB.Model(&task).Association("Watchers").Clear()
	var watchers []models.MemberModel

	ids := []string{}
	for _, v := range input.Watchers {
		ids = append(ids, v.ID)
		msg := gin.H{
			"task_id":      taskId,
			"message":      "Watcher changed successfully",
			"command":      "RELOAD_TASK",
			"column_id":    task.ColumnID,
			"project_id":   task.ProjectID,
			"sender_id":    c.MustGet("userID").(string),
			"recipient_id": v.ID,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}
	h.ctx.DB.Find(&watchers, "id in (?)", ids)
	h.ctx.DB.Model(&task).Association("Watchers").Append(watchers)
	// utils.LogJson(input.Watchers)
	msg := gin.H{
		"task_id":    taskId,
		"message":    "Task updated successfully",
		"column_id":  task.ColumnID,
		"project_id": task.ProjectID,
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), task.ColumnID, &taskId, "UPDATE_TASK", nil)
	c.JSON(200, gin.H{"message": "Task updated successfully"})
}

func (h *TaskHandler) AddCommentHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")
	var input struct {
		Comment string `json:"comment" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	memberID := c.MustGet("member").(models.MemberModel).ID
	comment := models.TaskCommentModel{Comment: input.Comment, MemberID: &memberID}
	err = h.pmService.TaskService.CreateComment(taskId, &comment, true)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"task_id":    task.ID,
		"message":    "Comment added successfully",
		"comment_id": comment.ID,
		"project_id": projectId,
		"sender_id":  c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	h.pmService.ProjectService.AddActivity(projectId, c.MustGet("memberID").(string), task.ColumnID, &taskId, "ADD_COMMENT", &input.Comment)

	c.JSON(200, gin.H{"message": "Comment added successfully"})
}

func (h *TaskHandler) MyTaskHandler(c *gin.Context) {
	tasks, err := h.pmService.TaskService.GetMyTask(*c.Request, c.Query("search"), c.MustGet("memberID").(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": tasks, "message": "My Tasks retrieved successfully"})
}
func (h *TaskHandler) WatchedTaskHandler(c *gin.Context) {
	tasks, err := h.pmService.TaskService.GetWatchedTask(*c.Request, c.Query("search"), c.MustGet("memberID").(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": tasks, "message": "My Tasks retrieved successfully"})
}
func (h *TaskHandler) GetTaskPluginsHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")

	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if task.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Task not found in project"})
		return
	}

	companyID := c.GetHeader("ID-Company")
	var plugins []rapid_api_models.RapidApiData
	err = h.ctx.DB.Preload("RapidApiEndpoint").Preload("RapidApiPlugin").Model(&rapid_api_models.RapidApiData{}).Find(&plugins, "company_id = ? and task_id = ? ", companyID, taskId).Error
	if err != nil {
		// c.JSON(500, gin.H{"error": err.Error()})
		fmt.Println(err)
		// return
	}

	for i, v := range plugins {
		if v.Data != "" {
			parsed := map[string]any{}
			json.Unmarshal([]byte(v.Data), &parsed)
			v.ParsedData = parsed
		}
		if v.Params != "" {
			parsedParams := []map[string]any{}
			json.Unmarshal([]byte(v.Params), &parsedParams)
			v.ParsedParams = parsedParams
		}
		v.Data = ""
		plugins[i] = v
	}
	c.JSON(200, gin.H{"data": plugins, "message": "Plugins retrieved successfully"})
}
func (h *TaskHandler) AddPluginHandler(c *gin.Context) {
	input := rapid_api_models.RapidApiData{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	projectId := c.Param("id")
	taskId := c.Param("taskId")

	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if task.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Task not found in project"})
		return
	}
	companyID := c.GetHeader("ID-Company")
	input.ID = utils.Uuid()
	input.Data = "{}"
	input.CompanyID = companyID

	err = h.ctx.DB.Save(&input).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	params := []map[string]any{}
	err = json.Unmarshal([]byte(input.Params), &params)
	if err != nil {
		fmt.Println("ERROR #1", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var plugin rapid_api_models.RapidApiPlugin
	err = h.ctx.DB.Find(&plugin, "id = ?", input.RapidApiPluginID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	var pluginEndpoint rapid_api_models.RapidApiEndpoint
	err = h.ctx.DB.Find(&pluginEndpoint, "id = ?", input.RapidApiEndpointID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.rapidApiService.GetData(plugin, pluginEndpoint, params, companyID)
	if err != nil {
		fmt.Println("ERROR #2", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	b, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("ERROR #3", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	input.Data = string(b)

	h.ctx.DB.Save(&input)

	input.ParsedData = resp

	var history rapid_api_models.RapidApiDataHistory = rapid_api_models.RapidApiDataHistory{
		RapidApiDataID:    input.ID,
		Data:              string(b),
		ChangedByMemberID: c.MustGet("memberID").(string),
		ChangedAt:         time.Now(),
	}
	history.ID = utils.Uuid()
	h.ctx.DB.Save(&history)

	c.JSON(200, gin.H{"data": task, "message": "Task retrieved successfully"})
}

func (h *TaskHandler) GetDataPluginHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")
	pluginDataId := c.Param("pluginId")

	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	if task.ProjectID != projectId {
		c.JSON(404, gin.H{"error": "Task not found in project"})
		return
	}

	pluginData := rapid_api_models.RapidApiData{}
	err = h.ctx.DB.Preload("RapidApiPlugin").Preload("RapidApiEndpoint").Find(&pluginData, "id = ?", pluginDataId).Error
	if err != nil {
		fmt.Println("ERROR #0", err)
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	params := []map[string]any{}
	err = json.Unmarshal([]byte(pluginData.Params), &params)
	if err != nil {
		fmt.Println("ERROR #1", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.rapidApiService.GetData(pluginData.RapidApiPlugin, pluginData.RapidApiEndpoint, params, pluginData.CompanyID)
	if err != nil {
		fmt.Println("ERROR #2", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	b, err := json.Marshal(resp)

	if err != nil {
		fmt.Println("ERROR #3", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	pluginData.Data = string(b)

	h.ctx.DB.Save(&pluginData)

	pluginData.ParsedData = resp

	var history rapid_api_models.RapidApiDataHistory = rapid_api_models.RapidApiDataHistory{

		RapidApiDataID:    pluginDataId,
		Data:              string(b),
		ChangedByMemberID: c.MustGet("memberID").(string),
		ChangedAt:         time.Now(),
	}

	history.ID = utils.Uuid()
	h.ctx.DB.Save(&history)

	c.JSON(200, gin.H{"data": pluginData, "message": "Data retrieved successfully"})
}
