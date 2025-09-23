package app

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"encoding/json"
	"fmt"
	"time"

	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/shared/objects"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/meta"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func (a *AppService) SendMessageWithTemplate(contact *models.ContactModel, metaService *meta.MetaService, conn *connection.ConnectionModel, template *models.MessageTemplate) error {
	metaService.WhatsappApiService.SetAccessToken(&conn.AccessToken)

	resp, err := metaService.WhatsappApiService.GetMessageTemplateByName(*template.BusinessID, *template.WhatsappTemplateID)
	if err != nil {
		return err
	}

	waResp, err := metaService.WhatsappApiService.SendTemplateMessage(conn.Session, contact, &resp.Data[0], template)
	if err != nil {
		return err
	}
	fmt.Println(waResp)

	return nil
}

func (a *AppService) SendTemplateMessageWhatsappAPI(
	customerRelationshipService *customer_relationship.CustomerRelationshipService,
	metaService *meta.MetaService,
	conn *connection.ConnectionModel,
	waDataReply models.WhatsappMessageModel,
	session *models.WhatsappMessageSession,
	member *models.MemberModel,
	files []models.FileModel,
	products []models.ProductModel,
	interactive *models.WhatsappInteractiveMessage,
) error {
	var quoteMessageID *string
	metaService.WhatsappApiService.SetAccessToken(&conn.AccessToken)

	thumbnail, restFiles := models.GetThumbnail(files)

	thumbnailFile := []*models.FileModel{}
	thumbnailFile2 := []models.FileModel{}
	if thumbnail != nil {
		thumbnailFile = append(thumbnailFile, thumbnail)
		thumbnailFile2 = append(thumbnailFile2, *thumbnail)
	}

	fmt.Println("thumbnailFile", thumbnailFile)
	// SEND PRIMARY MESSAGE
	resp, err := metaService.WhatsappApiService.SendMessage(conn.Session, waDataReply.Message, thumbnailFile, session.Contact, quoteMessageID, interactive)
	if err != nil {
		return err
	}

	a.SaveWhatsappAPIResponse(customerRelationshipService, metaService, conn, waDataReply, session, member, resp, thumbnailFile2, interactive)

	// SEND TEMPLATE MESSAGES
	// fmt.Println("FILES")
	// utils.LogJson(files)

	if interactive == nil {
		for _, v := range restFiles {
			if v.Caption != nil {
				waDataReply.Message = *v.Caption
			} else {
				waDataReply.Message = ""
			}
			resp, err := metaService.WhatsappApiService.SendMessage(conn.Session, waDataReply.Message, []*models.FileModel{&v}, session.Contact, quoteMessageID, nil)
			if err != nil {
				return err
			}
			a.SaveWhatsappAPIResponse(customerRelationshipService, metaService, conn, waDataReply, session, member, resp, []models.FileModel{v}, nil)
		}
	}

	return nil
}

func (a *AppService) SaveWhatsappAPIResponse(
	customerRelationshipService *customer_relationship.CustomerRelationshipService,
	metaService *meta.MetaService,
	conn *connection.ConnectionModel,
	waDataReply models.WhatsappMessageModel,
	session *models.WhatsappMessageSession,
	member *models.MemberModel,
	resp *objects.WaResponse,
	files []models.FileModel,
	interactive *models.WhatsappInteractiveMessage,
) error {
	if session.JID == "" {
		return nil
	}

	info := map[string]interface{}{
		"Timestamp": time.Now().Format(time.RFC3339),
	}
	infoByte, err := json.Marshal(info)
	if err != nil {
		return err
	}

	replyResponse := &models.WhatsappMessageModel{
		Receiver:  *session.Contact.Phone,
		Message:   waDataReply.Message,
		Session:   session.Session,
		JID:       session.JID,
		IsFromMe:  true,
		Info:      string(infoByte),
		IsGroup:   false,
		ContactID: &session.Contact.ID,
		CompanyID: conn.CompanyID,
		MemberID:  &member.ID,
	}

	if interactive != nil {
		b, _ := json.Marshal(interactive)
		replyResponse.InteractiveMessage = json.RawMessage(b)
		waDataReply.InteractiveMessage = json.RawMessage(b)
	}

	for _, v := range resp.Messages {
		replyResponse.MessageID = &v.ID
		waDataReply.MessageID = &v.ID
	}

	if len(files) > 0 {
		for _, file := range files {
			replyResponse.MediaURL = file.URL
			replyResponse.MimeType = file.MimeType
		}
		waDataReply.MediaURL = replyResponse.MediaURL
		waDataReply.MimeType = replyResponse.MimeType
	}

	err = customerRelationshipService.WhatsappService.CreateWhatsappMessage(replyResponse)
	if err != nil {
		return err
	}
	msgNotif := gin.H{
		"message":    waDataReply.Message,
		"command":    "WHATSAPP_RECEIVED",
		"session_id": session.ID,
		"data":       waDataReply,
	}
	msgNotifStr, _ := json.Marshal(msgNotif)
	a.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", config.App.Server.BaseURL, *session.CompanyID)
		return fmt.Sprintf("%s%s", config.App.Server.BaseURL, q.Request.URL.Path) == url
	})
	return nil
}
