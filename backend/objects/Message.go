package objects

import (
	"time"

	"github.com/AMETORY/ametory-erp-modules/shared/models"
)

type ScheduledMessage struct {
	To       string                      `json:"to"`
	Message  string                      `json:"message"`
	Duration time.Duration               `json:"duration"`
	Files    []models.FileModel          `json:"files"`
	Data     models.WhatsappMessageModel `json:"data"`
	Action   *models.ColumnAction        `json:"action"`
	Task     *models.TaskModel           `json:"task"`
}
