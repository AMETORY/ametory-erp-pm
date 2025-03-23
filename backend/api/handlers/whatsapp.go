package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"ametory-pm/services"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WhatsappHandler struct {
	erpContext                  *context.ERPContext
	waService                   *whatsmeow_client.WhatsmeowService
	appService                  *app.AppService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	geminiService               *google.GeminiService
}

// var eligibleKeyWords = []string{"Order", "order", "ORDER", "Orders", "orders", "ORDERS", "LOGIN", "login", "Login", "Menu", "MENU", "menu", "logout"}

func NewWhatsappHandler(erpContext *context.ERPContext) *WhatsappHandler {
	var waService *whatsmeow_client.WhatsmeowService
	waSrv, ok := erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
	if ok {
		waService = waSrv
	}

	var appService *app.AppService
	appSrv, ok := erpContext.AppService.(*app.AppService)
	if ok {
		appService = appSrv
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := erpContext.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	geminiService, ok := erpContext.ThirdPartyServices["GEMINI"].(*google.GeminiService)
	if !ok {
		panic("GeminiService is not found")
	}

	return &WhatsappHandler{
		erpContext:                  erpContext,
		waService:                   waService,
		appService:                  appService,
		customerRelationshipService: customerRelationshipService,
		geminiService:               geminiService,
	}
}

func (h *WhatsappHandler) SendMessage(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	var input struct {
		JID      string `json:"jid"`
		Session  string `json:"session"`
		Message  string `json:"message"`
		Receiver string `json:"receiver"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg, err := h.customerRelationshipService.WhatsappService.GetWhatsappLastMessages(input.JID, input.Session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	splitJID := strings.Split(input.JID, "@")
	splitSep := strings.Split(splitJID[0], ":")

	info := make(map[string]interface{})
	json.Unmarshal([]byte(msg.Info), &info)

	info["Timestamp"] = time.Now().Format("2006-01-02T15:04:05-07:00")

	b, _ := json.Marshal(info)
	msg.Info = string(b)

	var waDataReply models.WhatsappMessageModel = models.WhatsappMessageModel{
		Sender:   splitSep[0],
		Receiver: input.Receiver,
		Message:  input.Message,
		// MediaURL: mediaURLSaved,
		// MimeType: msg.MimeType,
		MessageInfo: info,
		Info:        msg.Info,
		Session:     input.Session,
		JID:         input.JID,
		IsFromMe:    true,
		IsGroup:     msg.MessageInfo["IsGroup"].(bool),
	}
	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	to := waDataReply.Receiver
	if waDataReply.IsGroup {
		to = waDataReply.Session
	}
	_, err = h.erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(whatsmeow_client.WaMessage{
		JID:     waDataReply.JID,
		Text:    waDataReply.Message,
		To:      to,
		IsGroup: waDataReply.IsGroup,
	})
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg, err = h.customerRelationshipService.WhatsappService.GetWhatsappLastMessages(input.JID, input.Session)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": msg})
}

func (h *WhatsappHandler) GetDevice(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	device, err := h.waService.GetDevices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(device, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
func (h *WhatsappHandler) DeleteDevice(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	err := h.waService.DeviceDelete(jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
func (h *WhatsappHandler) GetQR(c *gin.Context) {
	session := c.Params.ByName("session")
	if session == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	respBody, err := h.waService.GetQR(session)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": respBody})
}
func (h *WhatsappHandler) CreateQR(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	input := struct {
		Webhook   string `json:"webhook"`
		HeaderKey string `json:"header_key"`
		Session   string `json:"session"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	respBody, err := h.waService.CreateQR(input.Session, input.Webhook, input.HeaderKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var resp map[string]interface{}
	if err := json.Unmarshal(respBody, &resp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": resp})
}
func (h *WhatsappHandler) UpdateWebhook(c *gin.Context) {
	if h.waService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}
	input := struct {
		Webhook   string `json:"webhook"`
		HeaderKey string `json:"header_key"`
		Session   string `json:"session"`
	}{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.waService.UpdateWebhook(input.Session, input.Webhook, input.HeaderKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func (h *WhatsappHandler) GetMessagesHandler(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	messages, err := h.customerRelationshipService.WhatsappService.GetWhatsappMessages(*c.Request, c.DefaultQuery("search", ""), jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	items := messages.Items.(*[]models.WhatsappMessageModel)
	messages.Items = reverse(*items)

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": messages})
}
func (h *WhatsappHandler) GetConversationsHandler(c *gin.Context) {
	jid := c.Params.ByName("jid")
	if jid == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "session is required"})
		return
	}

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	conversations, err := h.customerRelationshipService.WhatsappService.GetMessageSession(jid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for i, v := range conversations {
		if v.IsGroup {
			groupInfo, err := h.waService.GetGroupInfo(jid, v.Session)

			if err == nil {
				groupName, ok := groupInfo["data"].(map[string]interface{})["Name"].(string)
				if ok {
					v.MessageInfo["GroupName"] = groupName
				}
			}
		}
		conversations[i] = v
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": conversations})
}

func reverse(messages []models.WhatsappMessageModel) []models.WhatsappMessageModel {
	for i, j := 0, len(messages)-1; i < len(messages)/2; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}

func (h *WhatsappHandler) WhatsappGetNumberHandler(c *gin.Context) {
	code, err := services.REDIS.Get(*h.erpContext.Ctx, "verify:"+c.Param("code")).Result()

	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid code"})
		return
	}
	data := map[string]interface{}{}
	err = json.Unmarshal([]byte(code), &data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Invalid code"})
		return
	}

	c.JSON(200, gin.H{"message": "Success", "data": data["phone_number"]})
}
func (h *WhatsappHandler) WhatsappRegisterHandler(c *gin.Context) {
	var input struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Phone   string `json:"phone"`
		Address string `json:"address"`
		Code    string `json:"code"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dataStr, err := services.REDIS.Get(*h.erpContext.Ctx, "verify:"+input.Code).Result()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{}

	err = json.Unmarshal([]byte(dataStr), &data)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	phone := data["phone_number"].(string)
	companyID, ok := data["company_id"].(string)
	if !ok {
		companyID = ""
	}

	if phone != input.Phone {
		c.JSON(400, gin.H{"error": "invalid code"})
		return
	}

	phoneNumber := utils.ParsePhoneNumber(input.Phone, "ID")

	var member models.ContactModel = models.ContactModel{
		Name:       input.Name,
		Email:      input.Email,
		Phone:      &phoneNumber,
		Address:    input.Address,
		CompanyID:  &companyID,
		IsCustomer: true,
	}
	member.ID = utils.Uuid()
	err = h.erpContext.DB.Create(&member).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	sendWAMessage(h.erpContext, data["jid"].(string), phoneNumber, "Registration has been completed")

	c.JSON(200, gin.H{"message": "Registration has been completed"})
}
func (h *WhatsappHandler) WhatsappWebhookHandler(c *gin.Context) {

	var body whatsapp.MsgObject

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// GET CONNECTION
	conn, err := h.appService.ConnectionService.GetConnectionBySession(body.SessionName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	convMsg := ""
	if body.Message.Conversation != nil {
		convMsg = *body.Message.Conversation
	}

	var sessionAuth *models.ContactModel
	if conn.SessionAuth {
		// CHECK IS PHONE NUMBER REGISTERED
		if !h.appService.ConnectionService.IsPhoneNumberRegistered(body.Sender) {
			randomStr := utils.RandString(8, false)

			data := map[string]interface{}{
				"phone_number": body.Sender,
				"jid":          body.JID,
				"company_id":   conn.CompanyID,
			}

			b, _ := json.Marshal(data)
			services.REDIS.Set(*h.erpContext.Ctx, fmt.Sprintf("verify:%s", randomStr), string(b), 30*time.Minute)
			replyData := whatsmeow_client.WaMessage{
				JID: body.JID,
				Text: fmt.Sprintf(`Selamat datang di %s
Anda belum terdaftar di sistem kami, silakan lakukan pendaftaran terlebih dahulu dengan mengikuti tautan berikut:

%s

*TERIMA KASIH*`, conn.Name, config.App.Server.FrontendURL+"/member/register/"+randomStr),
				To:      body.Sender,
				IsGroup: false,
			}

			h.waService.SendMessage(replyData)
			return
		}
		// CHECK SESSION AUTH

		session, err := h.appService.ConnectionService.GetActiveSession(body.Sender)
		if err != nil {
		}
		sessionAuth = session

		if convMsg == "LOGIN" && session == nil {
			err := doLogin(h.erpContext, body.JID, body.Sender, conn)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			return
		}

		if session == nil {
			sendWAMessage(h.erpContext, body.JID, body.Sender, "Anda belum Login, silakan ketik *LOGIN* lalu kirim untuk melakukan login")
			return
		}

	}

	fmt.Println("session", sessionAuth)

	if body.Sender == "status" {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
		return
	}

	if conn.GeminiAgent != nil {
		h.geminiService.SetupModel(conn.GeminiAgent.SetTemperature, conn.GeminiAgent.SetTopK, conn.GeminiAgent.SetTopP, conn.GeminiAgent.SetMaxOutputTokens, conn.GeminiAgent.ResponseMimetype, conn.GeminiAgent.Model)
		h.geminiService.SetUpSystemInstruction(fmt.Sprintf(`%s
		
%s`, conn.GeminiAgent.SystemInstruction, `

Tolong jawab dalam format : 
{
  "response": string,
  "type": string,
  "command": string,
  "params": object
}

Keterangan:
response: jawaban bila tipe nya pertanyaan
type: command atau question
command: jika tipe command
params: jika tipe command dibutuhkan parameter

`))

		var histories []models.GeminiHistoryModel
		err = h.erpContext.DB.Model(&models.GeminiHistoryModel{}).Find(&histories, "agent_id = ? and is_model = ?", conn.GeminiAgent.ID, true).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Agent histories is not found"})
			return
		}
		chatHistories := []map[string]any{}
		for _, v := range histories {
			chatHistories = append(chatHistories, map[string]any{
				"role":    "user",
				"content": v.Input,
			})
			chatHistories = append(chatHistories, map[string]any{
				"role":    "model",
				"content": v.Output,
			})
		}

		output, err := h.geminiService.GenerateContent(*h.erpContext.Ctx, convMsg, chatHistories, "", "")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var response geminiResponse
		err = json.Unmarshal([]byte(output), &response)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		sendWAMessage(h.erpContext, body.JID, body.Sender, response.Response)
	}

	var whatsappSession *models.WhatsappMessageSession
	err = h.erpContext.DB.First(&whatsappSession, "session = ?", body.SessionID).Error
	if err != nil {
		now := time.Now()
		if errors.Is(err, gorm.ErrRecordNotFound) {
			sessionData := models.WhatsappMessageSession{
				JID:          body.JID,
				Session:      body.SessionID,
				SessionName:  body.SessionName,
				LastOnlineAt: &now,
				LastMessage:  convMsg,
			}
			if sessionAuth != nil {
				sessionData.CompanyID = sessionAuth.CompanyID
			}
			h.erpContext.DB.Create(&sessionData)
		}
	}
	infoByte, err := json.Marshal(body.Info)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
		Sender:   body.Sender,
		Message:  convMsg,
		MimeType: body.MimeType,
		Info:     string(infoByte),
		Session:  body.SessionID,
		JID:      body.JID,
		IsFromMe: body.Info["IsFromMe"].(bool),
		IsGroup:  body.Info["IsGroup"].(bool),
	}

	if sessionAuth != nil {
		waData.ContactID = &sessionAuth.ID
		waData.CompanyID = sessionAuth.CompanyID
	}
	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
	if err != nil {
		// log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func doLogin(erpContext *context.ERPContext, jid, sender string, conn *connection.ConnectionModel) error {
	var member *models.ContactModel
	err := erpContext.DB.Find(&member, "phone = ?", sender).Error
	if err != nil {
		return err
	}
	b, err := json.Marshal(member)
	if err != nil {
		return err
	}
	services.REDIS.Set(*erpContext.Ctx, fmt.Sprintf("session:%s", *member.Phone), string(b), 7*24*time.Hour)
	msgContent := fmt.Sprintf(`
Hallo %s
Selamat datang di %s

Session anda akan berlaku selama 7 hari:

*ADMIN*
			`, member.Name, conn.Name)
	sendWAMessage(erpContext, jid, sender, msgContent)
	return nil
}

// func (h *WhatsappHandler) WebhookHandler(c *gin.Context) {
// 	setting, err := h.appService.GetSetting()
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	var session *mdl.Member

// 	// b, err := io.ReadAll(c.Request.Body)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	// 	return
// 	// }

// 	// log.Println(string(b))

// 	var msg objects.MsgObject
// 	if err := c.ShouldBindJSON(&msg); err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	session, _ = h.appService.MemberService.GetActiveSession(msg.Sender)
// 	if !h.appService.MemberService.IsPhoneNumberRegistered(msg.Sender) {
// 		randomStr := utils.RandString(8)

// 		data := map[string]interface{}{
// 			"phone_number": msg.Sender,
// 			"jid":          msg.JID,
// 		}

// 		b, _ := json.Marshal(data)
// 		services.REDIS.Set(fmt.Sprintf("verify:%s", randomStr), string(b), 30*time.Minute)
// 		replyData := whatsmeow_client.WaMessage{
// 			JID: msg.JID,
// 			Text: fmt.Sprintf(`Selamat datang di Sistem Informasi WA Dinkes Kota Sukabumi
// Anda belum terdaftar di sistem kami, silakan lakukan pendaftaran terlebih dahulu dengan mengikuti tautan berikut:

// %s

// *TERIMA KASIH*`, config.App.Server.FrontendURL+"/member/register/"+randomStr),
// 			To:      msg.Sender,
// 			IsGroup: false,
// 		}

// 		h.erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
// 		return
// 	}
// 	if msg.Message.Conversation != nil {
// 		conv := *msg.Message.Conversation
// 		splitConv := strings.Split(conv, " ")
// 		fmt.Println("splitConv", splitConv)
// 		if splitConv[0] != "" {
// 			for _, keyword := range eligibleKeyWords {
// 				if strings.Contains(strings.ToLower(*msg.Message.Conversation), strings.ToLower(keyword)) {
// 					infoByte, err := json.Marshal(msg.Info)
// 					if err != nil {
// 						log.Println(err)
// 						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 						return
// 					}
// 					var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
// 						Sender:   msg.Sender,
// 						Message:  conv,
// 						MimeType: msg.MimeType,
// 						Info:     string(infoByte),
// 						Session:  msg.SessionID,
// 						JID:      msg.JID,
// 						IsFromMe: msg.Info["IsFromMe"].(bool),
// 						IsGroup:  msg.Info["IsGroup"].(bool),
// 					}

// 					err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
// 					if err != nil {
// 						log.Println(err)
// 						c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 						return
// 					}
// 					waData.MessageInfo = msg.Info
// 					msgWs := gin.H{
// 						"type":    "NEW_MESSAGE",
// 						"session": msg.SessionID,
// 						"data":    waData,
// 					}

// 					msgWsByte, _ := json.Marshal(msgWs)
// 					services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 						return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 					})
// 					h.erpContext.TempData = msg
// 					cmd.TextCommand(h.erpContext, session, msg.JID, msg.Sender, splitConv...)
// 					return
// 				}
// 			}
// 		}

// 	}

// 	if session == nil {
// 		utils.SendWAMessage(h.erpContext, msg.JID, msg.Sender, "Anda belum Login, silakan ketik *LOGIN* lalu kirim untuk melakukan login")
// 		return
// 	}

// 	if msg.Sender == "status" {
// 		c.JSON(http.StatusOK, gin.H{"message": "ok"})
// 		return
// 	}

// 	b, _ := json.MarshalIndent(msg, "", "  ")
// 	log.Println("Received message:", string(b))
// 	var msgContent, fileUrl, mimeType string
// 	if msg.Message.Conversation != nil {
// 		msgContent = *msg.Message.Conversation
// 	}
// 	if msg.Message.ImageMessage != nil {
// 		msgContent = msg.Message.ImageMessage.Caption
// 		mimeType = msg.Message.ImageMessage.Mimetype
// 	}
// 	if msg.Message.VideoMessage != nil {
// 		msgContent = msg.Message.VideoMessage.Caption
// 		mimeType = msg.Message.VideoMessage.Mimetype
// 	}

// 	if msg.Message.DocumentMessage != nil {
// 		msgContent = msg.Message.DocumentMessage.Caption
// 		mimeType = msg.Message.DocumentMessage.Mimetype
// 	}
// 	var mediaURLSaved string
// 	if msg.MediaPath != "" {
// 		mediaURL := config.App.Whatsapp.BaseURL + msg.MediaPath

// 		resp, err := http.Get(mediaURL)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer resp.Body.Close()
// 		byteValue, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		path := filepath.Join("assets", msg.MediaPath)
// 		os.MkdirAll(filepath.Dir(path), os.ModePerm)
// 		if err := os.WriteFile(path, byteValue, 0644); err != nil {
// 			log.Println(err)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		mediaURLSaved = config.App.Server.BaseURL + "/" + path
// 		fileUrl = mediaURLSaved

// 	}

// 	infoByte, err := json.Marshal(msg.Info)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	var waData models.WhatsappMessageModel = models.WhatsappMessageModel{
// 		Sender:   msg.Sender,
// 		Message:  msgContent,
// 		MediaURL: mediaURLSaved,
// 		MimeType: msg.MimeType,
// 		Info:     string(infoByte),
// 		Session:  msg.SessionID,
// 		JID:      msg.JID,
// 		IsFromMe: msg.Info["IsFromMe"].(bool),
// 		IsGroup:  msg.Info["IsGroup"].(bool),
// 	}

// 	err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waData)
// 	if err != nil {
// 		log.Println(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	waData.MessageInfo = msg.Info
// 	msgWs := gin.H{
// 		"type":    "NEW_MESSAGE",
// 		"session": msg.SessionID,
// 		"data":    waData,
// 	}

// 	msgWsByte, _ := json.Marshal(msgWs)
// 	services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 		return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 	})
// 	splitJID := strings.Split(msg.JID, "@")
// 	splitSep := strings.Split(splitJID[0], ":")
// 	// utils.LogJson(msg.Message.ExtendedTextMessage)
// 	if msg.Message.ExtendedTextMessage != nil {

// 		contains := false
// 		for _, v := range msg.Message.ExtendedTextMessage.ContextInfo.MentionedJID {
// 			if v == splitSep[0]+"@"+splitJID[1] {
// 				contains = true
// 				break
// 			}
// 		}
// 		if contains {
// 			msgContent = msg.Message.ExtendedTextMessage.Text
// 		}

// 	}
// 	if waData.IsGroup && msg.Message.ExtendedTextMessage == nil {
// 		msgContent = ""
// 	}
// 	if msgContent != "" {

// 		if h.erpContext.ThirdPartyServices["GEMINI"] != nil && setting.AutoPilot {
// 			histories, err := services.REDIS.LRange(config.App.Redis.Key+"-"+msg.Sender, 0, 20).Result()
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			// fmt.Println("histories", histories)

// 			for i, j := 0, len(histories)-1; i < j; i, j = i+1, j-1 {
// 				histories[i], histories[j] = histories[j], histories[i]
// 			}

// 			userHistories := []map[string]interface{}{}
// 			for _, v := range histories {
// 				var history objects.History
// 				if err := json.Unmarshal([]byte(v), &history); err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}
// 				userHistories = append(userHistories, map[string]interface{}{
// 					"role":    "user",
// 					"content": history.Input,
// 				})
// 				userHistories = append(userHistories, map[string]interface{}{
// 					"role":    "model",
// 					"content": history.Output,
// 				})
// 			}

// 			// SEND TO AI
// 			log.Println("SEND MESSAGE TO AI", msgContent)
// 			resp, err := h.erpContext.ThirdPartyServices["GEMINI"].(*google.GeminiService).GenerateContent(*h.erpContext.Ctx, msgContent, userHistories, fileUrl, mimeType)
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}
// 			log.Println("Generated content:", resp)

// 			parseData := map[string]interface{}{}
// 			if err := json.Unmarshal([]byte(resp), &parseData); err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

// 			responseMsg := ""
// 			if parseData["type"] == "answer" {
// 				responseMsg = parseData["conversation"].(string)
// 			}
// 			if parseData["type"] == "command" {
// 				log.Println("Received command:", resp)
// 				msgContent, ok := parseData["message"].(string)
// 				if ok {
// 					responseMsg = msgContent
// 				}
// 				respCommand, err := command.GeminiResponse(parseData)
// 				if err == nil {
// 					responseMsg = respCommand
// 				}
// 			}
// 			// SEND MESSAGE
// 			if h.erpContext.ThirdPartyServices["WA"] != nil {
// 				to := msg.Sender
// 				if waData.IsGroup {
// 					to = msg.SessionID
// 				}
// 				replyData := whatsmeow_client.WaMessage{
// 					JID:     msg.JID,
// 					Text:    responseMsg,
// 					To:      to,
// 					IsGroup: waData.IsGroup,
// 				}

// 				utils.LogJson(replyData)
// 				_, err := h.erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
// 				if err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}

// 				info := make(map[string]interface{})
// 				json.Unmarshal(infoByte, &info)

// 				info["Timestamp"] = time.Now().Format("2006-01-02T15:04:05-07:00")

// 				i, _ := json.Marshal(info)
// 				infoByte = i

// 				var waDataReply models.WhatsappMessageModel = models.WhatsappMessageModel{
// 					Sender:   splitSep[0],
// 					Receiver: msg.Sender,
// 					Message:  responseMsg,
// 					// MediaURL: mediaURLSaved,
// 					// MimeType: msg.MimeType,
// 					MessageInfo: info,
// 					Info:        string(infoByte),
// 					Session:     msg.SessionID,
// 					JID:         msg.JID,
// 					IsFromMe:    true,
// 					IsGroup:     msg.Info["IsGroup"].(bool),
// 				}
// 				err = h.customerRelationshipService.WhatsappService.CreateWhatsappMessage(&waDataReply)
// 				if err != nil {
// 					log.Println(err)
// 					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 					return
// 				}

// 				msgWs := gin.H{
// 					"type":    "NEW_MESSAGE",
// 					"session": msg.SessionID,
// 					"data":    waDataReply,
// 				}

// 				msgWsByte, _ := json.Marshal(msgWs)
// 				services.WS.BroadcastFilter(msgWsByte, func(q *melody.Session) bool {
// 					return q.Request.URL.Path == "/api/v1/whatsapp/ws"
// 				})
// 			}

// 			// ADD HISTORY DATA
// 			historyData := objects.History{
// 				Sender: msg.Sender,
// 				Input:  msgContent,
// 				Output: resp,
// 			}

// 			b, err := json.Marshal(historyData)
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}
// 			err = services.REDIS.LPush(config.App.Redis.Key+"-"+msg.Sender, string(b)).Err()
// 			if err != nil {
// 				log.Println(err)
// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 				return
// 			}

//			}
//		}
//		c.JSON(http.StatusOK, gin.H{"message": "ok", "data": msg})
//	}

func sendWAMessage(erpContext *context.ERPContext, jid, to, message string) {
	replyData := whatsmeow_client.WaMessage{
		JID:     jid,
		Text:    message,
		To:      to,
		IsGroup: false,
	}
	erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(replyData)
}

type geminiResponse struct {
	Response string         `json:"response"`
	Type     string         `json:"type"`
	Command  string         `json:"command"`
	Params   map[string]any `json:"params"`
}
