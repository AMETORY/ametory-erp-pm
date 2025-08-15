package app

import (
	"ametory-pm/config"
	"ametory-pm/models/connection"
	"ametory-pm/models/whatsapp"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"

	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"moul.io/http2curl"
)

func SendWhatsappApiContactMessage(conn connection.ConnectionModel, contact models.ContactModel, message string, member *models.MemberModel, file []models.FileModel) error {
	imgID := ""
	if len(file) > 0 {
		for _, f := range file {

			imageId, err := SendWhatsappApiImage(conn, contact, f.Path, f.MimeType, member)
			if err != nil {
				return err
			}

			fmt.Println("IMAGE ID", *imageId)
			imgID = *imageId
		}

	}
	var payload map[string]any
	if imgID != "" {
		payload = map[string]any{
			"messaging_product": "whatsapp",
			"recipient_type":    "individual",
			"to":                contact.Phone,
			"type":              "image",
			"image": map[string]any{
				"id":      imgID,
				"caption": message,
			},
		}
	} else {
		payload = map[string]any{
			"messaging_product": "whatsapp",
			"recipient_type":    "individual",
			"to":                contact.Phone,
			"type":              "text",
			"text": map[string]any{
				"body": message,
			},
		}
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

	// utils.LogJson(waResponse)

	return nil

}

func SendWhatsappApiImage(conn connection.ConnectionModel, contact models.ContactModel, filePath, mimeType string, member *models.MemberModel) (*string, error) {
	url := fmt.Sprintf("%s/%s/media", config.App.Facebook.BaseURL, conn.Session)

	// Buat buffer untuk menampung body permintaan
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Tambahkan field 'messaging_product'
	_ = writer.WriteField("messaging_product", "whatsapp")

	// filePath, err := filepath.Abs(filePath)
	// if err != nil {
	// 	return nil, err
	// }
	fmt.Println("UPLOAD", filePath)
	// Buka file yang akan diunggah
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Gagal membuka file:", err)
		return nil, err
	}
	defer file.Close()

	// Buat form-data untuk file
	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="file"; filename="%s"`, filepath.Base(filePath)))
	header.Set("Content-Type", mimeType)

	// Buat bagian form-data menggunakan header yang sudah didefinisikan
	filePart, err := writer.CreatePart(header)
	if err != nil {
		return nil, err
	}

	// Salin isi file ke bagian form-data
	_, err = io.Copy(filePart, file)
	if err != nil {
		fmt.Println("Gagal menyalin file ke form-data:", err)
		return nil, err
	}

	// Selesaikan penulisan form-data
	writer.Close()

	// Buat URL endpoint

	// Buat permintaan HTTP POST
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		fmt.Println("Gagal membuat permintaan:", err)
		return nil, err
	}

	// Tambahkan header Authorization dan Content-Type
	req.Header.Set("Authorization", "Bearer "+conn.AccessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	fmt.Println("Authorization", conn.AccessToken)
	fmt.Println("Content-Type", writer.FormDataContentType())
	// Kirim permintaan
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Gagal mengirim permintaan:", err)
		return nil, err
	}
	defer resp.Body.Close()

	command, _ := http2curl.GetCurlCommand(req)
	fmt.Println("CURL", command)
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %v", err)
		}

		fmt.Println("ERROR SEND WA MESSAGE", string(body))

		return nil, fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}
	// Baca dan cetak respons
	fmt.Println("Status respons:", resp.Status)

	var waResponse struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&waResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	return &waResponse.ID, nil
}
