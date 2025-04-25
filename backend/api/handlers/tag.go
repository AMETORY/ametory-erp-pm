package handlers

import (
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/tag"
	"github.com/gin-gonic/gin"
)

type TagHandler struct {
	ctx        *context.ERPContext
	tagService *tag.TagService
}

func NewTagHandler(ctx *context.ERPContext) *TagHandler {
	tagService, ok := ctx.TagService.(*tag.TagService)
	if !ok {
		panic("TagService is not instance of tag.TagService")
	}

	return &TagHandler{
		ctx:        ctx,
		tagService: tagService,
	}
}

func (h *TagHandler) CreateTagHandler(c *gin.Context) {
	var input models.TagModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")
	input.CompanyID = &companyID
	err := h.tagService.CreateTag(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute created successfully", "id": input.ID})
}

func (h *TagHandler) GetTagsHandler(c *gin.Context) {
	attributes, err := h.tagService.ListTags(*c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attributes, "message": "Task attributes retrieved successfully"})
}

func (h *TagHandler) GetTagDetailHandler(c *gin.Context) {
	id := c.Param("id")
	attribute, err := h.tagService.GetTagByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": attribute, "message": "Task attribute detail retrieved successfully"})
}

func (h *TagHandler) UpdateTagHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.TagModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.tagService.UpdateTag(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute updated successfully"})
}

func (h *TagHandler) DeleteTagHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := h.tagService.GetTagByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	err = h.tagService.DeleteTag(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task attribute deleted successfully"})
}
