package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	tiktok "tiktokshop/open/sdk_golang/service"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ConnectionHandler struct {
	ctx                         *context.ERPContext
	appService                  *app.AppService
	whatsappWebService          *whatsmeow_client.WhatsmeowService
	tiktokService               *tiktok.TiktokService
	shopeeService               *services.ShopeeService
	lazadaService               *services.LazadaService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
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
	tiktokService, ok := ctx.ThirdPartyServices["Tiktok"].(*tiktok.TiktokService)
	if !ok {
		panic("ThirdPartyServices is not instance of tiktok.TiktokService")
	}
	shopeeService, ok := ctx.ThirdPartyServices["Shopee"].(*services.ShopeeService)
	if !ok {
		panic("ThirdPartyServices is not instance of services.ShopeeService")
	}
	lazadaService, ok := ctx.ThirdPartyServices["Lazada"].(*services.LazadaService)
	if !ok {
		panic("ThirdPartyServices is not instance of services.ShopeeService")
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	return &ConnectionHandler{
		ctx:                         ctx,
		appService:                  appService,
		whatsappWebService:          whatsappWebService,
		tiktokService:               tiktokService,
		shopeeService:               shopeeService,
		lazadaService:               lazadaService,
		customerRelationshipService: customerRelationshipService,
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
		if v.Type == "telegram" {
			if v.Status == "PENDING" {
				h.customerRelationshipService.TelegramService.SetToken(&v.SessionName, &v.AccessToken)
				resp, err := h.customerRelationshipService.TelegramService.GetMe()
				if err == nil {
					if resp["ok"].(bool) {
						v.Status = "ACTIVE"
						h.appService.ConnectionService.UpdateConnection(v.ID, &v)
					}
				}
			} else {

				if v.AuthData != nil {
					authData := map[string]any{}
					if err := json.Unmarshal([]byte(*v.AuthData), &authData); err == nil {
						_, ok := authData["webhook"].(string)
						if !ok {
							h.customerRelationshipService.TelegramService.SetToken(&v.SessionName, &v.AccessToken)
							info, err := h.customerRelationshipService.TelegramService.GetWebhookInfo()
							if err == nil {
								fmt.Println("WEBHOOK INFO", info)
							}
							authData["webhook"] = info["result"].(map[string]any)["url"].(string)
							b, _ := json.Marshal(authData)
							jsonData := json.RawMessage(b)
							v.AuthData = &jsonData
							h.appService.ConnectionService.UpdateConnection(v.ID, &v)
						}
					}

				}
			}

		}
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
	if connection.Type == "shopee" && connection.Status == "ACTIVE" {
		// h.shopeeService.GetShopProfile(connection.AccessToken, connection.Username)
	}

	// if connection.Type == "tiktok" && connection.Status == "ACTIVE" {
	// 	authData := map[string]interface{}{}
	// 	if connection.AuthData != nil {
	// 		if err := json.Unmarshal([]byte(*connection.AuthData), &authData); err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 			return
	// 		}
	// 		resp, err := h.tiktokService.CustomerService202309GetConversationsGet(authData["access_token"].(string), connection.Password, "", 10)
	// 		if err != nil {
	// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 			return
	// 		}

	// 		utils.LogJson(resp)
	// 	}

	// }
	c.JSON(http.StatusOK, gin.H{"data": connection, "message": "Connection retrieved successfully"})
}

func (h *ConnectionHandler) CreateConnectionHandler(c *gin.Context) {
	var connection connection.ConnectionModel
	if err := c.ShouldBindJSON(&connection); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")

	if connection.Type == "whatsapp" {
		ok, err := h.appService.ConnectionService.CheckConnectionBySession(connection.Session, companyID, string(connection.Type))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if ok {
			c.JSON(http.StatusOK, gin.H{"data": connection, "message": "Connection already exists"})
			return
		}
	}
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
func (h *ConnectionHandler) AuthorizeConnectionHandler(c *gin.Context) {
	id := c.Param("id")
	conn, err := h.appService.ConnectionService.GetConnection(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var requestBody map[string]interface{}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if requestBody["type"] == "shopee" {

		resp, err := h.shopeeService.GetAuthToken(requestBody["shopee_code"].(string), requestBody["shop_id"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// utils.LogJson(resp)

		expAt := time.Now().Add(time.Duration(resp.ExpireIn) * time.Second)
		conn.Username = requestBody["shop_id"].(string)
		conn.Password = requestBody["shopee_code"].(string)
		conn.AccessToken = resp.AccessToken
		conn.RefreshToken = resp.RefreshToken
		conn.AccessTokenExpiredAt = &expAt
		err = h.ctx.DB.Save(&conn).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		info, err := h.shopeeService.GetShopInfo(resp.AccessToken, conn.Username)
		if err == nil {
			fmt.Println("INFO")
			// utils.LogJson(info)
			b, err := json.Marshal(info)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			conn.SessionName = info.ShopName
			jsonInfo := json.RawMessage(b)
			conn.AuthData = &jsonInfo
			conn.Status = "ACTIVE"
			err = h.ctx.DB.Save(&conn).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully"})
		return
	} else {
		fmt.Println("ERROR SHOPEE", err)
	}

	if requestBody["type"] == "tiktok" {
		resp, err := h.tiktokService.AuthorizeConnection(id, requestBody["tiktok_code"].(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// conn.AccessToken = resp.Data.AccessToken

		b, err := json.Marshal(resp.Data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		data := json.RawMessage(b)
		conn.AuthData = &data
		conn.AccessToken = resp.Data.AccessToken
		conn.RefreshToken = resp.Data.RefreshToken
		expiredAt := time.Unix(int64(resp.Data.AccessTokenExpireIn), 0)
		conn.AccessTokenExpiredAt = &expiredAt
		conn.Status = "ACTIVE"
		err = h.ctx.DB.Save(&conn).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		shops, err := h.tiktokService.Authorization202309GetAuthorizedShopsGet(resp.Data.AccessToken)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully", "data": shops})
		return
	} else {
		fmt.Println("ERROR TIKTOK", err)
	}

	// You can now use requestBody to access all JSON request data

	c.JSON(http.StatusOK, gin.H{"message": "Connection updated successfully", "data": requestBody})
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

	if err := h.appService.ConnectionService.UpdateConnection(id, &input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if input.Type == "whatsapp-api" && input.AccessToken != "" && input.Session != "" {
		input.Status = "ACTIVE"
		if err := h.appService.ConnectionService.UpdateConnection(id, &input); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	}
	// utils.LogJson(input)
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

func (h *ConnectionHandler) GetShopeeAuthURLHandler(c *gin.Context) {
	resp, err := h.shopeeService.AuthShop()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Shopee Auth URL successfully", "data": resp})
}
func (h *ConnectionHandler) GetLazadaAuthURLHandler(c *gin.Context) {
	resp, err := h.lazadaService.GetAuthURL()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Get Lazada Auth URL successfully", "data": resp})
}
