package app

import (
	"ametory-pm/config"
	"ametory-pm/models"
	"ametory-pm/models/connection"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	"github.com/AMETORY/ametory-erp-modules/shared"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/meta"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
)

type BroadcastService struct {
	ctx                         *context.ERPContext
	appService                  *AppService
	whatsmeowService            *whatsmeow_client.WhatsmeowService
	customerRelationshipService *customer_relationship.CustomerRelationshipService
	metaService                 *meta.MetaService
}

func NewBroadcastService(ctx *context.ERPContext) *BroadcastService {
	if !ctx.SkipMigration {
		ctx.DB.AutoMigrate(&models.BroadcastModel{},
			&models.BroadcastGrouping{},
			&models.BroadcastContacts{},
			&models.MessageLog{},
			&models.MessageRetry{},
		)
	}
	appService, ok := ctx.AppService.(*AppService)
	if !ok {
		panic("AppService is not instance of app.AppService")
	}
	whatsmeowService, ok := ctx.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
	if !ok {
		panic("ThirdPartyServices is not instance of whatsmeow_client.WhatsmeowService")
	}
	var customerRelationshipService *customer_relationship.CustomerRelationshipService
	customerRelationshipSrv, ok := ctx.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	if ok {
		customerRelationshipService = customerRelationshipSrv
	}
	metaService, ok := ctx.ThirdPartyServices["Meta"].(*meta.MetaService)
	if !ok {
		panic("MetaService is not instance of meta.MetaService")
	}
	return &BroadcastService{
		ctx:                         ctx,
		appService:                  appService,
		whatsmeowService:            whatsmeowService,
		customerRelationshipService: customerRelationshipService,
		metaService:                 metaService,
	}
}

func (s *BroadcastService) CreateBroadcast(broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Create(broadcast).Error
}

func (s *BroadcastService) GetBroadcasts(pagination *Pagination, httpRequest http.Request, search string) ([]models.BroadcastModel, error) {
	var broadcasts []models.BroadcastModel
	db := s.ctx.DB.Scopes(paginate(broadcasts, pagination, s.ctx.DB))
	if err := db.Where("company_id = ?", httpRequest.Header.Get("ID-Company")).Order("created_at DESC").Find(&broadcasts).Error; err != nil {
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

func (s *BroadcastService) DeleteContact(broadcastID string, contactID string) error {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", broadcastID).First(&broadcast).Error; err != nil {
		return err
	}

	return s.ctx.DB.Model(&broadcast).Association("Contacts").Delete(mdl.ContactModel{BaseModel: shared.BaseModel{ID: contactID}})
}

func (s *BroadcastService) DeleteContactByIDs(broadcastID string, contactIDs []string) error {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", broadcastID).First(&broadcast).Error; err != nil {
		return err
	}

	contacts := make([]mdl.ContactModel, len(contactIDs))
	for i, id := range contactIDs {
		contacts[i] = mdl.ContactModel{BaseModel: shared.BaseModel{ID: id}}
	}

	return s.ctx.DB.Model(&broadcast).Association("Contacts").Delete(contacts)
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
	if err := s.ctx.DB.Preload("Products").Preload("Connections").Preload("Member.User").Where("id = ?", id).First(&broadcast).Error; err != nil {
		return nil, err
	}

	for i, v := range broadcast.Products {
		v.ProductImages, _ = s.ListImagesOfProduct(v.ID)
		broadcast.Products[i] = v
	}

	return &broadcast, nil
}

func (s *BroadcastService) ListImagesOfProduct(productID string) ([]mdl.FileModel, error) {
	var images []mdl.FileModel
	err := s.ctx.DB.Where("ref_id = ? and ref_type = ?", productID, "product").Find(&images).Error
	return images, err
}

