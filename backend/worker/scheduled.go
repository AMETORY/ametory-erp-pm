package worker

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
)

func ScheduledBroadcastWorker(erpContext *context.ERPContext) {
	fmt.Println("START SCHEDULED WORKER")

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
			go func() {
				time.Sleep(time.Until(*broadcastData.ScheduledAt))
				broadcastData.Status = "PROCESSING"
				err := erpContext.DB.First(&broadcastData, "id = ?", broadcastData.ID).Error
				if err != nil {
					log.Println("ERROR", err)
					return
				}
				erpContext.DB.Save(&broadcastData)
				broadcastSrv.StartBroadcast(&broadcastData)
			}()

		}
	}

}
