package worker

import (
	"ametory-pm/models"
	"ametory-pm/objects"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/customer_relationship"
	mdl "github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/whatsmeow_client"
)

func ScheduledBroadcastWorker(erpContext *context.ERPContext) {
	fmt.Println("START SCHEDULED BROADCAST WORKER")

	broadcastSrv, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	appService, ok2 := erpContext.AppService.(*app.AppService)
	if ok && ok2 {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "BROADCAST:SCHEDULED")
		for {
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			var broadcastData models.BroadcastModel
			err = json.Unmarshal([]byte(msg.Payload), &broadcastData)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("BROADCAST SCHEDULED", broadcastData.Description, broadcastData.ScheduledAt.Format("2006-01-02 15:04:05"))
			// go func() {
			time.Sleep(time.Until(*broadcastData.ScheduledAt))

			err = erpContext.DB.First(&broadcastData, "id = ?", broadcastData.ID).Error
			if err != nil {
				log.Println("ERROR", err)
				return
			}
			broadcastData.Status = "PROCESSING"
			erpContext.DB.Save(&broadcastData)
			broadcastSrv.StartBroadcast(&broadcastData, false)
			// }()

		}
	}

}
func BroadcastWorker(erpContext *context.ERPContext) {
	fmt.Println("START BROADCAST WORKER NOW")

	broadcastSrv, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	appService, ok2 := erpContext.AppService.(*app.AppService)
	if ok && ok2 {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "BROADCAST:NOW")
		for {
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
				continue
			}
			var broadcastData models.BroadcastModel
			err = json.Unmarshal([]byte(msg.Payload), &broadcastData)
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println("BROADCAST NOW", broadcastData.Description, time.Now().Format("2006-01-02 15:04:05"))
			go func() {
				broadcastData.Status = "PROCESSING"
				err := erpContext.DB.First(&broadcastData, "id = ?", broadcastData.ID).Error
				if err != nil {
					log.Println("ERROR", err)
					return
				}
				erpContext.DB.Save(&broadcastData)
				broadcastSrv.StartBroadcast(&broadcastData, false)
			}()

		}
	}

}

func ScheduledMessageWorker(erpContext *context.ERPContext) {
	fmt.Println("START SCHEDULED MESSAGE WORKER")
	customerRelationshipService, ok := erpContext.CustomerRelationshipService.(*customer_relationship.CustomerRelationshipService)
	appService, ok2 := erpContext.AppService.(*app.AppService)
	if ok && ok2 {
		dataSub := appService.Redis.Subscribe(*erpContext.Ctx, "MESSAGE:SCHEDULED")
		for {
			msg, err := dataSub.ReceiveMessage(*erpContext.Ctx)
			if err != nil {
				log.Println(err)
				continue
			}

			var msgData objects.ScheduledMessage
			err = json.Unmarshal([]byte(msg.Payload), &msgData)
			if err != nil {
				log.Println(err)
				continue
			}
			time.Sleep(msgData.Duration)

			log.Println("MESSAGE SCHEDULED", msgData.Message, msgData.Duration)
			whatsmeowService, ok := erpContext.ThirdPartyServices["WA"].(*whatsmeow_client.WhatsmeowService)
			if !ok {
				log.Println("ThirdPartyServices is not instance of whatsmeow_client.WhatsmeowService")
				continue
			}
			customerRelationshipService.WhatsappService.SetMsgData(whatsmeowService, &msgData.Data, msgData.To, msgData.Files, []mdl.ProductModel{}, false, nil)
			_, err = customer_relationship.SendCustomerServiceMessage(customerRelationshipService.WhatsappService)
			if err != nil {
				log.Println("ERROR", err)
				continue
			}
			if msgData.Action != nil {
				err = erpContext.DB.Model(&mdl.ColumnAction{}).Where("id = ?", msgData.Action.ID).Update("action_status", "READY").Error
				if err != nil {
					log.Println("ERROR", err)
					continue
				}
			}

			if msgData.Task != nil {
				if !msgData.Action.RunOnce {
					err = erpContext.DB.Model(&mdl.TaskModel{}).Where("id = ?", msgData.Task.ID).Update("last_action_trigger_at", time.Now()).Error
					if err != nil {
						log.Println("ERROR", err)
						continue
					}
				}

				erpContext.DB.Model(&msgData.Action).Association("ColumnActions").Append(&msgData.Action)
			}

		}
	}
}
