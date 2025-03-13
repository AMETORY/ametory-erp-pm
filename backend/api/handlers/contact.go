package handlers

import (
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	ctx            *context.ERPContext
	contactService *contact.ContactService
}

func NewContactHandler(ctx *context.ERPContext) *ContactHandler {
	contactService, ok := ctx.ContactService.(*contact.ContactService)
	if !ok {
		panic("invalid contact service")
	}

	return &ContactHandler{
		ctx:            ctx,
		contactService: contactService,
	}
}

func (h *ContactHandler) CreateContactHandler(c *gin.Context) {
	var contact models.ContactModel
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.contactService.CreateContact(&contact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func (h *ContactHandler) GetContactHandler(c *gin.Context) {
	id := c.Param("id")

	contact, err := h.contactService.GetContactByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contact)
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

	c.JSON(http.StatusOK, contact)
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

	var isCustomer, isVendor, isSupplier = false, false, false
	isCustomer = true
	contacts, err := h.contactService.GetContacts(*c.Request, c.Query("search"), &isCustomer, &isVendor, &isSupplier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contacts)
}
