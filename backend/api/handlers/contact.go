package handlers

import (
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type ContactHandler struct {
	ctx                         *context.ERPContext
	contactService              *contact.ContactService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	appService                  *app.AppService
}

func NewContactHandler(ctx *context.ERPContext) *ContactHandler {
	contactService, ok := ctx.ContactService.(*contact.ContactService)
	if !ok {
		panic("invalid contact service")
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}

	var appService *app.AppService
	appSrv, ok := ctx.AppService.(*app.AppService)
	if ok {
		appService = appSrv
	}

	return &ContactHandler{
		ctx:                         ctx,
		contactService:              contactService,
		customerRelationshipService: customerRelationshipService,
		appService:                  appService,
	}
}

func (h *ContactHandler) CreateContactHandler(c *gin.Context) {
	var contact models.ContactModel
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyID := c.GetHeader("ID-Company")
	contact.CompanyID = &companyID

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact created successfully"})
}

func (h *ContactHandler) GetContactHandler(c *gin.Context) {
	id := c.Param("id")

	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contact, "message": "Contact created successfully"})
}

func (h *ContactHandler) SendMessageContactHandler(c *gin.Context) {
	input := struct {
		Message      string `json:"message" binding:"required"`
		Type         string `json:"type" binding:"required"`
		ConnectionID string `json:"connection_id" binding:"required"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	now := time.Now()
	if input.Type == "whatsapp" {
		var conn connection.ConnectionModel
		err = h.ctx.DB.First(&conn, "id = ?", input.ConnectionID).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var whatsappSession models.WhatsappMessageSession
		err = h.ctx.DB.First(&whatsappSession, "contact_id = ?", contact.ID).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			parts := strings.Split(conn.Session, "@")
			userParts := strings.Split(parts[0], ":")
			refType := "connection"
			whatsappSession = models.WhatsappMessageSession{
				ContactID:    &contact.ID,
				JID:          conn.Session,
				Session:      fmt.Sprintf("%s@%s", userParts[0], parts[1]),
				LastOnlineAt: &now,
				LastMessage:  input.Message,
				CompanyID:    conn.CompanyID,
				SessionName:  conn.SessionName,
				RefID:        &conn.ID,
				RefType:      &refType,
			}
			err = h.ctx.DB.Create(&whatsappSession).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

		}

		sendWAMessage(h.ctx, whatsappSession.JID, *contact.Phone, input.Message)
		info := map[string]interface{}{
			"Timestamp": time.Now().Format(time.RFC3339),
		}
		infoByte, err := json.Marshal(info)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		replyResponse := &models.WhatsappMessageModel{
			Receiver:  *contact.Phone,
			Message:   input.Message,
			Session:   whatsappSession.Session,
			JID:       whatsappSession.JID,
			IsFromMe:  true,
			Info:      string(infoByte),
			IsGroup:   false,
			ContactID: &contact.ID,
			CompanyID: conn.CompanyID,
		}

		err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(replyResponse)
		if err != nil {
			// log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		replyResponse.SentAt = &now
		msg := gin.H{
			"message":    replyResponse.Message,
			"command":    "WHATSAPP_RECEIVED",
			"session_id": whatsappSession.ID,
			"data":       replyResponse,
		}
		b, _ := json.Marshal(msg)
		h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
			return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})

	}

	c.JSON(http.StatusOK, gin.H{"data": contact, "message": "Contact Send Message successfully"})
}
func (h *ContactHandler) UpdateContactHandler(c *gin.Context) {
	id := c.Param("id")
	var contact models.ContactModel
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.contactService.UpdateContact(id, &contact)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact created successfully"})
}

func (h *ContactHandler) DeleteContactHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.contactService.DeleteContact(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact deleted"})
}

func (h *ContactHandler) GetContactsHandler(c *gin.Context) {

	isCustomer := true
	contacts, err := h.contactService.GetContacts(*c.Request, c.Query("search"), &isCustomer, nil, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": contacts, "message": "Contact created successfully"})
}
