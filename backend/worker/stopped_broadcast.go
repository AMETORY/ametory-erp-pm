package worker

import (
	"ametory-pm/models"
	"ametory-pm/services/app"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"gorm.io/gorm/clause"
)

func GetStoppedBroadcasts(erpContext *context.ERPContext) ([]models.BroadcastModel, error) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "START GET STOPPED BROADCASTS")
	var broadcasts []models.BroadcastModel
	if err := erpContext.DB.Where("status IN (?)", []string{"STOPPED"}).Find(&broadcasts).Error; err != nil {
		return nil, err
	}

	broadcastSrv, ok := erpContext.ThirdPartyServices["BROADCAST"].(*app.BroadcastService)
	if ok {
		var wg sync.WaitGroup // 1. Inisialisasi WaitGroup
		for _, v := range broadcasts {
			broadcast := v // Buat salinan dari v untuk setiap iterasi
			wg.Add(1)      // 2. Tambah hitungan untuk setiap goroutine
			go func(b models.BroadcastModel) {
				defer wg.Done() // 3. Pastikan Done() dipanggil saat goroutine selesai
				erpContext.DB.Preload(clause.Associations).Find(&b)
				log.Println("START RESTARTING BROADCAST", b.ID, b.Description)
				b.Status = "PROCESSING"
				erpContext.DB.Save(&b)
				broadcastSrv.StartBroadcast(&b, true)
			}(broadcast) // Lewatkan salinan ke goroutine
		}
		wg.Wait() // 4. Tunggu semua goroutine selesai
		log.Println(time.Now().Format("2006-01-02 15:04:05"), "ALL STOPPED BROADCASTS HAVE BEEN RESTARTED")
	}

	return broadcasts, nil
}
