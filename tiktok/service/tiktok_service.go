package service

import (
	"ametory-pm/services/app"
	"encoding/json"
	"errors"
	"fmt"

	golangCtx "context"
	"tiktokshop/open/sdk_golang/apis"
	customer_service_v202309 "tiktokshop/open/sdk_golang/models/customer_service/v202309"
	"tiktokshop/open/sdk_golang/objects"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
)

type TiktokService struct {
	ctx                         *context.ERPContext
	appService                  *app.AppService
	appKey                      string
	appSecret                   string
	serviceID                   string
	customerRelationshipService *customer_relationship.CustomerRelationshipService
}

func NewTiktokService(ctx *context.ERPContext,
	appService *app.AppService,
	customerRelationshipService *customer_relationship.CustomerRelationshipService,
	appKey string,
	appSecret string,
	serviceID string,
) *TiktokService {
	return &TiktokService{
		ctx:                         ctx,
		appService:                  appService,
		appKey:                      appKey,
		appSecret:                   appSecret,
		serviceID:                   serviceID,
		customerRelationshipService: customerRelationshipService,
	}
}

func (s *TiktokService) SetAuth(appKey, appSecret, serviceID string) {
	s.appKey = appKey
	s.appSecret = appSecret
	s.serviceID = serviceID
}

func (s *TiktokService) AuthorizeConnection(id, code string) (*apis.ResponseInfo, error) {
	if s.appKey == "" || s.appSecret == "" || s.serviceID == "" {
		return nil, errors.New("tiktok credentials not set")
	}

	fmt.Println("CONNECTION ID", id)
	fmt.Println("TIKTOK CODE", code)
	auth := apis.NewAccessToken(s.appKey, s.appSecret)
	resp, err := auth.GetToken(code)
	if err != nil {
		return nil, err

	}
	return resp, nil
}

// For more details about the SDK, refer to the documentation:
// https://partner.tiktokshop.com/docv2/page/67c83e0799a75104986ae498
func (s *TiktokService) Authorization202309GetAuthorizedShopsGet(accessToken string) ([]objects.Shop, error) {
	if s.appKey == "" || s.appSecret == "" || s.serviceID == "" {
		return nil, errors.New("tiktok credentials not set")
	}

	fmt.Println("ACCESS TOKEN", accessToken)
	configuration := apis.NewConfiguration()
	configuration.AddAppInfo(s.appKey, s.appSecret)
	apiClient := apis.NewAPIClient(configuration)
	request := apiClient.AuthorizationV202309API.Authorization202309ShopsGet(golangCtx.Background())
	request = request.XTtsAccessToken(accessToken)
	request = request.ContentType("application/json")
	resp, httpResp, err := request.Execute()
	if err != nil || httpResp.StatusCode != 200 {
		return nil, fmt.Errorf("productsRequest err:%v resbody:%s", err, httpResp.Body)
	}
	if resp == nil {
		fmt.Printf("response is nil")
		return nil, errors.New("response is nil")
	}
	if resp.GetCode() != 0 {

		return nil, fmt.Errorf("response business is error, errorCode:%d errorMessage:%s", resp.GetCode(), resp.GetMessage())
	}
	respDataJson, _ := json.MarshalIndent(resp.GetData(), "", "  ")
	// fmt.Println("response data:", string(respDataJson))

	var respJson objects.GetShopResponse
	if err := json.Unmarshal(respDataJson, &respJson); err != nil {
		return nil, err
	}
	return respJson.Shops, nil
}

// For more details about the SDK, refer to the documentation:
// https://partner.tiktokshop.com/docv2/page/67c83e0799a75104986ae498
func (s *TiktokService) CustomerService202309GetConversationsGet(accessToken, shopChiper, pageToken string, limit int32) (*customer_service_v202309.CustomerService202309GetConversationsResponseData, error) {
	if s.appKey == "" || s.appSecret == "" || s.serviceID == "" {
		return nil, errors.New("tiktok credentials not set")
	}
	configuration := apis.NewConfiguration()
	configuration.AddAppInfo(s.appKey, s.appSecret)
	apiClient := apis.NewAPIClient(configuration)
	request := apiClient.CustomerServiceV202309API.CustomerService202309ConversationsGet(golangCtx.Background())
	request = request.XTtsAccessToken(accessToken)
	request = request.ContentType("application/json")
	request = request.PageToken(pageToken)
	request = request.PageSize(limit)
	request = request.Locale("id")
	request = request.ShopCipher(shopChiper)
	resp, httpResp, err := request.Execute()
	if err != nil || httpResp.StatusCode != 200 {
		fmt.Printf("productsRequest err:%v resbody:%s", err, httpResp.Body)
		return nil, fmt.Errorf("productsRequest err:%v resbody:%s", err, httpResp.Body)
	}
	if resp == nil {
		fmt.Printf("response is nil")
		return nil, errors.New("response is nil")
	}
	if resp.GetCode() != 0 {
		fmt.Printf("response business is error, errorCode:%d errorMessage:%s", resp.GetCode(), resp.GetMessage())
		return nil, fmt.Errorf("response business is error, errorCode:%d errorMessage:%s", resp.GetCode(), resp.GetMessage())
	}
	// respDataJson, _ := json.MarshalIndent(resp.GetData(), "", "  ")
	// fmt.Println("response data:", string(respDataJson))
	respData := resp.GetData()
	return &respData, nil
}

// For more details about the SDK, refer to the documentation:
// https://partner.tiktokshop.com/docv2/page/67c83e0799a75104986ae498
func (s *TiktokService) CustomerService202309GetConversationMessagesGet(accessToken, shopChiper, conversationId, pageToken string, limit int32) (*customer_service_v202309.CustomerService202309GetConversationMessagesResponseData, error) {
	if s.appKey == "" || s.appSecret == "" || s.serviceID == "" {
		return nil, errors.New("tiktok credentials not set")
	}

	fmt.Println("CONVERSATION ID", conversationId)
	configuration := apis.NewConfiguration()
	configuration.AddAppInfo(s.appKey, s.appSecret)
	apiClient := apis.NewAPIClient(configuration)
	request := apiClient.CustomerServiceV202309API.CustomerService202309ConversationsConversationIdMessagesGet(golangCtx.Background(), conversationId)
	request = request.XTtsAccessToken(accessToken)
	request = request.ContentType("application/json")
	request = request.PageToken(pageToken)
	request = request.PageSize(limit)
	request = request.Locale("id")
	request = request.ShopCipher(shopChiper)
	request = request.SortOrder("DESC")
	request = request.SortField("create_time")
	resp, httpResp, err := request.Execute()
	if err != nil || httpResp.StatusCode != 200 {
		fmt.Printf("productsRequest err:%v resbody:%s", err, httpResp.Body)
		return nil, fmt.Errorf("productsRequest err:%v resbody:%s", err, httpResp.Body)
	}
	if resp == nil {
		fmt.Printf("response is nil")
		return nil, errors.New("response is nil")
	}
	if resp.GetCode() != 0 {
		fmt.Printf("response business is error, errorCode:%d errorMessage:%s", resp.GetCode(), resp.GetMessage())
		return nil, fmt.Errorf("response business is error, errorCode:%d errorMessage:%s", resp.GetCode(), resp.GetMessage())
	}
	// respDataJson, _ := json.MarshalIndent(resp.GetData(), "", "  ")
	// fmt.Println("response data:", string(respDataJson))
	respData := resp.GetData()
	return &respData, nil
}
