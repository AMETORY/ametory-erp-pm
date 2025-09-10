package worker

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"fmt"
	"log"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
)

func GetStoppedBroadcasts(erpContext *context.ERPContext) ([]models.BroadcastModel, error) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "START GET STOPPED BROADCASTS")
	var broadcasts []models.BroadcastModel
	if err := erpContext.DB.Where("status IN (?)", []string{"PROCESSING"}).Find(&broadcasts).Error; err != nil {
		return nil, err
	}

	broadcastSrv, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	if ok {
		for _, v := range broadcasts {
			if v.Status == "STOPPED" || v.Status == "NOT_STARTED" {
				v.Status = "RESTARTING"
				broadcast, err := broadcastSrv.GetBroadcastByID(v.ID)
				if err != nil {
					log.Println("ERROR", err)
					continue
				}
				erpContext.DB.Save(broadcast)
				broadcastSrv.StartBroadcast(broadcast)
			}
		}
	}

	return broadcasts, nil
}
