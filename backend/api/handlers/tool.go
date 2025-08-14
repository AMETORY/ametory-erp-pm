package handlers

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
)

func SendWhatsappApiContactMessage(conn connection.ConnectionModel, contact models.ContactModel, message string, member *models.MemberModel) error {
	payload := map[string]interface{}{
		"messaging_product": "whatsapp",
		"recipient_type":    "individual",
		"to":                contact.Phone,
		"type":              "text",
		"text": map[string]interface{}{
			// "preview_url": true,
			"body": message,
		},
	}

	url := fmt.Sprintf("%s/%s/messages", config.App.Facebook.BaseURL, conn.Session)
	fmt.Println("URL", url, "\nPAYLOAD", payload)

	jsonPayload, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conn.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}

		fmt.Println("ERROR SEND WA MESSAGE", string(body))

		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	var waResponse whatsapp.WaResponse
	if err := json.NewDecoder(resp.Body).Decode(&waResponse); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	utils.LogJson(waResponse)
	return nil

}
