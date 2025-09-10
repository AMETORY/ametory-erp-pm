package worker

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"fmt"
	"log"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"gorm.io/gorm/clause"
)

func GetStoppedBroadcasts(erpContext *context.ERPContext) ([]models.BroadcastModel, error) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "START GET STOPPED BROADCASTS")
	var broadcasts []models.BroadcastModel
	if err := erpContext.DB.Where("status IN (?)", []string{"PROCESSING", "RESTARTING", "STOPPED"}).Find(&broadcasts).Error; err != nil {
		return nil, err
	}

	broadcastSrv, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	if ok {
		for _, v := range broadcasts {
			if v.Status == "STOPPED" || v.Status == "NOT_STARTED" {

				erpContext.DB.Preload(clause.Associations).Find(&v)
				log.Println("START RESTARTING BROADCAST", v.ID, v.Description)
				v.Status = "RESTARTING"
				erpContext.DB.Save(v)
				broadcastSrv.StartBroadcast(&v, true)
			}
		}
	}

	return broadcasts, nil
}
