package handlers

import (
	"ametory-pm/api/middlewares"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type TemplateHandler struct {
	ctx                         *context.ERPContext
	customerRelationshipService *customer_relationship.CustomerRelationshipService
}

func NewTemplateHandler(ctx *context.ERPContext) *TemplateHandler {
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}

	return &TemplateHandler{
		ctx:                         ctx,
		customerRelationshipService: customerRelationshipService,
	}
}

func (h *TemplateHandler) CreateTemplateHandler(c *gin.Context) {
	var input models.WhatsappMessageTemplate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")
	input.CompanyID = &companyID
	err := h.customerRelationshipService.WhatsappService.CreateWhatsappMessageTemplate(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute created successfully", "id": input.ID})
}

func (h *TemplateHandler) GetTemplatesHandler(c *gin.Context) {
	var memberIDStr = c.MustGet("memberID").(string)
	var memberID *string
	ok, _ := middlewares.CheckIsSuperAdminPermission(c.MustGet("user").(models.UserModel))
	if ok {
		memberID = nil

	}
	memberID = &memberIDStr
	attributes, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplates(*c.Request, c.Query("search"), memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attributes, "message": "Task attributes retrieved successfully"})
}

func (h *TemplateHandler) GetTemplateDetailHandler(c *gin.Context) {
	id := c.Param("id")
	attribute, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attribute, "message": "Task attribute detail retrieved successfully"})
}

func (h *TemplateHandler) UpdateTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.WhatsappMessageTemplate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.customerRelationshipService.WhatsappService.UpdateWhatsappMessageTemplate(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute updated successfully"})
}

func (h *TemplateHandler) DeleteTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = h.customerRelationshipService.WhatsappService.DeleteWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute deleted successfully"})
}
