package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"ametory-pm/services/app"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type WhatsappApiHandler struct {
	erpContext     *context.ERPContext
	contactService *contact.ContactService
	appService     *app.AppService
}

func NewWhatsappApiHandler(erpContext *context.ERPContext) *WhatsappApiHandler {
	contactService, ok := erpContext.ContactService.(*contact.ContactService)
	if !ok {
		panic("ContactService is not instance of contact.ContactService")
	}
	var appService *app.AppService
	appSrv, ok := erpContext.AppService.(*app.AppService)
	if ok {
		appService = appSrv
	}
	return &WhatsappApiHandler{
		erpContext:     erpContext,
		contactService: contactService,
		appService:     appService,
	}
}

func (h *WhatsappApiHandler) WhatsappApiWebhookHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		VerifyFacebookWebhook(c)
		return
	}
	// var bodyMap map[string]interface{}
	// if err := c.BindJSON(&bodyMap); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body"})
	// 	return
	// }

	// fmt.Println("BODY WEBHOOK")
	// utils.LogJson(bodyMap)

	var data whatsapp.WhatsappApiWebhookRequest
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid request body #2"})
		return
	}
	// fmt.Println("DATA WEBHOOK")
	// utils.LogJson(data)
	now := time.Now()
	// fmt.Println("DATA WEBHOOK")
	// utils.LogJson(data)

	var conn connection.ConnectionModel
	for _, entry := range data.Entry {
		for _, change := range entry.Changes {
			if change.Field == "messages" && change.Value.MessagingProduct == "whatsapp" {
				if len(change.Value.Contacts) > 0 {
					phoneNumberID := change.Value.Metadata.PhoneNumberID
					fmt.Println("PHONE NUMBER ID", phoneNumberID)
					err := h.erpContext.DB.First(&conn, "session = ?", phoneNumberID).Error
					if err != nil {
						fmt.Println("ERROR GET CONNECTION BY PHONE NUMBER ID", err)
						continue
					}

					var session mdl.WhatsappMessageSession
					err = h.erpContext.DB.Where("j_id = ?", phoneNumberID).First(&session).Error
					if err != nil {
						fmt.Println("ERROR GET CONNECTION BY PHONE NUMBER ID", err)
						continue
					}

					// GET CONNECTION BY PHONE NUMBER ID
					contact, err := h.contactService.GetContactByPhone(change.Value.Contacts[0].WAID, *conn.CompanyID)
					if err != nil {
						if errors.Is(err, gorm.ErrRecordNotFound) {
							contact = &mdl.ContactModel{
								CompanyID: conn.CompanyID,
								Name:      change.Value.Contacts[0].Profile.Name,
								Phone:     &change.Value.Contacts[0].WAID,
							}
							if err := h.erpContext.DB.Create(contact).Error; err != nil {
								fmt.Println("ERROR CREATE CONTACT", err)
								continue
							}
						} else {
							fmt.Println("ERROR GET CONTACT BY PHONE NUMBER ID", err)
							continue
						}
					}

					// GET CONTACT BY PHONE NUMBER ID
					fmt.Println("GET CONTACT")
					utils.SaveJson(contact)

					// GET MESSAGE
					for _, msg := range change.Value.Messages {
						message := ""
						// QUOTE MESSAGE

						if msg.Type == "text" {
							message = msg.Text.Body
						}
						if msg.Type == "image" && msg.Image != nil {
							message = msg.Image.Caption
							path, err := GetFacebookMediaObject(msg.Image.ID, conn.AccessToken)
							if err != nil {
								fmt.Println("ERROR", err)
							}
							fmt.Println("PATH", *path)
						}
						sessionWa := fmt.Sprintf("%s@%s", *contact.Phone, conn.Session)
						waMsg := mdl.WhatsappMessageModel{
							Message:   message,
							MessageID: &msg.ID,
							Sender:    msg.From,
							JID:       phoneNumberID,
							Contact:   contact,
							SentAt:    &now,
							Session:   sessionWa,
							CompanyID: conn.CompanyID,
						}

						if msg.Context != nil {
							waMsg.QuotedMessageID = &msg.Context.ID
						}

						// utils.LogJson(waMsg)

						if err := h.erpContext.DB.Create(&waMsg).Error; err != nil {
							fmt.Println("ERROR CREATE WHATSAPP MESSAGE #2", err)
							continue
						}

						msg := gin.H{
							"message":    message,
							"command":    "WHATSAPP_RECEIVED",
							"session_id": session.ID,
							"data":       waMsg,
						}
						b, _ := json.Marshal(msg)
						h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
							url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *conn.CompanyID)
							return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
						})

					}

				}
			}

		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func GetFacebookMediaObject(mediaID, accessToken string) (*string, error) {
	url := fmt.Sprintf("https://graph.facebook.com/v21.0/%s", mediaID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media object, status code: %d", resp.StatusCode)
	}

	var mediaObject whatsapp.FacebookMedia
	if err := json.NewDecoder(resp.Body).Decode(&mediaObject); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	req, err = http.NewRequest("GET", mediaObject.URL, nil)
	if err != nil {
		return nil, err
	}

	fmt.Println("URL", mediaObject.URL)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media content, status code: %d", resp.StatusCode)
	}

	extension := ""
	switch mediaObject.MimeType {
	case "image/jpeg":
		extension = ".jpg"
	case "image/png":
		extension = ".png"
	default:
		return nil, fmt.Errorf("unsupported media type: %s", mediaObject.MimeType)
	}

	fileName := fmt.Sprintf("%s%s", mediaObject.ID, extension)
	return downloadAndSaveMedia(mediaObject.URL, fileName)
}

func downloadAndSaveMedia(url, fileName string) (*string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get media, status code: %d", resp.StatusCode)
	}

	assetsFolder := "./assets/images/"
	if _, err := os.Stat(assetsFolder); os.IsNotExist(err) {
		os.Mkdir(assetsFolder, os.ModePerm)
	}

	destination := filepath.Join(assetsFolder, fileName)
	if err := saveResponseBodyToFile(resp.Body, destination); err != nil {
		return nil, err
	}

	if config.App.Server.StorageProvider == "google" {
		// destination, err = HandleGoogleStorage(destination)
		// if err != nil {
		// 	return nil, err
		// }
	}

	return &destination, nil
}

func saveResponseBodyToFile(respBody io.Reader, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, respBody)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	return nil
}
