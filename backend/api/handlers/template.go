package handlers

import (
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
	var memberIDStr = c.MustGet("memberID").(string)
	companyID := c.GetHeader("ID-Company")
	input.CompanyID = &companyID
	userID := c.MustGet("userID").(string)
	input.UserID = &userID
	var IsSuperAdmin bool = IsSuperAdmin(h.ctx, c)
	if !IsSuperAdmin {
		input.MemberID = &memberIDStr
	}
	err := h.customerRelationshipService.WhatsappService.CreateWhatsappMessageTemplate(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Template created successfully", "data": input})
}

func (h *TemplateHandler) GetTemplatesHandler(c *gin.Context) {
	var memberIDStr = c.MustGet("memberID").(string)
	var memberID *string
	memberID = &memberIDStr
	var IsSuperAdmin bool = IsSuperAdmin(h.ctx, c)

	if IsSuperAdmin {

		memberID = nil
	}

	templates, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplates(*c.Request, c.Query("search"), memberID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": templates, "message": "Templates retrieved successfully"})
}

func (h *TemplateHandler) GetTemplateDetailHandler(c *gin.Context) {
	id := c.Param("id")
	template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": template, "message": "Template detail retrieved successfully"})
}

func (h *TemplateHandler) AddMessageHandler(c *gin.Context) {
	var input models.MessageTemplate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = h.customerRelationshipService.WhatsappService.AddMessage(template.ID, &input)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Message added successfully"})

}
func (h *TemplateHandler) DeleteMessageTemplate(c *gin.Context) {

	id := c.Param("id")
	msgId := c.Param("msgId")
	template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = h.customerRelationshipService.WhatsappService.DeleteMessage(template.ID, msgId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Message added successfully"})

}
func (h *TemplateHandler) UpdateTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.WhatsappMessageTemplate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	template, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	for _, v := range template.Messages {
		for _, file := range v.Files {
			h.ctx.DB.Model(&file).Where("id = ?", file.ID).Update("ref_type", "")
			h.ctx.DB.Model(&file).Where("id = ?", file.ID).Update("ref_id", "")
		}
	}

	err = h.customerRelationshipService.WhatsappService.UpdateWhatsappMessageTemplate(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for _, v := range input.Messages {
		for _, file := range v.Files {
			file.RefType = "message_template"
			file.RefID = v.ID
			h.ctx.DB.Save(&file)
		}
		h.ctx.DB.Save(&v)
	}

	c.JSON(200, gin.H{"message": "Template updated successfully"})
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

	c.JSON(200, gin.H{"message": "Template deleted successfully"})
}

func IsSuperAdmin(erpContext *context.ERPContext, c *gin.Context) bool {
	user := c.MustGet("user").(models.UserModel)
	companyID := c.MustGet("companyID").(string)
	erpContext.DB.Preload("Roles").Find(&user)
	for _, v := range user.Roles {
		if v.IsSuperAdmin && *v.CompanyID == companyID {
			return true
		}
	}
	return false
}

func (p *TemplateHandler) AddImageTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	msgId := c.Param("msgId")
	var input models.ProductModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = p.customerRelationshipService.WhatsappService.AddProductWhatsappMessageTemplate(id, msgId, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Product Added successfully"})
}
