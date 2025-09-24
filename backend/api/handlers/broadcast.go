package handlers

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BroadcastHandler struct {
	ctx           *context.ERPContext
	broadcastServ *app.BroadcastService
	contactSrv    *contact.ContactService
	waService     *whatsmeow_client.WhatsmeowService
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
	var waService *whatsmeow_client.WhatsmeowService
	waSrv, ok := erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
	if ok {
		waService = waSrv
	}

	return &BroadcastHandler{
		ctx:           erpContext,
		broadcastServ: broadcastServ,
		contactSrv:    contactSrv,
		waService:     waService,
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
	memberID := c.MustGet("memberID").(string)
	input.MemberID = &memberID
	input.DelayTime = 1
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

	template, ok := broadcast.Template.(mdl.WhatsappMessageTemplate)
	if ok {
		for i, message := range template.Messages {
			for j, v := range message.Products {
				var images []mdl.FileModel
				h.ctx.DB.Where("ref_id = ? and ref_type = ?", v.ID, "product").Find(&images)
				v.ProductImages = images
				template.Messages[i].Products[j] = v
			}

		}
	}
	broadcast.Template = template
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

	if len(input.Connections) == 0 {
		fmt.Println("CLEAR CONNECTIOn")
		err := h.ctx.DB.Exec("DELETE FROM broadcast_connections where broadcast_model_id = ?", id).Error
		if err != nil {
			fmt.Println("ERROR CLEAR CONNECTION", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}
	for _, v := range input.Files {
		v.RefType = "broadcast"
		v.RefID = id
		h.ctx.DB.Save(&v)
	}

	err = h.ctx.DB.Model(&input).Association("Products").Replace(input.Products)
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

func (h *BroadcastHandler) AddContactFromFileHandler(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		FileURL string `json:"file_url" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contacts, err := h.importFile(map[string]string{
		"file_url":   input.FileURL,
		"user_id":    c.MustGet("userID").(string),
		"company_id": c.GetHeader("ID-Company"),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// utils.LogJson(contacts)
	fmt.Println("contacts", len(contacts))
	if err := h.broadcastServ.AddContact(id, contacts); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contacts imported successfully", "data": input.FileURL})
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

func (h *BroadcastHandler) importFile(data map[string]string) ([]mdl.ContactModel, error) {
	resp, err := http.Get(data["file_url"])
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	file, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	xlsx, err := excelize.OpenReader(bytes.NewReader(file))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	firstSheet := xlsx.GetSheetName(0)
	rows, err := xlsx.GetRows(firstSheet)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var contacts []mdl.ContactModel

	// var customDataKey = ""
	for i, row := range rows {
		if i == 0 {
			if len(rows[i]) > 5 && cleanString(rows[i][5]) != "" {
				fmt.Println("------------------------------------------------")
				fmt.Println(rows[i])
				// customDataKey = strings.ToLower(cleanString(rows[i][5]))
			}
			continue
		}
		for j, r := range row {
			fmt.Printf("%v.\t %s: %s\n", j+1, rows[0][j], r)
		}
		fmt.Println("------------------------------------------------")
		var phone *string
		if cleanString(rows[i][2]) != "" {
			phoneStr := utils.ParsePhoneNumber(cleanString(rows[i][2]), "")
			phone = &phoneStr
		}

		userID := data["user_id"]
		companyID := data["company_id"]

		var tags []mdl.TagModel
		tagStr := "IMPORT"
		if len(rows[i]) > 4 && rows[i][4] != "" {
			tagStr = "IMPORT," + cleanString(rows[i][4])
		}

		if cleanString(tagStr) != "" {
			dataTags := strings.Split(cleanString(tagStr), ",")
			for _, v := range dataTags {
				fmt.Println("TAG", cleanString(v))
				var checkTag mdl.TagModel
				err := h.ctx.DB.Where("name = ? and company_id = ?", cleanString(v), companyID).First(&checkTag).Error
				if err != nil {
					if errors.Is(err, gorm.ErrRecordNotFound) {
						tag := mdl.TagModel{
							Name:      cleanString(v),
							CompanyID: &companyID,
							Color:     randomColor(),
						}
						err := h.ctx.DB.Create(&tag).Error
						if err != nil {
							log.Println(err)
							continue
						}
						tags = append(tags, tag)
					}
				} else {
					tags = append(tags, checkTag)
				}
			}
		}
		// var products []mdl.ProductModel
		// if len(rows[i]) > 5 && cleanString(rows[i][5]) != "" {
		// 	dataProducts := strings.Split(cleanString(rows[i][5]), ",")
		// 	for _, v := range dataProducts {
		// 		fmt.Println("PRODUCT", cleanString(v))
		// 		var checkProduct mdl.ProductModel
		// 		err := h.ctx.DB.Where("name = ? or sku = ?", cleanString(v), cleanString(v)).Where("company_id = ?", companyID).First(&checkProduct).Error
		// 		if err == nil {
		// 			products = append(products, checkProduct)
		// 		}
		// 	}
		// }
		var customData = make(map[string]string)
		// if len(rows[i]) > 5 && cleanString(rows[i][5]) != "" && customDataKey != "" {
		// 	customData[customDataKey] = cleanString(rows[i][5])

		// }

		if len(rows[i]) > 5 {
			for j := 5; j < len(rows[i]); j++ {
				customData[rows[0][j]] = cleanString(rows[i][j])
			}
		}

		b, _ := json.Marshal(customData)

		newContact := mdl.ContactModel{
			Name:      rows[i][0],
			Email:     rows[i][1],
			Phone:     phone,
			UserID:    &userID,
			CompanyID: &companyID,
			Tags:      tags,
			// Products:   products,
			IsCustomer: true,
			CustomData: b,
		}

		if newContact.Phone != nil {
			var existingContact mdl.ContactModel
			if err := h.ctx.DB.Where("phone = ? and company_id = ?", newContact.Phone, *newContact.CompanyID).First(&existingContact).Error; err == nil {
				newContact.ID = existingContact.ID
			}
		}
		if newContact.Email != "" {
			var existingContact mdl.ContactModel
			if err := h.ctx.DB.Where("email = ? and company_id = ?", newContact.Email, *newContact.CompanyID).First(&existingContact).Error; err == nil {
				newContact.ID = existingContact.ID
			}
		}

		err := h.ctx.DB.Save(&newContact).Error
		if err != nil {
			log.Println(err)
			continue
		}

		contacts = append(contacts, newContact)
		fmt.Println("------------------------------------------------")
		fmt.Println("ADD", newContact.Name)
	}

	fmt.Println("TOTAL ADDED", len(contacts))

	return contacts, nil
}

func cleanString(s string) string {
	return strings.TrimSpace(s)
}

func randomColor() string {
	colors := []string{
		"#FF5733", "#33FF57", "#3357FF", "#FF33A1", "#A133FF",
		"#33FFF5", "#F5FF33", "#FF8C33", "#FF3380", "#33FFBD",
		"#FFC400", "#00BFFF", "#FF00FF", "#008000", "#4B0082",
		"#9400D3", "#00FF00", "#FFD700", "#C71585", "#00FFFF",
		"#FF1493", "#00BFFF", "#32CD32", "#6495ED", "#DC143C",
		"#006400", "#8B008B", "#B03060", "#FF6347", "#1E90FF",
		"#228B22", "#FF7F24", "#FFC0CB", "#8B0A1A", "#4682B4",
		"#228B22", "#00688B", "#C71585", "#00BFFF", "#9400D3",
		"#FFD700", "#00FF00", "#FF1493", "#4B0082", "#9400D3",
		"#FFC400", "#00BFFF", "#FF00FF", "#008000", "#4B0082",
		"#9400D3", "#00FF00", "#FFD700", "#C71585", "#00FFFF",
		"#FF1493", "#00BFFF", "#32CD32", "#6495ED", "#DC143C",
		"#006400", "#8B008B", "#B03060", "#FF6347", "#1E90FF",
		"#228B22", "#FF7F24", "#FFC0CB", "#8B0A1A", "#4682B4",
	}
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(colors))))
	if err != nil {
		log.Println(err)
		return "#000000" // default to black if there's an error
	}
	return colors[n.Int64()]
}
