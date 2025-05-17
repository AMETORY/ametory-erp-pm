package handlers

import (
	"ametory-pm/models/connection"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/context"
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

func (h *FacebookHandler) FacebookCallbackHandler(c *gin.Context) {
	connectionID := c.Param("connectionID")
	code := c.Query("code")

	if code != "" {
		redirectUrl := fmt.Sprintf(`%s/%s`, h.appService.Config.Facebook.RedirectURL, connectionID)
		url := fmt.Sprintf(`https://graph.facebook.com/v18.0/oauth/access_token?client_id=%s&redirect_uri=%s&client_secret=%s&code=%s`, h.appService.Config.Facebook.AppID,
			redirectUrl,
			h.appService.Config.Facebook.AppSecret,
			code)
		fmt.Println("url", url)
		fmt.Println("connectionID", connectionID)
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

		var result facebookAccessTokenResponse
		if err := json.Unmarshal(bodyBytes, &result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to parse response body",
			})
			return
		}

		fmt.Println("connectionID", connectionID)

		fmt.Println("Parsed response:")

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

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "code is empty",
		})
	}

}

type facebookAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}
