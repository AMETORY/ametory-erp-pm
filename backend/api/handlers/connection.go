package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	for i, v := range connections {
		if v.Type == "whatsapp" {
			resp, err := h.whatsappWebService.CheckConnected(v.Session)
			if err == nil {
				respJson := struct {
					IsConnected bool   `json:"is_connected"`
					Message     string `json:"message"`
				}{}
				if err := json.Unmarshal(resp, &respJson); err == nil {
					fmt.Println("respJson", respJson)
					v.Connected = respJson.IsConnected
				}
			}
		}
		connections[i] = v
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
	resp, err := h.whatsappWebService.CheckConnected(connection.Session)
	if err == nil {
		respJson := struct {
			IsConnected bool   `json:"is_connected"`
			Message     string `json:"message"`
		}{}
		if err := json.Unmarshal(resp, &respJson); err == nil {
			fmt.Println("respJson", respJson)
			connection.Connected = respJson.IsConnected
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": connection, "message": "Connection retrieved successfully"})
}

func (h *ConnectionHandler) CreateConnectionHandler(c *gin.Context) {
	var connection connection.ConnectionModel
	if err := c.ShouldBindJSON(&connection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyID := c.GetHeader("ID-Company")
	connection.APIKey = utils.RandString(32, true)
	connection.Status = "PENDING"
	connection.CompanyID = &companyID
	if err := h.appService.ConnectionService.CreateConnection(&connection); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Connection created successfully", "id": connection.ID})
}

func (h *ConnectionHandler) SyncContactConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	conn, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.whatsappWebService.GetContact(conn.Session, c.Query("search"), c.DefaultQuery("page", "1"), c.DefaultQuery("limit", "5000"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var respJson map[string]interface{}
	if err := json.Unmarshal(resp, &respJson); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	for _, v := range respJson["data"].(map[string]interface{})["items"].([]interface{}) {
		// fmt.Println(v.(map[string]any)["full_name"])
		name := v.(map[string]any)["full_name"].(string)
		if name == "" {
			name = v.(map[string]any)["business_name"].(string)
		}

		phone := v.(map[string]any)["phone_number"].(string)
		var contact models.ContactModel
		err := h.ctx.DB.First(&contact, "phone = ?", phone).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			contact.Phone = &phone
			contact.Name = name
			contact.CompanyID = conn.CompanyID
			contact.IsCustomer = true
			h.ctx.DB.Create(&contact)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": respJson})
}
func (h *ConnectionHandler) UpdateConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input connection.ConnectionModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.ConnectionService.UpdateConnection(&input); err != nil {
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
	baseURL := config.App.Server.BaseURL
	resp, err := h.whatsappWebService.CreateQR(connection.SessionName, fmt.Sprintf("%s/api/v1/whatsapp-webhook", baseURL), connection.APIKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var respQr map[string]any
	if err := json.Unmarshal(resp, &respQr); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// utils.LogJson(respQr)
	connection.Status = "ACTIVE"
	respData, ok := respQr["data"].(map[string]any)
	if ok {
		jid := respData["jid"].(string)
		connection.Session = jid
	}
	err = h.ctx.DB.Save(&connection).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully", "data": respQr})
}

func (h *ConnectionHandler) DeleteConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	connection, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.whatsappWebService.DisconnectDeviceByJID(connection.Session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.appService.ConnectionService.DeleteConnection(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Connection delete successfully"})
}
