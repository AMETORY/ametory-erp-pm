package worker

import (
	"log"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
)

func CheckIdleColumn(erpContext *context.ERPContext) {
	log.Println("START CHECK IDLE COLUMN")
	pmService := erpContext.ProjectManagementService.(*project_management.ProjectManagementService)

	pmService.ProjectService.CheckIdleColumn()
}
