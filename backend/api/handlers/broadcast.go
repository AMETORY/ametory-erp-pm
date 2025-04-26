package handlers

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"net/http"
	"strconv"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type BroadcastHandler struct {
	ctx           *context.ERPContext
	broadcastServ *app.BroadcastService
	contactSrv    *contact.ContactService
}

func NewBroadcastHandler(erpContext *context.ERPContext) *BroadcastHandler {
	broadcastServ, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	if !ok {
		panic("broadcast service not found")
	}

	contactSrv, ok := erpContext.ContactService.(*contact.ContactService)
	if !ok {
		panic("contact service not found")
	}

	return &BroadcastHandler{
		ctx:           erpContext,
		broadcastServ: broadcastServ,
		contactSrv:    contactSrv,
	}
}

func (h *BroadcastHandler) GetBroadcastsHandler(c *gin.Context) {
	var pagination app.Pagination

	limitStr := c.DefaultQuery("size", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pagination.Limit = limit
	pagination.Page = page

	broadcasts, err := h.broadcastServ.GetBroadcasts(&pagination, *c.Request, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": broadcasts, "pagination": pagination, "message": "Broadcasts retrieved successfully"})
}

func (h *BroadcastHandler) CreateBroadcastHandler(c *gin.Context) {
	var input models.BroadcastModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")
	input.CompanyID = &companyID
	err := h.broadcastServ.CreateBroadcast(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Broadcast created", "data": input})
}

func (h *BroadcastHandler) GetBroadcastHandler(c *gin.Context) {
	id := c.Param("id")
	broadcast, err := h.broadcastServ.GetBroadcastByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var pagination app.Pagination

	limitStr := c.DefaultQuery("size", "10")
	pageStr := c.DefaultQuery("page", "1")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	pagination.Limit = limit
	pagination.Page = page

	contacts, err := h.broadcastServ.GetContacts(id, &pagination, c.Query("search"))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	broadcast.Contacts = contacts
	broadcast.ContactCount = int(pagination.TotalRows)
	c.JSON(200, gin.H{"data": broadcast, "pagination": pagination, "message": "Broadcast retrieved successfully"})
}

func (h *BroadcastHandler) UpdateBroadcastHandler(c *gin.Context) {
	id := c.Param("id")
	var input models.BroadcastModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.broadcastServ.UpdateBroadcast(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Broadcast updated successfully"})
}

func (h *BroadcastHandler) DeleteBroadcastHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.broadcastServ.DeleteBroadcast(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Broadcast deleted successfully"})
}

func (h *BroadcastHandler) SendBroadcastHandler(c *gin.Context) {
	id := c.Param("id")
	broadcast, err := h.broadcastServ.GetBroadcastByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve broadcast"})
		return
	}

	h.ctx.DB.Preload(clause.Associations).Find(&broadcast)

	// Logic to send the broadcast, for example using a messaging service
	h.broadcastServ.Send(broadcast)

	c.JSON(200, gin.H{"message": "Broadcast sent successfully"})
}

func (h *BroadcastHandler) AddContactHandler(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		TagIDs     []string `json:"tag_ids"`
		ContactIDs []string `json:"contact_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contacts, err := h.contactSrv.GetContactByTagIDs(input.TagIDs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	contactIDs := []mdl.ContactModel{}
	for _, v := range input.ContactIDs {
		contactIDs = append(contactIDs, mdl.ContactModel{BaseModel: shared.BaseModel{ID: v}})
	}

	contacts = append(contacts, contactIDs...)

	if err := h.broadcastServ.AddContact(id, contacts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contacts added successfully"})
}

func (h *BroadcastHandler) DeleteContactHandler(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		TagIDs     []string `json:"tag_ids"`
		ContactIDs []string `json:"contact_ids"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.broadcastServ.DeleteContactByIDs(id, input.ContactIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted successfully"})
}