func (s *BroadcastService) GetContacts(id string, pagination *Pagination, search string) ([]mdl.ContactModel, error) {

	var contacts []mdl.ContactModel
	var selectContacts []models.CustomContactModel
	var totalRows int64
	db := s.ctx.DB.Model(&mdl.ContactModel{}).Preload("Tags").
		Joins("JOIN broadcast_contacts on broadcast_contacts.contact_model_id = contacts.id").
		Joins("JOIN broadcasts on broadcasts.id = broadcast_contacts.broadcast_model_id").
		Where("(contacts.name LIKE ? OR contacts.phone LIKE ?)", "%"+search+"%", "%"+search+"%").
		Where("broadcasts.id  = ?", id).
		Select("contacts.*", "broadcast_contacts.is_completed", "broadcast_contacts.is_success")

	err := db.Count(&totalRows).Error
	if err != nil {
		return nil, err
	}
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages

	if err := db.
		Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort()).
		Scan(&selectContacts).
		Error; err != nil {
		return nil, err
	}

	for _, contact := range selectContacts {
		var newContact mdl.ContactModel = contact.ContactModel
		var retry models.MessageRetry
		var logs []models.MessageLog

		s.ctx.DB.Model(&models.MessageRetry{}).Where("contact_id = ? and broadcast_id = ?", contact.ID, id).Find(&retry)
		s.ctx.DB.Model(&models.MessageLog{}).Where("contact_id = ? and broadcast_id = ?", contact.ID, id).Find(&logs)
		newContact.IsCompleted = contact.IsCompleted
		newContact.IsSuccess = contact.IsSuccess

		var data map[string]any = make(map[string]any)
		data["retry"] = retry
		data["logs"] = logs
		newContact.Data = data

		contacts = append(contacts, newContact)
	}

	return contacts, nil
}

func (s *BroadcastService) UpdateBroadcast(id string, broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Where("id = ?", id).Updates(broadcast).Error
}

func (s *BroadcastService) DeleteBroadcast(id string) error {
	return s.ctx.DB.Delete(&models.BroadcastModel{}, "id = ?", id).Error
}

func (s *BroadcastService) Send(b *models.BroadcastModel) {
	if b.ScheduledAt != nil {
		// key := fmt.Sprintf("broadcast:schedule:%v", b.ID)
		data, _ := json.Marshal(b)
		b.Status = "SCHEDULED"
		s.ctx.DB.Save(b)
		s.appService.Redis.Publish(*s.ctx.Ctx, "BROADCAST:SCHEDULED", data)
		// s.appService.Redis.Set(*s.ctx.Ctx, key, data, time.Until(*b.ScheduledAt))
		// go func() {
		// 	time.Sleep(time.Until(*b.ScheduledAt))
		// 	b.Status = "PROCESSING"
		// 	err := s.ctx.DB.First(&b, "id = ?", b.ID).Error
		// 	if err != nil {
		// 		log.Println("ERROR", err)
		// 		return
		// 	}
		// 	s.ctx.DB.Save(b)
		// 	s.StartBroadcast(b)
		// }()
	} else {
		b.Status = "PROCESSING"
		data, _ := json.Marshal(b)
		s.ctx.DB.Save(b)
		s.appService.Redis.Publish(*s.ctx.Ctx, "BROADCAST:NOW", data)

	}
}

func (s *BroadcastService) StartBroadcast(b *models.BroadcastModel) {
	log.Println("📢 Starting broadcast", b.ID)

	batches := chunkContacts(b.Contacts, b.MaxContactsPerBatch)
	log.Println("📢 Number of batches", len(batches), "<>", b.MaxContactsPerBatch)
	// utils.LogJson((batches))
	for i, batch := range batches {
		sender := b.Connections[i%len(b.Connections)]
		var group = models.BroadcastGrouping{
			BaseModel:   shared.BaseModel{ID: uuid.New().String()},
			BroadcastID: b.ID,
			Code:        utils.GenerateRandomNumber(6),
		}
		s.ctx.DB.Create(&group)
		for _, v := range batch {
			s.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ?", v.ID).Updates(map[string]any{
				"broadcast_grouping_id": group.ID,
				"connection_model_id":   sender.ID,
			})
		}
		if (b.SequenceDelayTime > 0) && (i > 0) {
			time.Sleep(time.Duration(b.SequenceDelayTime) * time.Second)
		}

		s.sendBatchWithDelay(sender, b.ID, batch, time.Duration(b.DelayTime)*time.Second)
	}
	s.StartRetrySchedulers(*b)
}

