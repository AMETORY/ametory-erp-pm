package worker

import (
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/project_management"
)

func CheckIdleColumn(erpContext *context.ERPContext) {
	pmService := erpContext.ProjectManagementService.(*project_management.ProjectManagementService)

	pmService.ProjectService.CheckIdleColumn()
}
