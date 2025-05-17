package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type FacebookHandler struct {
	ctx        *context.ERPContext
	appService *app.AppService
}

func NewFacebookHandler(ctx *context.ERPContext) *FacebookHandler {
	appService, ok := ctx.AppService.(*app.AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	return &FacebookHandler{ctx: ctx, appService: appService}
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

		utils.LogJson(response)
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
		var result facebookAccessTokenResponse
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

type facebookAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
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
