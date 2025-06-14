package services

import (
	"ametory-pm/models/connection"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/utils"
)

type ShopeeService struct {
	ctx         *context.ERPContext
	apiSecret   string
	partnerID   string
	host        string
	redirectUrl string
}

func NewShopeeService(ctx *context.ERPContext, apiSecret, partnerID, host, redirectUrl string) *ShopeeService {
	return &ShopeeService{ctx: ctx,
		apiSecret:   apiSecret,
		partnerID:   partnerID,
		host:        host,
		redirectUrl: redirectUrl,
	}
}

func (s *ShopeeService) GetHost() string {
	return s.host
}
func (s *ShopeeService) GetRedirectUrl() string {
	return s.redirectUrl
}
func (s *ShopeeService) GetPartnerID() string {
	return s.partnerID
}
func (s *ShopeeService) GetApiSecret() string {
	return s.apiSecret
}

func (s ShopeeService) AuthShop() (string, error) {
	if s.apiSecret == "" || s.host == "" || s.redirectUrl == "" || s.partnerID == "" {
		return "", fmt.Errorf("apiSecret or host or redirectUrl or partnerID is empty")
	}
	timest := strconv.FormatInt(time.Now().Unix(), 10)
	host := s.host
	path := "/api/v2/shop/auth_partner"
	redirectUrl := s.redirectUrl
	partnerId := s.partnerID
	partnerKey := s.apiSecret
	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timest)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))
	url := fmt.Sprintf(host+path+"?partner_id=%s&timestamp=%s&sign=%s&redirect=%s", partnerId, timest, sign, redirectUrl)
	fmt.Println(url)

	return url, nil
}

func (s ShopeeService) GetAccessTokenFromConnection(connection *connection.ConnectionModel) (string, error) {
	if connection.AccessTokenExpiredAt.Before(time.Now()) {
		resp, err := s.GetRefreshToken(connection.RefreshToken, connection.Username)
		if err != nil {
			return "", err
		}
		expAt := time.Now().Add(time.Duration(resp.ExpireIn) * time.Second)
		connection.AccessToken = resp.AccessToken
		connection.RefreshToken = resp.RefreshToken
		connection.AccessTokenExpiredAt = &expAt
		err = s.ctx.DB.Save(&connection).Error
		if err != nil {
			return "", err
		}
		return resp.AccessToken, nil
	}

	return connection.AccessToken, nil
}

func (s ShopeeService) GetRefreshToken(refreshToken, shopID string) (*TokenResponse, error) {
	if s.apiSecret == "" || s.host == "" || s.partnerID == "" {
		return nil, fmt.Errorf("apiSecret or host or partnerID is empty")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	path := "/api/v2/auth/access_token/get"
	partnerId := s.partnerID
	partnerKey := s.apiSecret
	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timestamp)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	endpointUrl := fmt.Sprintf("%s%s?partner_id=%s&timestamp=%s&sign=%s", s.host, path, partnerId, timestamp, sign)
	println(endpointUrl)
	intPartnerID, _ := strconv.Atoi(partnerId)
	intShopID, _ := strconv.Atoi(shopID)
	data := map[string]any{
		"refresh_token": refreshToken,
		"shop_id":       intShopID,
		"partner_id":    intPartnerID,
	}

	utils.LogJson(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpointUrl, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get auth token, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
func (s ShopeeService) GetAuthToken(authCode, shopID string) (*TokenResponse, error) {
	if s.apiSecret == "" || s.host == "" || s.partnerID == "" {
		return nil, fmt.Errorf("apiSecret or host or partnerID is empty")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	path := "/api/v2/auth/token/get"
	partnerId := s.partnerID
	partnerKey := s.apiSecret
	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timestamp)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	endpointUrl := fmt.Sprintf("%s%s?partner_id=%s&timestamp=%s&sign=%s", s.host, path, partnerId, timestamp, sign)
	println(endpointUrl)
	intPartnerID, _ := strconv.Atoi(partnerId)
	intShopID, _ := strconv.Atoi(shopID)
	data := map[string]any{
		"code":       authCode,
		"shop_id":    intShopID,
		"partner_id": intPartnerID,
	}

	utils.LogJson(data)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpointUrl, bytes.NewReader(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get auth token, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

func (s ShopeeService) GetShopProfile(accessToken, shopID string) (*ShopProfileResponse, error) {
	if s.apiSecret == "" || s.host == "" || s.partnerID == "" {
		return nil, fmt.Errorf("apiSecret or host or partnerID is empty")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	path := "/api/v2/shop/get_profile"
	partnerId := s.partnerID
	partnerKey := s.apiSecret
	baseString := fmt.Sprintf("%s%s%s%s%s", partnerId, path, timestamp, accessToken, shopID)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	url := fmt.Sprintf("%s%s?access_token=%s&partner_id=%s&shop_id=%s&sign=%s&timestamp=%s", s.host, path, accessToken, partnerId, shopID, sign, timestamp)
	println(url)
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get shop profile, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var shopResponse ShopProfileResponse
	if err := json.Unmarshal(body, &shopResponse); err != nil {
		return nil, err
	}

	return &shopResponse, nil
}
func (s ShopeeService) GetShopInfo(accessToken, shopID string) (*ShopInfoResponse, error) {
	if s.apiSecret == "" || s.host == "" || s.partnerID == "" {
		return nil, fmt.Errorf("apiSecret or host or partnerID is empty")
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	path := "/api/v2/shop/get_shop_info"
	partnerId := s.partnerID
	partnerKey := s.apiSecret
	baseString := fmt.Sprintf("%s%s%s%s%s", partnerId, path, timestamp, accessToken, shopID)
	h := hmac.New(sha256.New, []byte(partnerKey))
	h.Write([]byte(baseString))
	sign := hex.EncodeToString(h.Sum(nil))

	url := fmt.Sprintf("%s%s?access_token=%s&partner_id=%s&shop_id=%s&sign=%s&timestamp=%s", s.host, path, accessToken, s.partnerID, shopID, sign, timestamp)
	println(url)
	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get shop info, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var shopResponse ShopInfoResponse
	if err := json.Unmarshal(body, &shopResponse); err != nil {
		return nil, err
	}

	return &shopResponse, nil
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	ExpireIn     int    `json:"expire_in"`
	RequestID    string `json:"request_id"`
	Error        string `json:"error"`
	Message      string `json:"message"`
}

type ShopProfileResponse struct {
	RequestID string `json:"request_id"`
	Response  struct {
		ShopLogo    string `json:"shop_logo"`
		Description string `json:"description"`
		ShopName    string `json:"shop_name"`
	} `json:"response"`
}

type ShopInfoResponse struct {
	ShopName            string  `json:"shop_name"`
	Region              string  `json:"region"`
	Status              string  `json:"status"`
	IsSIP               bool    `json:"is_sip"`
	IsCB                bool    `json:"is_cb"`
	IsCNSC              bool    `json:"is_cnsc"`
	RequestID           string  `json:"request_id"`
	AuthTime            int     `json:"auth_time"`
	ExpireTime          int     `json:"expire_time"`
	Error               string  `json:"error"`
	Message             string  `json:"message"`
	ShopCBSC            string  `json:"shop_cbsc"`
	Is3PF               bool    `json:"is_3pf"`
	IsUpgradedCBSC      bool    `json:"is_upgraded_cbsc"`
	MerchantID          *string `json:"merchant_id"`
	ShopFulfillmentFlag string  `json:"shop_fulfillment_flag"`
}
