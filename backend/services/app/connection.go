package app

import (
	"ametory-pm/models/connection"
	srv "ametory-pm/services"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"gorm.io/gorm"
)

type ConnectionService struct {
	ctx *context.ERPContext
}

func NewConnectionService(erpContext *context.ERPContext) *ConnectionService {
	return &ConnectionService{
		ctx: erpContext,
	}
}

func (c *ConnectionService) GetConnections(pagination *Pagination, httpRequest http.Request, search string) ([]connection.ConnectionModel, error) {
	var connections []connection.ConnectionModel

	if err := c.ctx.DB.Scopes(paginate(connections, pagination, c.ctx.DB)).Preload("GeminiAgent").Find(&connections).Error; err != nil {
		return nil, err
	}
	return connections, nil
}

func (c *ConnectionService) GetConnection(id string) (*connection.ConnectionModel, error) {
	var con connection.ConnectionModel
	if err := c.ctx.DB.Where("id = ?", id).Preload("GeminiAgent").First(&con).Error; err != nil {
		return nil, err
	}
	return &con, nil
}

func (c *ConnectionService) GetConnectionBySession(session string) (*connection.ConnectionModel, error) {
	var con connection.ConnectionModel
	if err := c.ctx.DB.Where("session_name = ?", session).Preload("GeminiAgent").First(&con).Error; err != nil {
		return nil, err
	}
	return &con, nil
}

func (c *ConnectionService) CreateConnection(con *connection.ConnectionModel) error {
	if err := c.ctx.DB.Create(con).Error; err != nil {
		return err
	}
	return nil
}

func (c *ConnectionService) UpdateConnection(con *connection.ConnectionModel) error {
	if err := c.ctx.DB.Save(con).Error; err != nil {
		return err
	}
	return nil
}

func (c *ConnectionService) DeleteConnection(id string) error {
	if err := c.ctx.DB.Where("id = ?", id).Delete(&connection.ConnectionModel{}).Error; err != nil {
		return err
	}

	return nil
}

func (c *ConnectionService) GetActiveSession(phoneNumber string) (*models.ContactModel, error) {
	sessionStr, err := srv.REDIS.Get(*c.ctx.Ctx, fmt.Sprintf("session:%s", phoneNumber)).Result()
	if err != nil {
		return nil, err
	}
	session := models.ContactModel{}
	err = json.Unmarshal([]byte(sessionStr), &session)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (c *ConnectionService) RemoveSession(phoneNumber string) error {
	return srv.REDIS.Del(*c.ctx.Ctx, fmt.Sprintf("session:%s", phoneNumber)).Err()
}
func (c *ConnectionService) SetSession(phoneNumber string) error {
	var member *models.ContactModel
	err := c.ctx.DB.Model(&models.ContactModel{}).Where("phone = ?", phoneNumber).First(&member).Error
	if err != nil {
		return err
	}
	sessionJson, err := json.Marshal(member)
	if err != nil {
		return err
	}
	err = srv.REDIS.Set(*c.ctx.Ctx, fmt.Sprintf("session:%s", phoneNumber), string(sessionJson), 7*24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c *ConnectionService) IsPhoneNumberRegistered(phoneNumber string) bool {
	var member models.ContactModel
	err := c.ctx.DB.Model(&member).Where("phone = ?", phoneNumber).First(&member).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func paginate(value any, pagination *Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)
	pagination.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.Limit)))
	pagination.TotalPages = totalPages
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order(pagination.GetSort())
	}
}

type Pagination struct {
	Limit      int    `json:"limit,omitempty" query:"limit"`
	Page       int    `json:"page,omitempty" query:"page"`
	Sort       string `json:"sort,omitempty" query:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
	Rows       any    `json:"rows"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}
func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "id desc"
	}
	return p.Sort
}
