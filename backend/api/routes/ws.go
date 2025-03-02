package routes

import (
	"ametory-pm/config"
	"ametory-pm/services"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func SetupWSRoutes(r *gin.RouterGroup, erpContext *context.ERPContext) {
	r.GET("/ws/:channelId", func(c *gin.Context) {
		services.WS.HandleRequest(c.Writer, c.Request)
	})

	services.WS.HandleConnect(func(s *melody.Session) {
		userID, err := parseToken(s.Request.URL.Query().Get("token"))
		if err != nil {
			s.Close()
			return
		}
		msg := gin.H{
			"message":   "Connected",
			"sender_id": *userID,
		}
		b, _ := json.Marshal(msg)
		services.WS.BroadcastFilter(b, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
		fmt.Println("Connected", s.Request.URL.Path)
	})
	services.WS.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("Disconnected", s.Request.URL.Path)
	})

	services.WS.HandleMessage(func(s *melody.Session, msg []byte) {
		services.WS.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})
}

func parseToken(authToken string) (*string, error) {

	token, err := jwt.ParseWithClaims(authToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.App.Server.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}
	userID := token.Claims.(*jwt.StandardClaims).Id
	return &userID, nil
}
