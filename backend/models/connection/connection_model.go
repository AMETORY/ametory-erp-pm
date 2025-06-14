package connection

import (
	"encoding/json"
	"time"

	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConnectionType string

var (
	WhatsappWeb ConnectionType = "whatsapp_web"
)

type ConnectionModel struct {
	shared.BaseModel
	Name                  string               `json:"name" gorm:"column:name;not null;type:varchar(255)"`
	Description           string               `json:"description" gorm:"column:description;type:varchar(255)"`
	Type                  ConnectionType       `json:"type" gorm:"column:type;not null;type:varchar(255)"`
	Username              string               `json:"username" gorm:"column:username;type:varchar(255)"`
	SessionName           string               `json:"session_name" gorm:"column:session_name;type:varchar(255)"`
	Session               string               `json:"session" gorm:"column:session;type:varchar(255)"`
	Password              string               `json:"password" gorm:"column:password;type:varchar(255)"`
	ChannelID             string               `json:"channel_id" gorm:"column:channel_id;type:varchar(255)"`
	APIKey                string               `json:"api_key" gorm:"column:api_key;type:varchar(255)"`
	APISecret             string               `json:"api_secret" gorm:"column:api_secret;type:varchar(255)"`
	AccessToken           string               `json:"access_token" gorm:"column:access_token;type:varchar(255)"`
	AccessTokenExpiredAt  *time.Time           `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	RefreshToken          string               `json:"refresh_token" gorm:"column:refresh_token;type:varchar(255)"`
	Status                string               `json:"status" gorm:"column:status;type:varchar(255)"`
	GeminiAgentID         *string              `json:"gemini_agent_id" gorm:"column:gemini_agent_id;type:varchar(255)"`
	GeminiAgent           *models.GeminiAgent  `json:"gemini_agent" gorm:"foreignKey:GeminiAgentID;references:ID"`
	IsAutoPilot           bool                 `json:"is_auto_pilot" gorm:"column:is_auto_pilot;type:bool;default:false"`
	SessionAuth           bool                 `json:"session_auth" gorm:"column:session_auth;type:bool;default:false"`
	CompanyID             *string              `json:"company_id,omitempty" gorm:"column:company_id"`
	Company               *models.CompanyModel `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	Connected             bool                 `json:"connected" gorm:"-"`
	ProjectID             *string              `json:"project_id,omitempty" gorm:"column:project_id"`
	Project               *models.ProjectModel `gorm:"foreignKey:ProjectID" json:"project,omitempty"`
	NewSessionColumnID    *string              `json:"new_session_column_id,omitempty" gorm:"column:new_session_column_id"`
	NewSessionColumn      *models.ColumnModel  `gorm:"foreignKey:NewSessionColumnID" json:"new_session_column,omitempty"`
	IdleColumnID          *string              `json:"idle_column_id,omitempty" gorm:"column:idle_column_id"`
	IdleColumn            *models.ColumnModel  `gorm:"foreignKey:IdleColumnID" json:"idle_column,omitempty"`
	IdleDuration          float64              `json:"idle_duration" gorm:"column:idle_duration;default:7"`
	Color                 string               `json:"color" gorm:"column:color;type:varchar(255);default:'#34C759'"`
	AutoResponseStartTime *string              `json:"auto_response_start_time" gorm:"type:varchar(32)"`
	AutoResponseEndTime   *string              `json:"auto_response_end_time" gorm:"type:varchar(32)"`
	AutoResponseMessage   string               `json:"auto_response_message" gorm:"column:auto_response_message;type:varchar(255)"`
	AuthData              *json.RawMessage     `json:"auth_data" gorm:"type:json;default:'{}'"`
}

func (ConnectionModel) TableName() string {
	return "connections"
}

func (c *ConnectionModel) BeforeCreate(tx *gorm.DB) error {
	if c.ID == "" {
		tx.Statement.SetColumn("id", uuid.New().String())
	}
	return nil
}
