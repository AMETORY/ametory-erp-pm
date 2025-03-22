package whatsapp

type WhatsAppMessage struct {
	Conversation       *string          `json:"conversation"`
	ImageMessage       *ImageMessage    `json:"imageMessage,omitempty"`
	VideoMessage       *VideoMessage    `json:"videoMessage,omitempty"`
	DocumentMessage    *DocumentMessage `json:"documentMessage,omitempty"`
	MessageContextInfo struct {
		DeviceListMetadata struct {
			SenderKeyHash      string `json:"senderKeyHash"`
			SenderTimestamp    int64  `json:"senderTimestamp"`
			RecipientKeyHash   string `json:"recipientKeyHash"`
			RecipientTimestamp int64  `json:"recipientTimestamp"`
		} `json:"deviceListMetadata"`
		DeviceListMetadataVersion int    `json:"deviceListMetadataVersion"`
		MessageSecret             string `json:"messageSecret"`
	} `json:"messageContextInfo"`
	ExtendedTextMessage *ExtendedTextMessage `json:"extendedTextMessage,omitempty"`
}

type ExtendedTextMessage struct {
	Text        string `json:"text"`
	ContextInfo struct {
		MentionedJID []string `json:"mentionedJID"`
	} `json:"contextInfo"`
}

type ImageMessage struct {
	URL                  string `json:"URL"`
	Mimetype             string `json:"mimetype"`
	Caption              string `json:"caption"`
	FileSHA256           string `json:"fileSHA256"`
	FileLength           int    `json:"fileLength"`
	Height               int    `json:"height"`
	Width                int    `json:"width"`
	MediaKey             string `json:"mediaKey"`
	FileEncSHA256        string `json:"fileEncSHA256"`
	DirectPath           string `json:"directPath"`
	MediaKeyTimestamp    int64  `json:"mediaKeyTimestamp"`
	JPEGThumbnail        string `json:"JPEGThumbnail"`
	FirstScanSidecar     string `json:"firstScanSidecar"`
	FirstScanLength      int    `json:"firstScanLength"`
	ScansSidecar         string `json:"scansSidecar"`
	ScanLengths          []int  `json:"scanLengths"`
	MidQualityFileSHA256 string `json:"midQualityFileSHA256"`
	ImageSourceType      int    `json:"imageSourceType"`
}

type VideoMessage struct {
	URL                string `json:"URL"`
	Mimetype           string `json:"mimetype"`
	Caption            string `json:"caption"`
	FileSHA256         string `json:"fileSHA256"`
	FileLength         int    `json:"fileLength"`
	Seconds            int    `json:"seconds"`
	MediaKey           string `json:"mediaKey"`
	Height             int    `json:"height"`
	Width              int    `json:"width"`
	FileEncSHA256      string `json:"fileEncSHA256"`
	DirectPath         string `json:"directPath"`
	MediaKeyTimestamp  int64  `json:"mediaKeyTimestamp"`
	JPEGThumbnail      string `json:"JPEGThumbnail"`
	StreamingSidecar   string `json:"streamingSidecar"`
	AccessibilityLabel string `json:"accessibilityLabel"`
}

type DocumentMessage struct {
	URL               string `json:"URL"`
	Mimetype          string `json:"mimetype"`
	Title             string `json:"title"`
	FileSHA256        string `json:"fileSHA256"`
	FileLength        int    `json:"fileLength"`
	MediaKey          string `json:"mediaKey"`
	FileName          string `json:"fileName"`
	FileEncSHA256     string `json:"fileEncSHA256"`
	DirectPath        string `json:"directPath"`
	MediaKeyTimestamp int64  `json:"mediaKeyTimestamp"`
	ContactVcard      bool   `json:"contactVcard"`
	Caption           string `json:"caption"`
}

type MsgObject struct {
	Message     WhatsAppMessage `json:"message"`
	Info        map[string]any  `json:"info"`
	Sender      string          `json:"sender"`
	JID         string          `json:"jid"`
	SessionID   string          `json:"session_id"`
	SessionName string          `json:"session_name"`
	MediaPath   string          `json:"media_path"`
	MimeType    string          `json:"mime_type"`
}
