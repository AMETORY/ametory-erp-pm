package models

import (
	"ametory-pm/models/connection"
	"fmt"
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
	CompletedCount      int                          `json:"completed_count" gorm:"-"`
	TemplateID          *string                      `json:"template_id,omitempty"`
	Template            any                          `gorm:"-" json:"template,omitempty"`
	MemberID            *string                      `json:"member_id,omitempty" gorm:"column:member_id;constraint:OnDelete:CASCADE;"`
	Member              *models.MemberModel          `gorm:"foreignKey:MemberID" json:"member,omitempty"`
	DelayTime           int                          `json:"delay_time" gorm:"default:1000"`
	SequenceDelayTime   int                          `json:"sequence_delay_time" gorm:"default:0"`
	Files               []models.FileModel           `json:"files,omitempty" gorm:"-"`
	Products            []models.ProductModel        `gorm:"many2many:broadcast_products;" json:"products,omitempty"`
	LastBroadcastAt     *time.Time                   `json:"last_broadcast_at,omitempty" gorm:"-"`
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
		Success  int64 `json:"success"`
		Failed   int64 `json:"failed"`
		Complete int64 `json:"complete"`
	}
	var countData count
	tx.Model(&BroadcastContacts{}).Where("broadcast_model_id = ?", b.ID).Select("COUNT(CASE WHEN is_completed = 't' THEN 1 END) as complete, COUNT(CASE WHEN is_success = 't' THEN 1 END) as success, COUNT(CASE WHEN is_success = 'f' THEN 1 END) as failed").Scan(&countData)
	b.SuccessCount = int(countData.Success)
	b.FailedCount = int(countData.Failed)
	b.CompletedCount = int(countData.Complete)

	if b.TemplateID != nil {
		var template models.WhatsappMessageTemplate
		tx.Model(&models.WhatsappMessageTemplate{}).Where("id = ?", *b.TemplateID).First(&template)

		b.Template = template
	}

	var files []models.FileModel
	tx.Model(&models.FileModel{}).Where("ref_id = ? and ref_type = 'broadcast'", b.ID).Find(&files)
	b.Files = files

	var lastBroadcastAt time.Time
	tx.Model(&MessageLog{}).
		Where("broadcast_id = ?", b.ID).
		Order("created_at DESC").
		Limit(1).
		Select("created_at").
		Scan(&lastBroadcastAt)

	if lastBroadcastAt.IsZero() {
		b.LastBroadcastAt = nil
	} else {
		b.LastBroadcastAt = &lastBroadcastAt
	}

	var countLog int64
	tx.Model(&MessageLog{}).Where("broadcast_id = ?", b.ID).Count(&countLog)

	var contactCount int64
	tx.Model(&BroadcastContacts{}).Where("broadcast_model_id = ?", b.ID).Count(&contactCount)
	b.ContactCount = int(contactCount)
	fmt.Println("b.LastBroadcastAt", b.LastBroadcastAt)

	if countLog == 0 && b.ContactCount > 0 && time.Now().After(b.UpdatedAt.Add(time.Duration(b.DelayTime)*time.Second)) {
		b.Status = "NOT_STARTED"
		tx.Save(b)
	}
	if (b.Status == "PROCESSING" || b.Status == "RESTARTING" || b.Status == "STOPPED") && b.LastBroadcastAt != nil {
		if b.ContactCount > 0 && b.DelayTime > 0 {
			var expectedTime time.Time
			if b.LastBroadcastAt != nil {
				expectedTime = time.Date(
					b.LastBroadcastAt.Year(),
					b.LastBroadcastAt.Month(),
					b.LastBroadcastAt.Day(),
					b.LastBroadcastAt.Hour(),
					b.LastBroadcastAt.Minute(),
					b.LastBroadcastAt.Second(),
					0,
					b.LastBroadcastAt.Location()).Add(time.Duration(b.DelayTime) * time.Second)
			}

			fmt.Printf("EXPECTED TIME[%v] %v, NOW %v, \n", b.Description, expectedTime, time.Now())
			expectedTimeStr := fmt.Sprintf("atau %v hari %v jam %v menit",
				int(time.Since(expectedTime).Hours()/24),
				int(time.Since(expectedTime).Hours())%24,
				int(time.Since(expectedTime).Minutes())%60)
			fmt.Println(expectedTimeStr)
			if time.Since(expectedTime) > 3*24*time.Hour {
				b.Status = "EXPIRED"
				tx.Save(b)
			}

			// else if time.Now().After(expectedTime) && b.Status != "RESTARTING" {
			// 	b.Status = "STOPPED"
			// 	// tx.Save(b)
			// }

		}
	}

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
