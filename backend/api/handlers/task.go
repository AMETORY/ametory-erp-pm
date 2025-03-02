package handlers

import (
	"ametory-pm/services/app"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
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

func (h *TaskHandler) GetTasksHandler(c *gin.Context) {
	projectId := c.Param("id")

	tasks, err := h.pmService.TaskService.GetTasks(*c.Request, c.Query("search"), &projectId)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": tasks, "message": "Tasks retrieved successfully"})
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

	c.JSON(200, gin.H{"message": "Task created successfully"})
}
