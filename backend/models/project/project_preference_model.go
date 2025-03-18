package project

import (
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectPreferenceModel struct {
	shared.BaseModel
	ProjectID              string              `gorm:"type:char(36);not null" json:"project_id,omitempty"`
	Project                models.ProjectModel `json:"project,omitempty" gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE;"`
	RapidApiEnabled        sql.NullBool        `json:"rapid_api_enabled,omitempty" gorm:"type:bool;default:false"`
	ContactEnabled         sql.NullBool        `json:"contact_enabled,omitempty" gorm:"type:bool;default:false"`
	CustomAttributeEnabled sql.NullBool        `json:"custom_attribute_enabled,omitempty" gorm:"type:bool;default:false"`
	GeminiEnabled          sql.NullBool        `json:"gemini_enabled,omitempty" gorm:"type:bool;default:false"`
	FormEnabled            sql.NullBool        `json:"form_enabled,omitempty" gorm:"type:bool;default:false"`
}

func (ProjectPreferenceModel) TableName() string {
	return "project_preferences"
}

func (p *ProjectPreferenceModel) BeforeCreate(tx *gorm.DB) error {
	if p.ProjectID == "" {
		return errors.New("project_id is required")
	}
	if p.ID == "" {
		tx.Statement.SetColumn("id", uuid.New().String())
	}
	return nil
}

func (p ProjectPreferenceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ProjectID              string `json:"project_id"`
		RapidApiEnabled        bool   `json:"rapid_api_enabled"`
		ContactEnabled         bool   `json:"contact_enabled"`
		CustomAttributeEnabled bool   `json:"custom_attribute_enabled"`
		GeminiEnabled          bool   `json:"gemini_enabled"`
		FormEnabled            bool   `json:"form_enabled"`
	}{
		ProjectID:              p.ProjectID,
		RapidApiEnabled:        p.RapidApiEnabled.Bool,
		ContactEnabled:         p.ContactEnabled.Bool,
		CustomAttributeEnabled: p.CustomAttributeEnabled.Bool,
		GeminiEnabled:          p.GeminiEnabled.Bool,
		FormEnabled:            p.FormEnabled.Bool,
	})
}

func (p *ProjectPreferenceModel) UnmarshalJSON(data []byte) error {
	type AliasProjectPreferenceModel ProjectPreferenceModel
	var a struct {
		AliasProjectPreferenceModel
		ProjectID              string `json:"project_id,omitempty"`
		RapidApiEnabled        bool   `json:"rapid_api_enabled"`
		ContactEnabled         bool   `json:"contact_enabled"`
		CustomAttributeEnabled bool   `json:"custom_attribute_enabled"`
		GeminiEnabled          bool   `json:"gemini_enabled"`
		FormEnabled            bool   `json:"form_enabled"`
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*p = ProjectPreferenceModel(a.AliasProjectPreferenceModel)
	p.ProjectID = a.ProjectID
	p.RapidApiEnabled = sql.NullBool{
		Bool:  a.RapidApiEnabled,
		Valid: a.RapidApiEnabled,
	}
	p.ContactEnabled = sql.NullBool{
		Bool:  a.ContactEnabled,
		Valid: a.ContactEnabled,
	}
	p.CustomAttributeEnabled = sql.NullBool{
		Bool:  a.CustomAttributeEnabled,
		Valid: a.CustomAttributeEnabled,
	}
	p.GeminiEnabled = sql.NullBool{
		Bool:  a.GeminiEnabled,
		Valid: a.GeminiEnabled,
	}
	p.FormEnabled = sql.NullBool{
		Bool:  a.FormEnabled,
		Valid: a.FormEnabled,
	}

	return nil
}
