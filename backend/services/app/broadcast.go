package app

import (
	"ametory-pm/models"
	"ametory-pm/models/connection"
	srv "ametory-pm/services"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/AMETORY/ametory-erp-modules/shared"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type BroadcastService struct {
	ctx        *context.ERPContext
	logHandler func(log models.MessageLog)
}

func NewBroadcastService(ctx *context.ERPContext) *BroadcastService {
	if !ctx.SkipMigration {
		ctx.DB.AutoMigrate(&models.BroadcastModel{}, &models.BroadcastGrouping{}, &models.BroadcastContacts{})
	}
	return &BroadcastService{
		ctx: ctx,
	}
}

func (s *BroadcastService) CreateBroadcast(broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Create(broadcast).Error
}

func (s *BroadcastService) GetBroadcasts(companyID string) ([]models.BroadcastModel, error) {
	var broadcasts []models.BroadcastModel
	if err := s.ctx.DB.Where("company_id = ?", companyID).Find(&broadcasts).Error; err != nil {
		return nil, err
	}
	return broadcasts, nil
}

func (s *BroadcastService) AddContact(broadcastID string, contacts []mdl.ContactModel) error {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", broadcastID).First(&broadcast).Error; err != nil {
		return err
	}

	return s.ctx.DB.Model(&broadcast).Association("Contacts").Append(contacts)
}

func (s *BroadcastService) AddConnection(broadcastID string, connections []connection.ConnectionModel) error {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", broadcastID).First(&broadcast).Error; err != nil {
		return err
	}

	return s.ctx.DB.Model(&broadcast).Association("Connections").Append(connections)
}

func (s *BroadcastService) GetBroadcastByID(id string) (*models.BroadcastModel, error) {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", id).First(&broadcast).Error; err != nil {
		return nil, err
	}
	return &broadcast, nil
}

func (s *BroadcastService) UpdateBroadcast(id string, broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Where("id = ?", id).Save(broadcast).Error
}

func (s *BroadcastService) DeleteBroadcast(id string) error {
	return s.ctx.DB.Delete(&models.BroadcastModel{}, "id = ?", id).Error
}

func (s *BroadcastService) Send(b *models.BroadcastModel) {
	if b.ScheduledAt == nil {
		key := fmt.Sprintf("broadcast:schedule:%v", b.ID)
		data, _ := json.Marshal(b)
		srv.REDIS.Set(*s.ctx.Ctx, key, data, time.Until(*b.ScheduledAt))
		b.Status = "SCHEDULED"
		s.ctx.DB.Save(b)
		go func() {
			time.Sleep(time.Until(*b.ScheduledAt))
			s.startBroadcast(b)
		}()
	} else {
		s.startBroadcast(b)
	}
}

func (s *BroadcastService) startBroadcast(b *models.BroadcastModel) {
	fmt.Println("📢 Starting broadcast", b.ID)

	batches := chunkContacts(b.Contacts, b.MaxContactsPerBatch)
	for i, batch := range batches {
		sender := b.Connections[i%len(b.Connections)]
		go s.sendBatchWithDelay(sender, b.ID, batch, 1*time.Second)
	}
}

func (s *BroadcastService) sendBatchWithDelay(sender connection.ConnectionModel, broadcastID string, contacts []mdl.ContactModel, delay time.Duration) {
	for _, contact := range contacts {
		sendWithRetryHandling(
			sender,
			broadcastID,
			contact,
			1,
			delay,
			s.logHandler,
			func(mr models.MessageRetry) {
				s.saveToRetryQueue(mr)
			},
		)
	}
}

func (s *BroadcastService) saveToRetryQueue(retry models.MessageRetry) {
	key := fmt.Sprintf("retry:sender:%d", retry.Sender.ID)
	data, _ := json.Marshal(retry)
	srv.REDIS.RPush(*s.ctx.Ctx, key, data)
}

func (s *BroadcastService) StartRetrySchedulers(b models.BroadcastModel) {
	for _, sender := range b.Connections {
		go s.scheduleRetrySender(sender)
	}
}

func (s *BroadcastService) scheduleRetrySender(sender connection.ConnectionModel) {
	c := cron.New()
	_, _ = c.AddFunc("*/1 * * * *", func() {
		ctx := s.ctx.Ctx
		key := fmt.Sprintf("retry:sender:%d", sender.ID)

		for {
			retryData, err := srv.REDIS.LPop(*ctx, key).Result()
			if err == redis.Nil {
				break
			} else if err != nil {
				break
			}

			var retryItem models.MessageRetry
			if err := json.Unmarshal([]byte(retryData), &retryItem); err != nil {
				continue
			}

			contact := mdl.ContactModel{
				BaseModel: shared.BaseModel{ID: retryItem.Contact.ID},
				Phone:     retryItem.Contact.Phone,
			}

			sendWithRetryHandling(
				retryItem.Sender,
				retryItem.BroadcastID,
				contact,
				retryItem.Attempt,
				1*time.Second,
				s.logHandler,
				func(mr models.MessageRetry) {
					s.saveToRetryQueue(mr)
				},
			)
		}
	})
	c.Start()
}

func chunkContacts(contacts []mdl.ContactModel, size int) [][]mdl.ContactModel {
	var batches [][]mdl.ContactModel
	for size < len(contacts) {
		contacts, batches = contacts[size:], append(batches, contacts[0:size:size])
	}
	return append(batches, contacts)
}

func sendWithRetryHandling(
	sender connection.ConnectionModel,
	broadcastID string,
	contact mdl.ContactModel,
	attempt int,
	delay time.Duration,
	logHandler func(log models.MessageLog),
	retryHandler func(retry models.MessageRetry),
) {
	time.Sleep(delay)

	success := simulateSend(contact)

	log := models.MessageLog{
		BroadcastID: broadcastID,
		ContactID:   contact.ID,
		SenderID:    sender.ID,
		SentAt:      time.Now(),
	}

	if success {
		log.Status = "success"
		logHandler(log)
	} else {
		if attempt >= 3 {
			log.Status = "undeliverable"
			log.ErrorMessage = fmt.Sprintf("attempt %d failed", attempt)
			logHandler(log)
		} else {
			log.Status = "failed"
			log.ErrorMessage = fmt.Sprintf("attempt %d failed", attempt)
			logHandler(log)

			retryHandler(models.MessageRetry{
				BroadcastID: broadcastID,
				Contact:     contact,
				Sender:      sender,
				Attempt:     attempt + 1,
				LastError:   log.ErrorMessage,
				LastTriedAt: time.Now(),
			})
		}
	}
}

func simulateSend(contact mdl.ContactModel) bool {

	fmt.Println("Simulate send to", contact.Phone)
	// 90% berhasil
	return rand.Intn(100) < 90
}
