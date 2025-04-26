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
	Description         string                       `json:"description"`
	Message             string                       `json:"message"`
	CompanyID           *string                      `json:"company_id,omitempty"`
	Company             *models.CompanyModel         `gorm:"foreignKey:CompanyID" json:"company,omitempty"`
	ScheduledAt         *time.Time                   `json:"scheduled_at,omitempty"`
	Status              string                       `json:"status" gorm:"default:DRAFT"`
	Connections         []connection.ConnectionModel `gorm:"many2many:broadcast_connections;" json:"connections,omitempty"`
	Contacts            []models.ContactModel        `gorm:"many2many:broadcast_contacts;" json:"contacts,omitempty"`
	MaxContactsPerBatch int                          `json:"max_contacts_per_batch" gorm:"default:100"`
	Groups              []BroadcastGrouping          `gorm:"foreignKey:BroadcastID" json:"groups,omitempty"`
	ContactCount        int                          `json:"contact_count" gorm:"-"`
	GroupCount          int                          `json:"group_count" gorm:"-"`
	SuccessCount        int                          `json:"success_count" gorm:"-"`
	FailedCount         int                          `json:"failed_count" gorm:"-"`
}

func (b *BroadcastModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == "" {
		tx.Statement.SetColumn("id", uuid.New().String())
	}
	return nil
}

func (b *BroadcastModel) AfterFind(tx *gorm.DB) error {
	// var completedCount int64
	// tx.Model(&BroadcastContacts{}).Where(" broadcast_model_id = ?", b.ID).Count(&completedCount)
	// if int(completedCount) == b.ContactCount {
	// 	b.Status = "COMPLETED"
	// 	return tx.Save(b).Error
	// }
	var countGroups int64
	tx.Model(&BroadcastGrouping{}).Where(" broadcast_id = ?", b.ID).Count(&countGroups)
	b.GroupCount = int(countGroups)
	type count struct {
		Success int64 `json:"success"`
		Failed  int64 `json:"failed"`
	}
	var countData count
	tx.Model(&BroadcastContacts{}).Where("broadcast_model_id = ?", b.ID).Select("COUNT(CASE WHEN is_success = 't' THEN 1 END) as success, COUNT(CASE WHEN is_success = 'f' THEN 1 END) as failed").Scan(&countData)
	b.SuccessCount = int(countData.Success)
	b.FailedCount = int(countData.Failed)

	return nil
}

func (BroadcastModel) TableName() string {
	return "broadcasts"
}

type BroadcastGrouping struct {
	shared.BaseModel
	Code        string         `json:"code"`
	BroadcastID string         `json:"broadcast_id"`
	Broadcast   BroadcastModel `gorm:"foreignKey:BroadcastID" json:"broadcast,omitempty"`
}

type BroadcastContacts struct {
	BroadcastGroupingID string `gorm:"size:36" json:"broadcast_grouping_id"`
	ConnectionModelID   string `gorm:"size:36" json:"connection_model_id"`
	ContactModelID      string `gorm:"uniqueIndex:idx_broadcast_contact;type:char(36)" json:"contact_model_id"`
	BroadcastModelID    string `gorm:"uniqueIndex:idx_broadcast_contact;type:char(36)" json:"broadcast_model_id"`
	IsCompleted         bool   `json:"is_completed"`
	IsSuccess           bool   `json:"is_success"`
}

type MessageLog struct {
	shared.BaseModel
	BroadcastID  string                     `json:"broadcast_id"`
	Broadcast    BroadcastModel             `gorm:"foreignKey:BroadcastID" json:"broadcast,omitempty"`
	ContactID    string                     `json:"contact_id"`
	Contact      models.ContactModel        `gorm:"foreignKey:ContactID" json:"contact"`
	SenderID     string                     `json:"sender_id"`
	Sender       connection.ConnectionModel `gorm:"foreignKey:SenderID" json:"sender"`
	Status       string                     `json:"status"` // "success" or "failed"
	ErrorMessage string                     `json:"error_message,omitempty"`
	SentAt       time.Time                  `json:"sent_at"`
}

func (MessageLog) TableName() string {
	return "broadcast_message_logs"
}

type MessageRetry struct {
	shared.BaseModel
	BroadcastID string                     `json:"broadcast_id"`
	Broadcast   BroadcastModel             `gorm:"foreignKey:BroadcastID" json:"broadcast,omitempty"`
	ContactID   string                     `json:"contact_id"`
	Contact     models.ContactModel        `gorm:"foreignKey:ContactID" json:"contact"`
	SenderID    string                     `json:"sender_id"`
	Sender      connection.ConnectionModel `gorm:"foreignKey:SenderID" json:"sender"`
	Attempt     int                        `json:"attempt"`
	LastError   string                     `json:"last_error"`
	LastTriedAt time.Time                  `json:"last_tried_at"`
}

func (MessageRetry) TableName() string {
	return "broadcast_message_retries"
}

type CustomContactModel struct {
	models.ContactModel
	IsCompleted bool           `json:"is_completed"`
	IsSuccess   bool           `json:"is_success"`
	Retries     []MessageRetry `gorm:"-" json:"retries,omitempty"`
	MessageLog  MessageLog     `gorm:"-" json:"message_log,omitempty"`
}
