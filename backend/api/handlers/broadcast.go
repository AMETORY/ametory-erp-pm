package handlers

import (
	"ametory-pm/models"
	"ametory-pm/services/app"

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

func (h *BroadcastHandler) CreateBroadcast(c *gin.Context) {
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

func (h *BroadcastHandler) GetBroadcasts(c *gin.Context) {
	broadcasts, err := h.broadcastServ.GetBroadcasts(c.MustGet("companyID").(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, broadcasts)
}

func (h *BroadcastHandler) GetBroadcastByID(c *gin.Context) {
	id := c.Param("id")
	broadcast, err := h.broadcastServ.GetBroadcastByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": broadcast})
}

func (h *BroadcastHandler) UpdateBroadcast(c *gin.Context) {
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

func (h *BroadcastHandler) DeleteBroadcast(c *gin.Context) {
	id := c.Param("id")
	err := h.broadcastServ.DeleteBroadcast(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Broadcast deleted successfully"})
}
