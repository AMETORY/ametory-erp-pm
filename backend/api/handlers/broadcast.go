package handlers

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"net/http"
	"strconv"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

type BroadcastHandler struct {
	ctx           *context.ERPContext
	broadcastServ *app.BroadcastService
}

func NewBroadcastHandler(erpContext *context.ERPContext) *BroadcastHandler {
	broadcastServ, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	if !ok {
		panic("broadcast service not found")
	}
	return &BroadcastHandler{
		ctx:           erpContext,
		broadcastServ: broadcastServ,
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
	c.JSON(200, gin.H{"data": broadcast})
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

	// Logic to send the broadcast, for example using a messaging service
	h.broadcastServ.Send(broadcast)

	c.JSON(200, gin.H{"message": "Broadcast sent successfully"})
}
