package worker

import (
	"ametory-pm/services/app"
	"encoding/json"
	"log"

	"github.com/AMETORY/ametory-erp-modules/shared/models"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
)

func CheckIdleColumn(erpContext *context.ERPContext) {
	log.Println("START CHECK IDLE COLUMN")
	pmService := erpContext.ProjectManagementService.(*project_management.ProjectManagementService)
	pmService.ProjectService.CheckIdleColumn(func(sm models.ScheduledMessage) {
		log.Println("IDLE COLUMN FOUND, SENDING SCHEDULED MESSAGE")
		runScheduledMessages(erpContext, sm)
	})
}

func runScheduledMessages(erpContext *context.ERPContext, dataSchedule models.ScheduledMessage) {
	appService, ok := erpContext.AppService.(*app.AppService)
	if !ok {
		log.Println("Failed to cast AppService")
		return
	}
	b, err := json.Marshal(dataSchedule)
	if err != nil {
		log.Println("Error marshalling scheduled message:", err)
		return
	}
	log.Println("SENDING SCHEDULED MESSAGE", string(b))
	appService.Redis.Publish(*erpContext.Ctx, "MESSAGE:SCHEDULED", b)
}
