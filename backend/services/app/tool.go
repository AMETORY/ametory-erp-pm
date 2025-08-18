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

func (a *AppService) SendTemplateMessageWhatsappAPI(
	customerRelationshipService *customer_relationship.CustomerRelationshipService,
	metaService *meta.MetaService,
	conn *connection.ConnectionModel,
	waDataReply models.WhatsappMessageModel,
	session *models.WhatsappMessageSession,
	member *models.MemberModel,
	files []models.FileModel,
	products []models.ProductModel,
) error {
	var quoteMessageID *string
	metaService.WhatsappApiService.SetAccessToken(&conn.AccessToken)

	// SEND PRIMARY MESSAGE
	resp, err := metaService.WhatsappApiService.SendMessage(conn.Session, waDataReply.Message, []*models.FileModel{}, session.Contact, quoteMessageID)
	if err != nil {
		return err
	}

	a.SaveWhatsappAPIResponse(customerRelationshipService, metaService, conn, waDataReply, session, member, resp, []models.FileModel{})

	// SEND TEMPLATE MESSAGES

	for _, v := range files {
		if v.Caption != nil {
			waDataReply.Message = *v.Caption
		} else {
			waDataReply.Message = ""
		}
		resp, err := metaService.WhatsappApiService.SendMessage(conn.Session, waDataReply.Message, []*models.FileModel{&v}, session.Contact, quoteMessageID)
		if err != nil {
			return err
		}
		a.SaveWhatsappAPIResponse(customerRelationshipService, metaService, conn, waDataReply, session, member, resp, []models.FileModel{v})
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
