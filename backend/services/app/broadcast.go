package app

import (
	"ametory-pm/models"

	"github.com/AMETORY/ametory-erp-modules/context"
)

type BroadcastService struct {
	ctx *context.ERPContext
}

func NewBroadcastService(ctx *context.ERPContext) *BroadcastService {
	if !ctx.SkipMigration {
		ctx.DB.AutoMigrate(&models.BroadcastModel{}, &models.BroadcastGrouping{}, &models.BroadcastContacts{})
	}
	return &BroadcastService{
		ctx: ctx,
	}
}

func (s *BroadcastService) CreateBroadcast(broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Create(broadcast).Error
}

func (s *BroadcastService) GetBroadcasts(companyID string) ([]models.BroadcastModel, error) {
	var broadcasts []models.BroadcastModel
	if err := s.ctx.DB.Where("company_id = ?", companyID).Find(&broadcasts).Error; err != nil {
		return nil, err
	}
	return broadcasts, nil
}

func (s *BroadcastService) GetBroadcastByID(id string) (*models.BroadcastModel, error) {
	var broadcast models.BroadcastModel
	if err := s.ctx.DB.Where("id = ?", id).First(&broadcast).Error; err != nil {
		return nil, err
	}
	return &broadcast, nil
}

func (s *BroadcastService) UpdateBroadcast(id string, broadcast *models.BroadcastModel) error {
	return s.ctx.DB.Where("id = ?", id).Save(broadcast).Error
}

func (s *BroadcastService) DeleteBroadcast(id string) error {
	return s.ctx.DB.Delete(&models.BroadcastModel{}, "id = ?", id).Error
}
