package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/contact"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/project_management"
	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"gorm.io/gorm"
)

type FacebookHandler struct {
	ctx                         *context.ERPContext
	appService                  *app.AppService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	geminiService               *google.GeminiService
	pmService                   *project_management.ProjectManagementService
	contactSrv                  *contact.ContactService
}

func NewFacebookHandler(ctx *context.ERPContext) *FacebookHandler {
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	geminiService, ok := ctx.ThirdPartyServices["GEMINI"].(*google.GeminiService)
	if !ok {
		panic("GeminiService is not found")
	}

	pmService, ok := ctx.ProjectManagementService.(*project_management.ProjectManagementService)
	if !ok {
		panic("ProjectManagementService is not instance of project_management.ProjectManagementService")
	}
	contactSrv, ok := ctx.ContactService.(*contact.ContactService)
	if !ok {
		panic("ContactService is not instance of contact.ContactService")
	}
	return &FacebookHandler{ctx: ctx,
		appService:                  appService,
		customerRelationshipService: customerRelationshipService,
		geminiService:               geminiService,
		pmService:                   pmService,
		contactSrv:                  contactSrv}
}

func (h *FacebookHandler) GetInstagramSessionsHandler(c *gin.Context) {
	sessions, err := h.customerRelationshipService.InstagramService.GetSessionMessageBySessionName("", *c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": sessions})
}

func (h *FacebookHandler) GetSessionMessagesHandler(c *gin.Context) {
	sessionId := c.Query("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session *models.InstagramMessageSession
	err := h.ctx.DB.First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	messages, err := h.customerRelationshipService.InstagramService.GetMessageSessionChatBySessionName(session.ID, session.ContactID, *c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	messages.Items = reverseInstagram(*messages.Items.(*[]models.InstagramMessage))

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": messages})
}
func (h *FacebookHandler) GetInstagramSessionDetailHandler(c *gin.Context) {
	sessionId := c.Params.ByName("session_id") // c.Params.ByName("sessionId")

	if h.customerRelationshipService == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
		return
	}

	var session models.InstagramMessageSession
	err := h.ctx.DB.Preload("Contact").First(&session, "id = ?", sessionId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	var connection connection.ConnectionModel
	err = h.ctx.DB.First(&connection, "id = ?", session.Session).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ok", "data": session, "connection": connection})
}

func (h *FacebookHandler) SendInstagramMessageHandler(c *gin.Context) {
	sessionID := c.Params.ByName("sessionId")
	var input struct {
		Message string `json:"message" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var session models.InstagramMessageSession
	if err := h.ctx.DB.Preload("Contact").First(&session, "id = ?", sessionID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	var connection connection.ConnectionModel
	if err := h.ctx.DB.First(&connection, "id = ?", session.Session).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if err := h.customerRelationshipService.InstagramService.SendInstagramMessage(connection.SessionName, *session.Contact.InstagramID, input.Message, connection.AccessToken); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	userID := c.MustGet("userID").(string)
	memberID := c.MustGet("memberID").(string)

	instagramMsgData := models.InstagramMessage{
		BaseModel:                 shared.BaseModel{ID: utils.Uuid()},
		ContactID:                 session.ContactID,
		Message:                   input.Message,
		Session:                   connection.ID,
		CompanyID:                 connection.CompanyID,
		InstagramMessageSessionID: &session.ID,
		IsFromMe:                  true,
		UserID:                    &userID,
		MemberID:                  &memberID,
	}
	err := h.ctx.DB.Create(&instagramMsgData).Error
	if err != nil {
		utils.LogPrintf(" error create message: %s\n", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msgNotif := gin.H{
		"message":    input.Message,
		"command":    "INSTAGRAM_RECEIVED",
		"session_id": session.ID,
		"data":       instagramMsgData,
	}
	msgNotifStr, _ := json.Marshal(msgNotif)
	h.appService.Websocket.BroadcastFilter(msgNotifStr, func(q *melody.Session) bool {
		url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
		return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
	})

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
func (h *FacebookHandler) FacebookWebhookHandler(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		VerifyFacebookWebhook(c)
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("[%s] error read req body: %s\n", time.Now().Format(time.RFC3339), err.Error())
		return
	}
	log.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), string(body))

	var req objects.FacebookWebhookResponse
	if err := json.Unmarshal(body, &req); err != nil {
		log.Printf("[%s] error unmarshal req body: %s\n", time.Now().Format(time.RFC3339), err.Error())
		return
	}

	utils.SaveJson(req)
	if req.Object == "instagram" {
		if len(req.Entry) > 0 {
			if len(req.Entry[0].Messaging) > 0 {

				now := time.Now()
				var senderID = req.Entry[0].Messaging[0].Sender.ID
				var recipientID = req.Entry[0].Messaging[0].Recipient.ID
				var instagramMsg = req.Entry[0].Messaging[0].Message.Text
				var attachments = req.Entry[0].Messaging[0].Message.Attachments

				// GET CONNECTION FROM RECIPIENT ID

				var connection connection.ConnectionModel
				if err := h.ctx.DB.Model(&connection).Where("session_name = ?", recipientID).First(&connection).Error; err != nil {
					utils.LogPrintf(" error get connection: %s\n", err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				// GET PROFILE OF SENDER
				senderData := getContact(senderID, connection.AccessToken)
				if senderData == nil {
					utils.LogPrintf("error get sender profile: %s\n", errors.New("error get sender profile"))
					c.JSON(http.StatusInternalServerError, gin.H{"error": "error get sender profile"})
					return
				}

				profilePic := ""
				profilePicData := getProfilePicture(senderID, connection.AccessToken)
				if profilePicData != nil {
					profilePic = profilePicData.ProfilePictureURL
				}

				utils.LogPrintf(" sender: %s, recipientID: %s, msg: %s\n", senderData.Name, recipientID, instagramMsg)

				// GET CONTACT By senderID
				var contact models.ContactModel
				err := h.ctx.DB.Model(&contact).Where("instagram_id = ?", senderID).First(&contact).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					connType := "instagram"
					contact.ID = utils.Uuid()
					contact.Name = senderData.Name
					contact.InstagramID = &senderID
					contact.ConnectionType = &connType
					contact.CompanyID = connection.CompanyID
					err := h.ctx.DB.Create(&contact).Error
					if err != nil {
						utils.LogPrintf(" error create contact: %s\n", err.Error())
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
				} else if err != nil {
					log.Printf("[%s] error get contact: %s\n", time.Now().Format(time.RFC1123), err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				if profilePic != "" {
					path, mimeType, err := saveFileContenFromUrl(profilePic)
					if err == nil {
						mediaURLSaved := config.App.Server.FrontendURL + "/" + path
						var file models.FileModel
						file.ID = utils.Uuid()
						file.FileName = profilePic
						file.Path = path
						file.URL = mediaURLSaved
						file.RefType = "contact"
						file.RefID = contact.ID
						file.MimeType = mimeType
						err = h.ctx.DB.Create(&file).Error
						if err != nil {
							utils.LogPrintf(" error create file: %s\n", err.Error())
							c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
							return
						}

					}

				}

				var session models.InstagramMessageSession
				err = h.ctx.DB.First(&session, "contact_id = ?", contact.ID).Error
				if errors.Is(err, gorm.ErrRecordNotFound) {
					session.ID = utils.Uuid()
					session.ContactID = &contact.ID
					session.Session = connection.ID
					session.SessionName = contact.Name
					session.LastMessage = instagramMsg
					session.LastOnlineAt = &now
					session.CompanyID = connection.CompanyID
					err = h.ctx.DB.Create(&session).Error
					if err != nil {
						utils.LogPrint(err.Error())
						c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
				}
				if len(attachments) > 0 {
					instagramMsg = ""
				}
				instagramMsgData := models.InstagramMessage{
					BaseModel:                 shared.BaseModel{ID: utils.Uuid()},
					ContactID:                 &contact.ID,
					Message:                   instagramMsg,
					Session:                   connection.ID,
					CompanyID:                 connection.CompanyID,
					MessageID:                 &req.Entry[0].ID,
					InstagramMessageSessionID: &session.ID,
				}
				err = h.ctx.DB.Create(&instagramMsgData).Error
				if err != nil {
					utils.LogPrintf(" error create message: %s\n", err.Error())
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				if len(attachments) > 0 {
					for _, v := range attachments {
						// utils.LogPrintf(" sender: %s, recipientID: %s, msg: %s\n", senderData.Name, recipientID, v.Payload.URL)
						path, mimeType, _ := saveFileContenFromUrl(v.Payload.URL)
						if path != "" {
							mediaURLSaved := config.App.Server.FrontendURL + "/" + path
							var file models.FileModel
							file.ID = utils.Uuid()
							file.FileName = filepath.Base(path)
							file.Path = path
							file.URL = mediaURLSaved
							file.RefType = "instagram_message"
							file.RefID = instagramMsgData.ID
							file.MimeType = mimeType
							err = h.ctx.DB.Create(&file).Error

							if err != nil {
								utils.LogPrintf(" error create file: %s\n", err.Error())
								c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
								return
							}
						}

					}
				}

				msg := gin.H{
					"message":    instagramMsg,
					"command":    "INSTAGRAM_RECEIVED",
					"session_id": session.ID,
					"data":       instagramMsgData,
				}

				fmt.Println(msg)
				b, _ := json.Marshal(msg)
				h.appService.Websocket.BroadcastFilter(b, func(q *melody.Session) bool {
					url := fmt.Sprintf("%s/api/v1/ws/%s", h.appService.Config.Server.BaseURL, *session.CompanyID)
					fmt.Println("URL", url)
					return fmt.Sprintf("%s%s", h.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
				})

			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

func VerifyFacebookWebhook(c *gin.Context) {
	verifyToken := c.Query("hub.verify_token")
	challenge := c.Query("hub.challenge")

	if verifyToken == config.App.Facebook.FacebookVerifyToken {
		c.String(http.StatusOK, challenge)
	} else {
		c.String(http.StatusUnauthorized, "Invalid verify token")
	}
}
func (h *FacebookHandler) InstagramCallbackHandler(c *gin.Context) {
	connectionID := c.Query("connection_id")
	code := c.Query("code")

	state := c.Query("state")
	if state != "" {
		// &state=connection_id-f14cba4f-12e8-4901-874c-3c5a5b8df04f#_
		connID := strings.ReplaceAll(state, "connection_id-", "")

		stateParts := strings.Split(connID, "#_")
		if len(stateParts) > 0 {
			state = stateParts[0]
		}
		connectionID = state
	}

	if code != "" && connectionID != "" {
		redirectUri := h.appService.Config.Facebook.IGRedirectURL
		formData := url.Values{
			"client_id":     {config.App.Facebook.AppIGID},
			"client_secret": {config.App.Facebook.AppIGSecret},
			"grant_type":    {"authorization_code"},
			"redirect_uri":  {redirectUri},
			"code":          {code},
		}
		url := "https://api.instagram.com/oauth/access_token"
		resp, err := http.PostForm(url, formData)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access token"})
			return
		}

		body, _ := io.ReadAll(resp.Body)
		fmt.Println(string(body))

		var response struct {
			AccessToken string   `json:"access_token"`
			UserID      int      `json:"user_id"`
			Permissions []string `json:"permissions"`
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println("ERROR UNMARSHAL", err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		utils.SaveJson(response)
		var instagramToken = ""
		instagramToken = response.AccessToken
		token, _ := exchangeInstagramToken(response.AccessToken)
		if token != "" {
			instagramToken = token
		}
		var conn connection.ConnectionModel

		err = h.ctx.DB.First(&conn, "id = ?", connectionID).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		conn.AccessToken = instagramToken
		conn.Status = "ACTIVE"
		sessionName := fetchProfile(instagramToken)
		conn.SessionName = sessionName
		err = h.ctx.DB.Save(&conn).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/connection", h.appService.Config.Server.FrontendURL))

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hallo",
	})
}
func (h *FacebookHandler) FacebookCallbackHandler(c *gin.Context) {
	connectionID := c.Query("connection_id")
	code := c.Query("code")

	state := c.Query("state")
	if state != "" {
		// &state=connection_id-f14cba4f-12e8-4901-874c-3c5a5b8df04f#_
		connID := strings.ReplaceAll(state, "connection_id-", "")

		stateParts := strings.Split(connID, "#_")
		if len(stateParts) > 0 {
			state = stateParts[0]
		}
		connectionID = state
	}

	if code != "" && connectionID != "" {
		redirectUrl := fmt.Sprintf(`%s/%s`, h.appService.Config.Facebook.RedirectURL, connectionID)
		url := fmt.Sprintf(`https://graph.facebook.com/v18.0/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s`, h.appService.Config.Facebook.AppID,
			redirectUrl,
			h.appService.Config.Facebook.AppSecret,
			code)
		log.Println("url", url)
		log.Println("connectionID", connectionID)
		resp, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get access token",
			})
			return
		}

		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to read response body",
			})
			return
		}

		log.Println("resp", string(bodyBytes))
		var result objects.FacebookAccessTokenResponse
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to parse response body",
			})
			return
		}

		log.Println("connectionID", connectionID)

		log.Println("Parsed response:")
		log.Println("Access token:", result.AccessToken)
		log.Println("Token type:", result.TokenType)
		log.Println("Expires in:", result.ExpiresIn)

		var conn connection.ConnectionModel

		err = h.ctx.DB.First(&conn, "id = ?", connectionID).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		conn.AccessToken = result.AccessToken
		conn.Status = "ACTIVE"
		err = h.ctx.DB.Save(&conn).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, fmt.Sprintf("%s/connection", h.appService.Config.Server.FrontendURL))

	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Hallo",
	})

}

