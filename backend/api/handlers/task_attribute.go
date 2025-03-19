package handlers

import (
	"ametory-pm/services/app"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type TaskAttibuteHandler struct {
	ctx        *context.ERPContext
	pmService  *project_management.ProjectManagementService
	appService *app.AppService
}

func NewTasAttributekHandler(ctx *context.ERPContext) *TaskAttibuteHandler {
	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}

	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}

	return &TaskAttibuteHandler{
		ctx:        ctx,
		pmService:  pmService,
		appService: appService,
	}
}

func (h *TaskAttibuteHandler) CreateTaskAttributeHandler(c *gin.Context) {
	var input models.TaskAttributeModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.pmService.TaskAttributeService.CreateTaskAttribute(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute created successfully", "id": input.ID})
}

func (h *TaskAttibuteHandler) GetTaskAttributesHandler(c *gin.Context) {
	attributes, err := h.pmService.TaskAttributeService.GetTaskAttributes(c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attributes, "message": "Task attributes retrieved successfully"})
}

func (h *TaskAttibuteHandler) GetTaskAttributeDetailHandler(c *gin.Context) {
	id := c.Param("id")
	attribute, err := h.pmService.TaskAttributeService.GetTaskAttributeByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attribute, "message": "Task attribute detail retrieved successfully"})
}

func (h *TaskAttibuteHandler) UpdateTaskAttributeHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.TaskAttributeModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.pmService.TaskAttributeService.UpdateTaskAttribute(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute updated successfully"})
}

func (h *TaskAttibuteHandler) DeleteTaskAttributeHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := h.pmService.TaskAttributeService.GetTaskAttributeByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = h.pmService.TaskAttributeService.DeleteTaskAttribute(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute deleted successfully"})
}
