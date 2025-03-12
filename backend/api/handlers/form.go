package handlers

import (
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/file"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type FormHandler struct {
	ctx         *context.ERPContext
	csrService  *customer_relationship.CustomerRelationshipService
	fileService *file.FileService
	pmService   *project_management.ProjectManagementService
	appService  *app.AppService
}

func NewFormHandler(ctx *context.ERPContext) *FormHandler {
	csrSevice, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if !ok {
		panic("CustomerRelationshipService is not found")
	}
	fileService, ok := ctx.FileService.(*file.FileService)
	if !ok {
		panic("FileService is not instance of file.FileService")
	}
	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	return &FormHandler{ctx: ctx, csrService: csrSevice, fileService: fileService, pmService: pmService, appService: appService}
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

	if form.Headers == "" {
		form.Headers = "{}"
	}
	companyID := c.GetHeader("ID-Company")
	memberID := c.MustGet("memberID").(string)
	userID := c.MustGet("userID").(string)
	form.Status = "ACTIVE"
	form.CompanyID = &companyID
	form.CreatedByMemberID = &memberID
	form.CreatedByID = &userID
	if err := h.csrService.FormService.CreateForm(&form); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if form.Cover != nil {
		form.Cover.RefID = form.ID
		form.Cover.RefType = "form-cover"
		h.ctx.DB.Save(&form.Cover)
	}
	if form.Picture != nil {
		form.Picture.RefID = form.ID
		form.Picture.RefType = "form-picture"
		h.ctx.DB.Save(&form.Picture)
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Form created successfully", "id": form.ID})
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

	if form.Cover != nil {
		form.Cover.RefID = form.ID
		form.Cover.RefType = "form-cover"
		h.ctx.DB.Save(&form.Cover)
	}
	if form.Picture != nil {
		form.Picture.RefID = form.ID
		form.Picture.RefType = "form-picture"
		h.ctx.DB.Save(&form.Picture)
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

func (h *FormHandler) PublicFormHandler(c *gin.Context) {
	formCode := c.Param("formCode")

	form, err := h.csrService.FormService.GetFormByCode(formCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if !form.IsPublic {
		c.JSON(http.StatusForbidden, gin.H{"error": "Form is not public"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": form, "message": "Form retrieved successfully"})

}

func (h *FormHandler) FormResponseHandler(c *gin.Context) {
	formCode := c.Param("formCode")

	form, err := h.csrService.FormService.GetFormByCode(formCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	input := []any{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, s := range input {
		section := s.(map[string]any)
		for j, f := range section["fields"].([]any) {
			field := f.(map[string]any)
			if field["type"].(string) == string(models.FileUpload) {
				file := models.FileModel{
					FileName: utils.RandString(8, false),
				}
				base64str := strings.Split(field["value"].(string), ",")[1]
				h.fileService.UploadFileFromBase64(base64str, "local", "files", &file)

				input[i].(map[string]any)["fields"].([]any)[j].(map[string]any)["value"] = file.URL
			}
		}
	}

	b, _ := json.Marshal(input)

	var formResponse models.FormResponseModel = models.FormResponseModel{
		FormID:   form.ID,
		Data:     string(b),
		Metadata: "{}",
	}

	formResponse.ID = utils.Uuid()

	err = h.ctx.DB.Save(&formResponse).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if form.ColumnID != nil {
		formResponse.RefID = *form.ColumnID
		formResponse.RefType = "column"

		task := models.TaskModel{
			Name:           fmt.Sprintf("From Form - %s", form.Title),
			ProjectID:      *form.ProjectID,
			ColumnID:       form.ColumnID,
			StartDate:      &formResponse.CreatedAt,
			EndDate:        &formResponse.CreatedAt,
			FormResponseID: &formResponse.ID,
			AssigneeID:     form.CreatedByMemberID,
		}
		err = h.pmService.TaskService.CreateTask(&task)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		msg := gin.H{
			"message":   "Task created successfully",
			"column_id": form.ColumnID,
			"sender_id": form.CreatedByID,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *form.CompanyID)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}

	if form.SubmitURL != "" {
		headers := map[string]string{}
		json.Unmarshal([]byte(form.Headers), &headers)
		_, err := sendHTTPRequest(form.Method, form.SubmitURL, headers, input)
		if err != nil {
			fmt.Println("ERROR", err)
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			// return
		}
		// resp.Body

	}

	c.JSON(http.StatusOK, gin.H{"data": form, "message": " submitted successfully"})
}

func sendHTTPRequest(method, url string, headers map[string]string, body any) (*http.Response, error) {
	fmt.Println("SEND REQUEST TO", url)
	client := &http.Client{}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}