func exchangeInstagramToken(shortLivedToken string) (string, error) {
	clientSecret := config.App.Facebook.AppIGSecret
	url := fmt.Sprintf("https://graph.instagram.com/access_token?grant_type=ig_exchange_token&client_secret=%s&access_token=%s", clientSecret, shortLivedToken)
	fmt.Println("URL", url)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to exchange token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR", fmt.Errorf("failed to exchange token, status code: %d", resp.StatusCode))
		return "", fmt.Errorf("failed to exchange token, status code: %d", resp.StatusCode)
	}

	var response struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Println("ERROR", err.Error())
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	return response.AccessToken, nil
}

func fetchProfile(instagramToken string) string {
	profileURL := fmt.Sprintf("https://graph.instagram.com/v21.0/me?fields=user_id,name,profile_picture_url&access_token=%s", instagramToken)
	resp, err := http.Get(profileURL)
	if err != nil {
		fmt.Println("ERROR", err.Error())

		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("ERROR", fmt.Errorf("failed to fetch instagram profile, status code: %d", resp.StatusCode))

		return ""
	}

	var profileResponse struct {
		UserID            string `json:"user_id"`
		Name              string `json:"name"`
		ProfilePictureURL string `json:"profile_picture_url"`
	}
	err = json.NewDecoder(resp.Body).Decode(&profileResponse)
	if err != nil {
		fmt.Println("ERROR", err.Error())

		return ""
	}
	return profileResponse.UserID
}

