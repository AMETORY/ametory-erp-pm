package handlers

import (
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	ctx        *context.ERPContext
	csrService *customer_relationship.CustomerRelationshipService
}

func NewFormHandler(ctx *context.ERPContext) *FormHandler {
	csrSevice, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if !ok {
		panic("CustomerRelationshipService is not found")
	}
	return &FormHandler{ctx: ctx, csrService: csrSevice}
}

// Add CRUD operations for FormTemplate

// CreateFormTemplateHandler handles the creation of a new form template
func (h *FormHandler) CreateFormTemplateHandler(c *gin.Context) {
	var formTemplate models.FormTemplate
	if err := c.ShouldBindJSON(&formTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	memberID := c.MustGet("memberID").(string)
	userID := c.MustGet("userID").(string)
	companyID := c.GetHeader("ID-Company")
	formTemplate.CreatedByMemberID = &memberID
	formTemplate.CreatedByID = &userID
	formTemplate.CompanyID = &companyID
	if err := h.csrService.FormService.CreateFormTemplate(&formTemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": formTemplate, "message": "Form template created successfully", "id": formTemplate.ID})
}

// GetFormTemplatesHandler retrieves all form templates
func (h *FormHandler) GetFormTemplatesHandler(c *gin.Context) {
	formTemplates, err := h.csrService.FormService.GetFormTemplates(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": formTemplates, "message": "Form templates retrieved successfully"})
}

// GetFormTemplateHandler retrieves a form template by ID
func (h *FormHandler) GetFormTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	formTemplate, err := h.csrService.FormService.GetFormTemplate(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": formTemplate, "message": "Form template retrieved successfully"})
}

// UpdateFormTemplateHandler updates a form template by ID
func (h *FormHandler) UpdateFormTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	var formTemplate models.FormTemplate
	if err := c.ShouldBindJSON(&formTemplate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.csrService.FormService.UpdateFormTemplate(id, &formTemplate); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": formTemplate})
}

// DeleteFormTemplateHandler deletes a form template by ID
func (h *FormHandler) DeleteFormTemplateHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.csrService.FormService.DeleteFormTemplate(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Form template deleted successfully"})
}

// CreateFormHandler handles the creation of a new form
func (h *FormHandler) CreateFormHandler(c *gin.Context) {
	var form models.FormModel
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.csrService.FormService.CreateForm(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Form created successfully"})
}

// GetFormsHandler retrieves all forms
func (h *FormHandler) GetFormsHandler(c *gin.Context) {
	form, err := h.csrService.FormService.GetForms(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": form, "message": "Forms retrieved successfully"})
}

// GetFormHandler retrieves a form by ID
func (h *FormHandler) GetFormHandler(c *gin.Context) {
	id := c.Param("id")
	form, err := h.csrService.FormService.GetForm(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": form, "message": "Form retrieved successfully"})
}

// UpdateFormHandler updates a form by ID
func (h *FormHandler) UpdateFormHandler(c *gin.Context) {
	id := c.Param("id")
	var form models.FormModel
	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.csrService.FormService.UpdateForm(id, &form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Form updated successfully"})
}

// DeleteFormHandler deletes a form by ID
func (h *FormHandler) DeleteFormHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.csrService.FormService.DeleteForm(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Form deleted successfully"})
}
