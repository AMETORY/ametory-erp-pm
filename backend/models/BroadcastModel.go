package models

import (
	"ametory-pm/models/connection"
	"time"

	"github.com/AMETORY/ametory-erp-modules/shared"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BroadcastModel struct {
	shared.BaseModel
	Description string                        `json:"description"`
	Message     string                        `json:"message"`
	CompanyID   *string                       `gorm:"primaryKey" json:"company_id,omitempty"`
	Company     *models.CompanyModel          `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	ScheduledAt *time.Time                    `json:"scheduled_at,omitempty"`
	Status      string                        `json:"status" gorm:"default:DRAFT"`
	Connections []*connection.ConnectionModel `gorm:"many2many:broadcast_connections;" json:"connections,omitempty"`
	Contacts    []*models.ContactModel        `gorm:"many2many:broadcast_contacts;" json:"contacts,omitempty"`
}

func (b *BroadcastModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		tx.Statement.SetColumn("id", uuid.New().String())
	}
	return nil
}

func (BroadcastModel) TableName() string {
	return "broadcasts"
}

type BroadcastGrouping struct {
	shared.BaseModel
	Code        string `json:"code"`
	BroadcastID string `gorm:"size:36" json:"broadcast_id"`
}

type BroadcastContacts struct {
	BroadcastGroupingID string `gorm:"size:36" json:"broadcast_grouping_id"`
	ContactModelID      string `gorm:"size:36" json:"contact_model_id"`
	ConnectionModelID   string `gorm:"size:36" json:"connection_model_id"`
}