func getContact(id string, accessToken string) *objects.FacebookUser {
	url := fmt.Sprintf("https://graph.instagram.com/v21.0/%s?fields=name&access_token=%s", id, accessToken)
	fmt.Println("FETCH USER PROFILE", url)
	resp, err := http.Get(url)
	if err == nil {

		var data objects.FacebookUser
		if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
			return &data
		}
	}
	defer resp.Body.Close()
	return nil
}
func getProfilePicture(id string, accessToken string) *objects.FacebookUser {
	url := fmt.Sprintf("https://graph.instagram.com/v21.0/%s?fields=name,profile_picture_url&access_token=%s", id, accessToken)
	fmt.Println("FETCH USER PROFILE", url)
	resp, err := http.Get(url)
	if err == nil {

		var data objects.FacebookUser
		if err := json.NewDecoder(resp.Body).Decode(&data); err == nil {
			return &data
		}
	}
	defer resp.Body.Close()
	return nil
}

func saveFileContenFromUrl(url string) (string, string, error) {
	filename := utils.GenerateRandomString(10)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return "", "", err
	}

	defer resp.Body.Close()
	byteValue, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	mtype := mimetype.Detect(byteValue)

	mimeType := mtype.String()
	switch mimeType {
	case "image/jpeg":
		filename = filename + ".jpg"
	case "image/png":
		filename = filename + ".png"
	}
	path := filepath.Join("assets/static", filename)
	os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err := os.WriteFile(path, byteValue, 0644); err != nil {
		log.Println(err)
		return "", "", err
	}
	return path, mimeType, nil
}

func reverseInstagram(messages []models.InstagramMessage) []models.InstagramMessage {
	for i, j := 0, len(messages)-1; i < len(messages)/2; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}
	return messages
}
