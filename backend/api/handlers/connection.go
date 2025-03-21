package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type ConnectionHandler struct {
	ctx                *context.ERPContext
	appService         *app.AppService
	whatsappWebService *whatsmeow_client.WhatsmeowService
}

func NewConnectionHandler(ctx *context.ERPContext) *ConnectionHandler {
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}

	whatsappWebService, ok := ctx.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
	if !ok {
		panic("ThirdPartyServices is not instance of whatsmeow_client.WhatsmeowService")
	}
	return &ConnectionHandler{
		ctx:                ctx,
		appService:         appService,
		whatsappWebService: whatsappWebService,
	}
}

func (h *ConnectionHandler) GetConnectionsHandler(c *gin.Context) {
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

	connections, err := h.appService.ConnectionService.GetConnections(&pagination, *c.Request, c.Query("search"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": connections, "pagination": pagination, "message": "Connections retrieved successfully"})
}

func (h *ConnectionHandler) GetConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	connection, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": connection, "message": "Connection retrieved successfully"})
}

func (h *ConnectionHandler) CreateConnectionHandler(c *gin.Context) {
	var connection connection.ConnectionModel
	if err := c.ShouldBindJSON(&connection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	connection.APIKey = utils.RandString(32, true)
	connection.Status = "PENDING"
	if err := h.appService.ConnectionService.CreateConnection(&connection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Connection created successfully", "id": connection.ID})
}

func (h *ConnectionHandler) UpdateConnectionHandler(c *gin.Context) {
	var connection connection.ConnectionModel
	if err := c.ShouldBindJSON(&connection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.appService.ConnectionService.UpdateConnection(&connection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully"})
}
func (h *ConnectionHandler) GetQRDeviceHandler(c *gin.Context) {
	session := c.Params.ByName("session")
	resp, err := h.whatsappWebService.GetQR(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Connection successfully", "data": resp})
}
func (h *ConnectionHandler) ConnectDeviceHandler(c *gin.Context) {
	id := c.Param("id")
	connection, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.whatsappWebService.CreateQR(connection.SessionName, fmt.Sprintf("%s/api/v1/whatsapp-webhook", config.App.Server.BaseURL), connection.APIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var respQr any
	if err := json.Unmarshal(resp, &respQr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	connection.Status = "ACTIVE"
	err = h.ctx.DB.Save(&connection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully", "data": respQr})
}

func (h *ConnectionHandler) DeleteConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.appService.ConnectionService.DeleteConnection(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Connection delete successfully"})
}
