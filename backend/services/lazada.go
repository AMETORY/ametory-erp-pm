package services

import (
	"fmt"
	iop "iop-go-sdk/iop"

	"github.com/AMETORY/ametory-erp-modules/context"
)

type LazadaService struct {
	ctx       *context.ERPContext
	apiKey    string
	apiSecret string
	region    string
	client    *iop.IopClient
}

func NewLazadaService(ctx *context.ERPContext, apiKey, apiSecret, region, callbackUrl string) *LazadaService {
	var clientOptions = iop.ClientOptions{
		APIKey:    apiKey,
		APISecret: apiSecret,
		Region:    region,
	}
	client := iop.NewClient(&clientOptions)
	client.CallbackURL = callbackUrl

	return &LazadaService{
		ctx:       ctx,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		region:    region,
		client:    client,
	}
}

func (s *LazadaService) GetAuthURL() (string, error) {
	if s.client == nil {
		return "", fmt.Errorf("lazada client is not initialized")
	}
	return s.client.MakeAuthURL(), nil
}
