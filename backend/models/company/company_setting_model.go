package company

import (
	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
)

type CompanySetting struct {
	shared.BaseModel
	CompanyID             *string              `gorm:"primaryKey" json:"company_id,omitempty"`
	Company               *models.CompanyModel `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	GeminiAPIKey          string               `gorm:"type:text" json:"gemini_api_key"`
	WhatsappWebHost       string               `gorm:"type:text" json:"whatsapp_web_host"`
	WhatsappWebMockNumber string               `gorm:"type:text" json:"whatsapp_web_mock_number"`
	WhatsappWebIsMocked   string               `gorm:"type:text" json:"whatsapp_web_is_mocked"`
}
