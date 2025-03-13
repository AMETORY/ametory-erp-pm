package rapid_api_models

import (
	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
)

type RapidApiPlugin struct {
	shared.BaseModel
	Name              string             `json:"name" gorm:"not null;type:varchar(100)"`
	Key               string             `json:"key" gorm:"not null;type:varchar(100)"`
	URL               string             `json:"url" gorm:"not null;type:varchar(255)"`
	IsActive          bool               `json:"is_active" gorm:"default:true"`
	RapidApiEndpoints []RapidApiEndpoint `json:"rapid_api_endpoints" gorm:"foreignKey:RapidApiPluginID"`
}

type RapidApiEndpoint struct {
	shared.BaseModel
	Title            string                   `json:"title" gorm:"not null;type:varchar(100)"`
	Key              string                   `json:"key" gorm:"not null;type:varchar(100)"`
	Method           string                   `json:"method" gorm:"not null;type:varchar(100)"`
	Params           string                   `json:"params" gorm:"not null;type:JSON"`
	URL              string                   `json:"url" gorm:"not null;type:varchar(255)"`
	RapidApiPluginID string                   `json:"rapid_api_plugin_id" gorm:"not null;foreignKey:RapidApiPluginID"`
	ParamData        []RapidApiEndpointParams `gorm:"-"`
}

type RapidApiData struct {
	shared.BaseModel
	CompanyID          string              `json:"company_id" gorm:"not null"`
	Company            models.CompanyModel `gorm:"foreignKey:CompanyID"`
	Title              string              `json:"title" gorm:"not null;type:varchar(100)"`
	RapidApiEndpointID string              `json:"rapid_api_endpoint_id" gorm:"not null;foreignKey:RapidApiEndpointID"`
	RapidApiEndpoint   RapidApiEndpoint    `gorm:"foreignKey:RapidApiEndpointID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	RapidApiPluginID   string              `json:"rapid_api_plugin_id" gorm:"not null;foreignKey:RapidApiPluginID"`
	RapidApiPlugin     RapidApiPlugin      `gorm:"foreignKey:RapidApiPluginID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Data               string              `json:"data" gorm:"type:JSON"`
	Params             string              `json:"params" gorm:"type:JSON"`
	TaskID             string              `json:"task_id" gorm:"not null"`
	Task               models.TaskModel    `gorm:"foreignKey:TaskID" json:"task"`
	ThumbnailURL       string              `json:"thumbnail_url"`
}

type RapidApiEndpointParams struct {
	Key  string `json:"key"`
	Type string `json:"type"`
}

type CompanyRapidApiPlugin struct {
	CompanyID        string              `gorm:"primaryKey" json:"company_id,omitempty"`
	Company          models.CompanyModel `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	RapidApiPluginID string              `gorm:"primaryKey" json:"rapid_api_plugin_id,omitempty"`
	RapidApiPlugin   RapidApiPlugin      `gorm:"foreignKey:RapidApiPluginID" json:"rapid_api_plugin,omitempty"`
	RapidApiKey      string              `json:"rapid_api_key,omitempty"`
	RapidApiHost     string              `json:"rapid_api_host,omitempty"`
	Endpoints        []RapidApiEndpoint  `gorm:"-" json:"endpoints,omitempty"`
}
