package handlers

import (
	"math"
	"strings"
	"time"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

type AnalyticHandler struct {
	ctx *context.ERPContext
}

func NewAnalyticHandler(ctx *context.ERPContext) *AnalyticHandler {
	return &AnalyticHandler{
		ctx: ctx,
	}
}

func (h *AnalyticHandler) CustomerInteractionHandler(c *gin.Context) {
	var input struct {
		StartDate string   `json:"start_date" binding:"required"`
		EndDate   string   `json:"end_date" binding:"required"`
		MemberIDs []string `json:"member_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var messages models.WhatsappMessageModel
	var count struct {
		NewCustomerCount      int64 `gorm:"column:new_customer_count" json:"new_customer_count"`
		ExistingCustomerCount int64 `gorm:"column:existing_customer_count" json:"existing_customer_count"`
	}

	db := h.ctx.DB.
		Model(&messages).
		Select("SUM(CASE WHEN is_new = true THEN 1 ELSE 0 END) as new_customer_count, SUM(CASE WHEN is_new = false THEN 1 ELSE 0 END) as existing_customer_count")

	db = db.Where("created_at >= ? AND created_at <= ?", input.StartDate, input.EndDate)
	db = db.Where("is_from_me = ?", false)
	// if len(input.MemberIDs) > 0 {
	// 	db = db.Where("member_id IN ?", input.MemberIDs)
	// }

	companyID := c.MustGet("companyID").(string)
	err := db.Where("company_id = ?", companyID).Group("company_id").Scan(&count).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Customer Interaction retrieved successfully", "data": map[string]float64{"existing": float64(count.ExistingCustomerCount), "new": float64(count.NewCustomerCount)}})
}

func (h *AnalyticHandler) AverageTimeReplyHandler(c *gin.Context) {
	var input struct {
		StartDate string   `json:"start_date" binding:"required"`
		EndDate   string   `json:"end_date" binding:"required"`
		MemberIDs []string `json:"member_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var messages models.WhatsappMessageModel
	var count []struct {
		AvgTimeReplyNew      float64 `gorm:"column:avg_time_reply_new" json:"avg_time_reply_new"`
		AvgTimeReplyExisting float64 `gorm:"column:avg_time_reply_existing" json:"avg_time_reply_existing"`
	}

	db := h.ctx.DB.
		Model(&messages).
		Select("whatsapp_messages.company_id, whatsapp_messages.ref_id, AVG(CASE WHEN ref.is_new = true THEN whatsapp_messages.response_time ELSE 0 END) as avg_time_reply_new, AVG(CASE WHEN ref.is_new = false THEN whatsapp_messages.response_time ELSE 0 END) as avg_time_reply_existing").
		Joins("JOIN whatsapp_messages ref ON whatsapp_messages.ref_id = ref.id").
		Where("whatsapp_messages.created_at >= ? AND whatsapp_messages.created_at <= ?", input.StartDate, input.EndDate).
		Where("whatsapp_messages.response_time IS NOT NULL").
		Where("whatsapp_messages.company_id = ?", c.MustGet("companyID").(string))

	if len(input.MemberIDs) > 0 {
		db = db.Where("whatsapp_messages.member_id IN ?", input.MemberIDs)
	}

	err := db.Group("whatsapp_messages.company_id, whatsapp_messages.ref_id").Scan(&count).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existing, new, countExisting, countNew float64
	for _, v := range count {
		existing += v.AvgTimeReplyExisting
		new += v.AvgTimeReplyNew
		if v.AvgTimeReplyExisting > 0 {
			countExisting++
		}
		if v.AvgTimeReplyNew > 0 {
			countNew++
		}
	}

	if len(count) == 0 {
		c.JSON(200, gin.H{"message": "Average Time Reply Customer Interaction retrieved successfully", "data": map[string]float64{"existing": 0, "new": 0}})
		return
	}

	if countExisting > 0 {
		existing /= countExisting
	}

	if countNew > 0 {
		new /= countNew
	}

	c.JSON(200, gin.H{"message": "Average Time Reply Customer Interaction retrieved successfully", "data": map[string]float64{"existing": existing, "new": new}})
}

func (h *AnalyticHandler) HourlyCustomerInteractionHandler(c *gin.Context) {
	var input struct {
		StartDate string   `json:"start_date" binding:"required"`
		EndDate   string   `json:"end_date" binding:"required"`
		MemberIDs []string `json:"member_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var messages models.WhatsappMessageModel
	var count []struct {
		Hour                  string `gorm:"column:hour" json:"hour"`
		NewCustomerCount      int64  `gorm:"column:new_customer_count" json:"new_customer_count"`
		ExistingCustomerCount int64  `gorm:"column:existing_customer_count" json:"existing_customer_count"`
	}

	db := h.ctx.DB.
		Model(&messages).
		Select("TO_CHAR(created_at, 'HH24:00') AS hour, SUM(CASE WHEN is_new = true THEN 1 ELSE 0 END) as new_customer_count, SUM(CASE WHEN is_new = false THEN 1 ELSE 0 END) as existing_customer_count").
		Where("created_at >= ? AND created_at <= ?", input.StartDate, input.EndDate).
		Where("company_id = ?", c.MustGet("companyID").(string)).
		Group("hour").
		Order("hour ASC")

	db = db.Where("is_from_me = ?", false)

	// if len(input.MemberIDs) > 0 {
	// 	db = db.Where("member_id IN ?", input.MemberIDs)
	// }

	err := db.Group("company_id, ref_id").Scan(&count).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hourlyMap := make(map[string][]int64)
	for _, r := range count {
		hourlyMap[r.Hour] = []int64{r.NewCustomerCount, r.ExistingCustomerCount}
	}

	loc, _ := time.LoadLocation("Local")
	start := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)

	newHourlyMap := []map[string]any{}
	for i := 0; i < 24; i++ {
		current := start.Add(time.Duration(i) * time.Hour)
		hourStr := current.Format("15:04")
		total, ok := hourlyMap[hourStr] // jika tidak ada, default = 0
		if ok {
			newHourlyMap = append(newHourlyMap, map[string]any{"hour": strings.Split(hourStr, ":")[0], "new": total[0], "existing": total[1]})
		} else {
			newHourlyMap = append(newHourlyMap, map[string]any{"hour": strings.Split(hourStr, ":")[0], "new": 0, "existing": 0})
		}
	}

	c.JSON(200, gin.H{"message": "Customer Interaction retrieved successfully", "data": newHourlyMap})
}

func (h *AnalyticHandler) HourlyAverageTimeReplyHandler(c *gin.Context) {
	var input struct {
		StartDate string   `json:"start_date" binding:"required"`
		EndDate   string   `json:"end_date" binding:"required"`
		MemberIDs []string `json:"member_ids"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var messages models.WhatsappMessageModel
	var count []struct {
		Hour                 string  `gorm:"column:hour" json:"hour"`
		AvgTimeReplyNew      float64 `gorm:"column:avg_time_reply_new" json:"avg_time_reply_new"`
		AvgTimeReplyExisting float64 `gorm:"column:avg_time_reply_existing" json:"avg_time_reply_existing"`
	}

	db := h.ctx.DB.
		Model(&messages).
		Select("TO_CHAR(created_at, 'HH24:00') AS hour, AVG(CASE WHEN is_new = true THEN response_time ELSE 0 END) as avg_time_reply_new, AVG(CASE WHEN is_new = false THEN response_time ELSE 0 END) as avg_time_reply_existing").
		Where("created_at >= ? AND created_at <= ?", input.StartDate, input.EndDate).
		Where("company_id = ?", c.MustGet("companyID").(string)).
		Where("response_time IS NOT NULL").
		Group("hour").
		Order("hour ASC")

	db = db.Where("is_from_me = ?", true)

	if len(input.MemberIDs) > 0 {
		db = db.Where("member_id IN ?", input.MemberIDs)
	}

	err := db.Group("company_id, ref_id").Scan(&count).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hourlyMap := make(map[string][]float64)
	for _, r := range count {
		hourlyMap[r.Hour] = []float64{r.AvgTimeReplyNew, r.AvgTimeReplyExisting}
	}

	loc, _ := time.LoadLocation("Local")
	start := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, loc)

	newHourlyMap := []map[string]any{}
	for i := 0; i < 24; i++ {
		current := start.Add(time.Duration(i) * time.Hour)
		hourStr := current.Format("15:04")
		total, ok := hourlyMap[hourStr] // jika tidak ada, default = 0

		if ok {
			newHourlyMap = append(newHourlyMap, map[string]any{
				"hour":     strings.Split(hourStr, ":")[0],
				"new":      math.Round(total[0]),
				"existing": math.Round(total[1]),
			})

		} else {
			newHourlyMap = append(newHourlyMap, map[string]any{"hour": strings.Split(hourStr, ":")[0], "new": 0, "existing": 0})
		}
	}

	c.JSON(200, gin.H{"message": "Customer Interaction retrieved successfully", "data": newHourlyMap})
}