func (s *BroadcastService) sendBatchWithDelay(sender connection.ConnectionModel, broadcastID string, contacts []mdl.ContactModel, delay time.Duration) {
	for _, contact := range contacts {
		s.sendWithRetryHandling(
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
		// s.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("connection_model_id", sender.ID)
	}
}

func (s *BroadcastService) saveToRetryQueue(retry models.MessageRetry) {
	key := fmt.Sprintf("retry:sender:%v", retry.Sender.ID)
	data, _ := json.Marshal(retry)
	s.appService.Redis.RPush(*s.ctx.Ctx, key, data)
	retry.ID = utils.Uuid()
	s.ctx.DB.Create(&retry)
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
		key := fmt.Sprintf("retry:sender:%v", sender.ID)

		for {
			retryData, err := s.appService.Redis.LPop(*ctx, key).Result()
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

			s.sendWithRetryHandling(
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
		contacts, batches = contacts[size:], append(batches, contacts[:size:size])
	}

	return append(batches, contacts)
}

func (b *BroadcastService) sendWithRetryHandling(
	sender connection.ConnectionModel,
	broadcastID string,
	contact mdl.ContactModel,
	attempt int,
	delay time.Duration,
	logHandler func(log models.MessageLog),
	retryHandler func(retry models.MessageRetry),
) {
	// var broadcast models.BroadcastModel
	// b.ctx.DB.Where("id = ?", broadcastID).Preload("Member.User").First(&broadcast)

	broadcast, err := b.GetBroadcastByID(broadcastID)
	if err != nil {
		fmt.Println("ERROR GET BROADCAST", err)
		return
	}
	time.Sleep(delay)
	var success bool
	var isNotOnWhatsapp bool

	// USE SIMULATION
	if config.App.Server.SimulateBroadcast {
		if broadcast.TemplateID == nil {
			success = simulateSend(contact, parseMsgTemplate(contact, broadcast.Member, broadcast.Message))
		} else {
			// USE TEMPLATE
			var template mdl.WhatsappMessageTemplate
			template, err := b.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(*broadcast.TemplateID)
			if err == nil {
				for _, v := range template.Messages {
					success = simulateSend(contact, parseMsgTemplate(contact, broadcast.Member, v.Body))
					for _, file := range v.Files {
						fmt.Println("SIMULATE SEND FILE", file.URL)
					}
				}
			}

		}
	} else {

		fmt.Println("PREPARE TO SEND BROADCAST", contact.Name)

		// USE REAL API
		if contact.Phone != nil {
			fmt.Println("WITH PHONE NUMBER", *contact.Phone)
			resp, err := b.whatsmeowService.CheckNumber(sender.Session, *contact.Phone)
			if err != nil {
				log.Println("ERROR CHECK NUMBER", resp)
			}

			var respCheck QueryIsOnWhatsapp
			if err := json.Unmarshal(resp, &respCheck); err != nil {
				log.Println("ERROR CHECK NUMBER PARSE RESPONSE")
			}

			for _, v := range respCheck.Query {
				if !v.IsIn {
					isNotOnWhatsapp = true
				}
			}

			if isNotOnWhatsapp {
				log.Println("NUMBER IS NOT REGISTERED ON WHATSAPP")
			}
			msgData := mdl.WhatsappMessageModel{
				JID:     sender.Session,
				Message: parseMsgTemplate(contact, broadcast.Member, broadcast.Message),
			}

			fmt.Println("PRODUCTS BROADCAST")
			utils.LogJson(broadcast.Products)
			// USE REGULAR MESSAGE
			if broadcast.TemplateID == nil {
				success = true
				if sender.Type == "whatsapp-api" {
					var session *mdl.WhatsappMessageSession
					session.Contact = &contact
					err := b.appService.SendTemplateMessageWhatsappAPI(b.customerRelationshipService, b.metaService, &sender, msgData, session, broadcast.Member, broadcast.Files, broadcast.Products)
					// err := SendWhatsappApiContactMessage(sender, contact, msgData.Message, nil, broadcast.Files, broadcast.Products)
					if err != nil {
						log.Println("ERROR SEND MESSAGE REGULAR (WHATSAPP API)", err)
						success = false
					}
				} else {

					b.customerRelationshipService.WhatsappService.SetMsgData(b.whatsmeowService, &msgData, *contact.Phone, broadcast.Files, broadcast.Products, false, nil)
					_, err := customer_relationship.SendCustomerServiceMessage(b.customerRelationshipService.WhatsappService)
					if err != nil {
						log.Println("ERROR SEND MESSAGE REGULAR", err)
						success = false
					}
				}

				// err = b.customerRelationshipService.WhatsappService.SendWhatsappMessage(b.whatsmeowService, &msgData, *contact.Phone, broadcast.Files, broadcast.Products, false)
				// if err != nil {
				// 	log.Println("ERROR SEND MESSAGE REGULAR", err)
				// 	success = false
				// }
			} else {

				// USE TEMPLATE
				var template mdl.WhatsappMessageTemplate
				template, err := b.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(*broadcast.TemplateID)
				if err == nil {
					for _, v := range template.Messages {
						if sender.Type == "whatsapp-api" {
							parsedMsg := parseMsgTemplate(contact, broadcast.Member, v.Body)

							var session *mdl.WhatsappMessageSession = &mdl.WhatsappMessageSession{
								Contact: &contact,
							}
							msgData.Message = parsedMsg
							err := b.appService.SendTemplateMessageWhatsappAPI(b.customerRelationshipService, b.metaService, &sender, msgData, session, broadcast.Member, broadcast.Files, broadcast.Products)
							// err := SendWhatsappApiContactMessage(sender, contact, parsedMsg, nil, broadcast.Files)
							if err != nil {
								log.Println("ERROR SEND MESSAGE REGULAR (WHATSAPP API)", err)
								success = false
							}
						} else {
							success = true
							msgData := mdl.WhatsappMessageModel{
								JID:     sender.Session,
								Message: parseMsgTemplate(contact, broadcast.Member, v.Body),
							}
							b.customerRelationshipService.WhatsappService.SetMsgData(b.whatsmeowService, &msgData, *contact.Phone, v.Files, v.Products, false, nil)
							_, err := customer_relationship.SendCustomerServiceMessage(b.customerRelationshipService.WhatsappService)
							if err != nil {
								log.Println("ERROR SEND MESSAGE REGULAR", err)
								success = false
							}
						}

						// err = b.customerRelationshipService.WhatsappService.SendWhatsappMessage(b.whatsmeowService, &msgData, *contact.Phone, v.Files, v.Products, false)
						// if err != nil {
						// 	log.Println("ERROR SEND MESSAGE WITH TEMPLATE", err)
						// 	success = false
						// }
					}
				}

			}
			// 			if broadcast.TemplateID == nil {
			// 				resp, err := b.whatsmeowService.CheckNumber(sender.Session, *contact.Phone)
			// 				if err != nil {
			// 					log.Println("ERROR CHECK NUMBER", resp)
			// 				}

			// 				var respCheck QueryIsOnWhatsapp
			// 				if err := json.Unmarshal(resp, &respCheck); err != nil {
			// 					log.Println("ERROR CHECK NUMBER PARSE RESPONSE")
			// 				}

			// 				for _, v := range respCheck.Query {
			// 					if !v.IsIn {
			// 						isNotOnWhatsapp = true
			// 					}
			// 				}

			// 				thumbnail, restFiles := mdl.GetThumbnail(broadcast.Files)
			// 				var fileType, fileUrl string
			// 				if thumbnail != nil {
			// 					fileType = "image"
			// 					fileUrl = thumbnail.URL
			// 				}
			// 				waData := whatsmeow_client.WaMessage{
			// 					JID:      sender.Session,
			// 					Text:     parseMsgTemplate(contact, broadcast.Member, broadcast.Message),
			// 					To:       *contact.Phone,
			// 					IsGroup:  false,
			// 					FileType: fileType,
			// 					FileUrl:  fileUrl,
			// 				}
			// 				fmt.Println("SEND MESSAGE", *contact.Phone)
			// 				// utils.LogJson(waData)
			// 				if !isNotOnWhatsapp {
			// 					_, err = b.whatsmeowService.SendMessage(waData)
			// 					if err != nil {
			// 						success = false
			// 					} else {
			// 						success = true
			// 					}

			// 					for _, v := range restFiles {
			// 						if strings.Contains(v.MimeType, "image") && v.URL != "" {
			// 							resp, _ := b.whatsmeowService.SendMessage(whatsmeow_client.WaMessage{
			// 								JID:      sender.Session,
			// 								Text:     "",
			// 								To:       *contact.Phone,
			// 								IsGroup:  false,
			// 								FileType: "image",
			// 								FileUrl:  v.URL,
			// 							})
			// 							fmt.Println("RESPONSE", resp)
			// 						} else {
			// 							resp, _ := b.whatsmeowService.SendMessage(whatsmeow_client.WaMessage{
			// 								JID:      sender.Session,
			// 								Text:     "",
			// 								To:       *contact.Phone,
			// 								IsGroup:  false,
			// 								FileType: "document",
			// 								FileUrl:  v.URL,
			// 							})
			// 							fmt.Println("RESPONSE", resp)
			// 						}

			// 					}

			// 					b.ctx.DB.Preload("Products").Find(&broadcast)
			// 					for _, v := range broadcast.Products {
			// 						desc := ""
			// 						var images []mdl.FileModel
			// 						b.ctx.DB.Where("ref_id = ? and ref_type = ?", v.ID, "product").Find(&images)

			// 						if v.Description != nil {
			// 							desc = *v.Description
			// 						}
			// 						dataMsg := fmt.Sprintf(`*%s*
			// _%s_

			// %s
			// 		`, v.DisplayName, utils.FormatRupiah(v.Price), desc)
			// 						productMsg := whatsmeow_client.WaMessage{
			// 							JID:     sender.Session,
			// 							Text:    dataMsg,
			// 							To:      *contact.Phone,
			// 							IsGroup: false,
			// 						}

			// 						if len(images) > 0 {
			// 							productMsg.FileType = "image"
			// 							productMsg.FileUrl = images[0].URL
			// 						}
			// 						resp, _ := b.ctx.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(productMsg)
			// 						fmt.Println("RESPONSE", resp)

			// 					}
			// 				}

			// 			} else {
			// 				// USE TEMPLATE

			// 				var template mdl.WhatsappMessageTemplate
			// 				template, err := b.customerRelationshipService.WhatsappService.GetWhatsappMessageTemplate(*broadcast.TemplateID)
			// 				if err == nil {
			// 					for _, msg := range template.Messages {
			// 						fmt.Println("CHECK NUMBER")
			// 						resp, err := b.whatsmeowService.CheckNumber(sender.Session, *contact.Phone)
			// 						if err != nil {
			// 							log.Println("ERROR CHECK NUMBER", resp)
			// 						}

			// 						var respCheck QueryIsOnWhatsapp
			// 						if err := json.Unmarshal(resp, &respCheck); err != nil {
			// 							log.Println("ERROR CHECK NUMBER PARSE RESPONSE")
			// 							return
			// 						}
			// 						utils.LogJson(respCheck)
			// 						for _, v := range respCheck.Query {
			// 							if !v.IsIn {
			// 								isNotOnWhatsapp = true
			// 							}
			// 						}

			// 						if !isNotOnWhatsapp {
			// 							thumbnail, restFiles := mdl.GetThumbnail(msg.Files)
			// 							var fileType, fileUrl string
			// 							if thumbnail != nil {
			// 								fileType = "image"
			// 								fileUrl = thumbnail.URL
			// 							}
			// 							waData := whatsmeow_client.WaMessage{
			// 								JID:      sender.Session,
			// 								Text:     parseMsgTemplate(contact, broadcast.Member, msg.Body),
			// 								To:       *contact.Phone,
			// 								IsGroup:  false,
			// 								FileType: fileType,
			// 								FileUrl:  fileUrl,
			// 							}

			// 							_, err = b.whatsmeowService.SendMessage(waData)
			// 							if err != nil {
			// 								success = false
			// 							} else {
			// 								success = true
			// 							}

			// 							for _, v := range restFiles {
			// 								if strings.Contains(v.MimeType, "image") && v.URL != "" {
			// 									resp, _ := b.whatsmeowService.SendMessage(whatsmeow_client.WaMessage{
			// 										JID:      sender.Session,
			// 										Text:     "",
			// 										To:       *contact.Phone,
			// 										IsGroup:  false,
			// 										FileType: "image",
			// 										FileUrl:  v.URL,
			// 									})
			// 									fmt.Println("RESPONSE", resp)
			// 								} else {
			// 									resp, _ := b.whatsmeowService.SendMessage(whatsmeow_client.WaMessage{
			// 										JID:      sender.Session,
			// 										Text:     "",
			// 										To:       *contact.Phone,
			// 										IsGroup:  false,
			// 										FileType: "document",
			// 										FileUrl:  v.URL,
			// 									})
			// 									fmt.Println("RESPONSE", resp)
			// 								}

			// 							}
			// 							for _, v := range msg.Products {
			// 								desc := ""
			// 								var images []mdl.FileModel
			// 								b.ctx.DB.Where("ref_id = ? and ref_type = ?", v.ID, "product").Find(&images)

			// 								if v.Description != nil {
			// 									desc = *v.Description
			// 								}
			// 								dataMsg := fmt.Sprintf(`*%s*
			// _%s_

			// %s
			// 				`, v.DisplayName, utils.FormatRupiah(v.Price), desc)
			// 								productMsg := whatsmeow_client.WaMessage{
			// 									JID:     sender.Session,
			// 									Text:    dataMsg,
			// 									To:      *contact.Phone,
			// 									IsGroup: false,
			// 								}
			// 								if len(images) > 0 {
			// 									productMsg.FileType = "image"
			// 									productMsg.FileUrl = images[0].URL
			// 								}
			// 								resp, _ := b.ctx.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService).SendMessage(productMsg)
			// 								fmt.Println("RESPONSE", resp)

			// 							}
			// 						}

			// 					}
			// 				}

			// 			}

		}
	}

	msgLog := models.MessageLog{
		BroadcastID: broadcastID,
		ContactID:   contact.ID,
		SenderID:    sender.ID,
		SentAt:      time.Now(),
	}

	if success {
		msgLog.Status = "success"
		logHandler(msgLog)
		b.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("is_success", true)
		b.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("is_completed", true)
	} else {
		b.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("is_success", false)
		if isNotOnWhatsapp {
			msgLog.Status = "failed"
			msgLog.ErrorMessage = "number is not registered on whatsapp"
			logHandler(msgLog)
			b.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("is_completed", true)
		} else if attempt >= 3 {
			msgLog.Status = "undeliverable"
			msgLog.ErrorMessage = fmt.Sprintf("attempt %d failed", attempt)
			logHandler(msgLog)
			b.ctx.DB.Model(&models.BroadcastContacts{}).Where("contact_model_id = ? and broadcast_model_id = ?", contact.ID, broadcastID).Update("is_completed", true)
		} else {
			msgLog.Status = "failed"
			msgLog.ErrorMessage = fmt.Sprintf("attempt %d failed", attempt)
			logHandler(msgLog)

			retryHandler(models.MessageRetry{
				BroadcastID: broadcastID,
				Contact:     contact,
				Sender:      sender,
				Attempt:     attempt + 1,
				LastError:   msgLog.ErrorMessage,
				LastTriedAt: time.Now(),
			})
		}
	}

	var completedCount int64
	b.ctx.DB.Model(&models.BroadcastContacts{}).Where(" broadcast_model_id = ? and is_completed = ?", broadcastID, true).Count(&completedCount)
	var contactCount int64
	b.ctx.DB.Model(&models.BroadcastContacts{}).Where(" broadcast_model_id = ?", broadcastID).Count(&contactCount)
	fmt.Println("COUNT", completedCount, contactCount)
	if completedCount == contactCount {
		b.ctx.DB.Where("id = ?", broadcastID).Model(&broadcast).Update("status", "COMPLETED")
		msg := map[string]any{
			"message":      "All contacts have been sent",
			"broadcast_id": broadcastID,
			"command":      "BROADCAST_COMPLETED",
		}
		msgStr, _ := json.Marshal(msg)
		b.appService.Websocket.BroadcastFilter(msgStr, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", b.appService.Config.Server.BaseURL, *broadcast.CompanyID)
			return fmt.Sprintf("%s%s", b.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	} else {
		type count struct {
			Complete int64 `json:"complete"`
			Success  int64 `json:"success"`
			Failed   int64 `json:"failed"`
		}
		var countData count
		b.ctx.DB.Model(&models.BroadcastContacts{}).Where("broadcast_model_id = ?", broadcastID).Select("COUNT(CASE WHEN is_completed = 't' THEN 1 END) as complete, COUNT(CASE WHEN is_success = 't' THEN 1 END) as success, COUNT(CASE WHEN is_success = 'f' THEN 1 END) as failed").Scan(&countData)
		msg := map[string]any{
			"message":      "Broadcast in progress",
			"broadcast_id": broadcastID,
			"command":      "BROADCAST_PROGRESS",
			"data": map[string]any{
				"success":   countData.Success,
				"failed":    countData.Failed,
				"completed": countData.Complete,
			},
		}
		msgStr, _ := json.Marshal(msg)
		b.appService.Websocket.BroadcastFilter(msgStr, func(q *melody.Session) bool {
			url := fmt.Sprintf("%s/api/v1/ws/%s", b.appService.Config.Server.BaseURL, *broadcast.CompanyID)
			return fmt.Sprintf("%s%s", b.appService.Config.Server.BaseURL, q.Request.URL.Path) == url
		})
	}

}

func simulateSend(contact mdl.ContactModel, msg string) bool {

	fmt.Println("Simulate send to", contact.Name, *contact.Phone, "with message \n", msg)
	// 90% berhasil
	return rand.Intn(100) < 90
}

func (s BroadcastService) logHandler(log models.MessageLog) {
	log.ID = utils.Uuid()
	s.ctx.DB.Create(&log)
}

func parseMsgTemplate(contact mdl.ContactModel, member *mdl.MemberModel, msg string) string {
	re := regexp.MustCompile(`@\[([^\]]+)\]|\(\{\{([^}]+)\}\}\)`)

	// Replace
	result := re.ReplaceAllStringFunc(msg, func(s string) string {
		matches := re.FindStringSubmatch(s)

		fmt.Println("MATCHES", matches)
		re2 := regexp.MustCompile(`@\[([^\]]+)\]`)
		if re2.MatchString(s) {
			return ""
		}

		if matches[0] == "({{user}})" {
			return contact.Name
		}
		if matches[0] == "({{phone}})" {
			return *contact.Phone
		}

		if matches[0] == "({{agent}})" && member != nil {
			return member.User.FullName
		}
		if matches[0] == "({{product}})" {
			var customData map[string]string
			json.Unmarshal([]byte(contact.CustomData), &customData)
			return customData["product"]
		}
		return s // Kalau tidak ada datanya, biarkan
	})

	return result
}

type QueryJID struct {
	Query        string `json:"Query"`
	JID          string `json:"JID"`
	IsIn         bool   `json:"IsIn"`
	VerifiedName any    `json:"VerifiedName"`
}

type QueryIsOnWhatsapp struct {
	Query   []QueryJID `json:"is_on_whatsapp"`
	Message string     `json:"message"`
}
