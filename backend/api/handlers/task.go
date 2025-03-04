package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type TaskHandler struct {
	ctx        *context.ERPContext
	pmService  *project_management.ProjectManagementService
	appService *app.AppService
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
	return &TaskHandler{
		ctx:        ctx,
		pmService:  pmService,
		appService: appService,
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
	c.JSON(200, gin.H{"data": task, "message": "Task retrieved successfully"})
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
	_, err := h.pmService.ProjectService.GetProjectByID(projectId)
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
	err = h.ctx.DB.Save(&task).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"message":          "Task moved successfully",
		"column_id":        input.ColumnID,
		"source_column_id": input.SourceColumnID,
		"sender_id":        c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})

	c.JSON(200, gin.H{"message": "Task moved successfully"})
}

func (h *TaskHandler) RearrangeTaskHandler(c *gin.Context) {
	projectId := c.Param("id")
	var input models.ColumnModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, err := h.pmService.ProjectService.GetProjectByID(projectId)
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
		return q.Request.URL.Path == url
	})

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
		"message":   "Task created successfully",
		"column_id": input.ColumnID,
		"sender_id": c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})

	c.JSON(200, gin.H{"message": "Task created successfully"})
}

func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	projectId := c.Param("id")
	taskId := c.Param("taskId")
	var input models.TaskModel
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.pmService.ProjectService.GetProjectByID(projectId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	task, err := h.pmService.TaskService.GetTaskByID(taskId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	err = h.ctx.DB.Save(&input).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	msg := gin.H{
		"task_id":   taskId,
		"message":   "Task updated successfully",
		"column_id": task.ColumnID,
		"sender_id": c.MustGet("userID").(string),
	}
	b, _ := json.Marshal(msg)
	h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, c.MustGet("companyID").(string))
		return q.Request.URL.Path == url
	})

	c.JSON(200, gin.H{"message": "Task updated successfully"})
}
