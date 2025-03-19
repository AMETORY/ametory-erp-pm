package handlers

import (
	com "ametory-pm/models/company"
	"encoding/json"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/google"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type GeminiHandler struct {
	ctx           *context.ERPContext
	geminiService *google.GeminiService
}

func NewGeminiHandler(ctx *context.ERPContext) *GeminiHandler {
	geminiService, ok := ctx.ThirdPartyServices["GEMINI"].(*google.GeminiService)
	if !ok {
		panic("GeminiService is not found")
	}

	return &GeminiHandler{
		ctx:           ctx,
		geminiService: geminiService,
	}
}

func (h *GeminiHandler) GenerateContentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)
	var agentID *string
	agentId := c.Query("agent_id")
	var input struct {
		Content string
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var companySetting com.CompanySetting
	err := h.ctx.DB.Find(&companySetting, "company_id = ?", member.CompanyID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Company setting is not found"})
		return
	}

	var histories []models.GeminiHistoryModel

	skipHistory := c.Query("skip_history")
	if skipHistory == "false" {
		err = h.ctx.DB.Model(&models.GeminiHistoryModel{}).Find(&histories, "agent_id = ? and is_model = ?", agentId, true).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Agent histories is not found"})
			return
		}
	}

	h.geminiService.SetupAPIKey(companySetting.GeminiAPIKey, c.Query("skip_history") == "true")

	if agentId != "" {
		var agent models.GeminiAgent
		err := h.ctx.DB.Find(&agent, "id = ?", agentId).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Agent is not found"})
			return
		}
		h.geminiService.SetupModel(agent.SetTemperature, agent.SetTopK, agent.SetTopP, agent.SetMaxOutputTokens, agent.ResponseMimetype, agent.Model)
		h.geminiService.SetUpSystemInstruction(agent.SystemInstruction)
		agentID = &agentId
	}

	chatHistories := []map[string]any{}
	for _, v := range histories {
		chatHistories = append(chatHistories, map[string]any{
			"role":    "user",
			"content": v.Input,
		})
		chatHistories = append(chatHistories, map[string]any{
			"role":    "model",
			"content": v.Output,
		})
	}

	// h.geminiService.SetupModel(companySetting.GeminiAPIKey)
	utils.LogJson(chatHistories)
	output, err := h.geminiService.GenerateContent(*h.ctx.Ctx, input.Content, chatHistories, "", "")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	outputResp := map[string]any{}
	json.Unmarshal([]byte(output), &outputResp)

	if c.Query("skip_save") == "true" {

	} else {
		var history models.GeminiHistoryModel = models.GeminiHistoryModel{
			Input:   input.Content,
			Output:  output,
			AgentID: agentID,
		}

		err = h.ctx.DB.Create(&history).Error
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(200, gin.H{"data": outputResp})
}

func (h *GeminiHandler) GetAgentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)
	var companySetting com.CompanySetting
	err := h.ctx.DB.Find(&companySetting, "company_id = ?", member.CompanyID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Company setting is not found"})
		return
	}

	// h.geminiService.SetupAPIKey(companySetting.GeminiAPIKey)
	agents, err := h.geminiService.GetAgents(*c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": agents})
}

func (h *GeminiHandler) GetAgentDetailHandler(c *gin.Context) {
	id := c.Param("id")

	agent, err := h.geminiService.GetAgent(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": agent})

}
func (h *GeminiHandler) GetAgentHistoriesHandler(c *gin.Context) {
	id := c.Param("id")

	// agent, err := h.geminiService.GetAgent(id)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	var histories []models.GeminiHistoryModel
	err := h.ctx.DB.Model(&models.GeminiHistoryModel{}).Find(&histories, "agent_id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Agent histories is not found"})
		return
	}
	c.JSON(200, gin.H{"data": histories})

}

func (h *GeminiHandler) DeleteHistoryHandler(c *gin.Context) {
	historyId := c.Param("historyId")

	err := h.ctx.DB.Delete(&models.GeminiHistoryModel{}, "id = ?", historyId).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "History deleted successfully"})
}

func (h *GeminiHandler) UpdateHistoryHandler(c *gin.Context) {
	// id := c.Param("id")
	historyId := c.Param("historyId")

	var input models.GeminiHistoryModel
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.ctx.DB.Model(&models.GeminiHistoryModel{}).Where("id = ?", historyId).Updates(map[string]any{
		"input":    input.Input,
		"output":   input.Output,
		"agent_id": input.AgentID,
	}).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "History updated successfully"})
}

func (h *GeminiHandler) ToggleModelHistoryHandler(c *gin.Context) {
	historyId := c.Param("historyId")

	var history models.GeminiHistoryModel
	err := h.ctx.DB.Find(&history, "id = ?", historyId).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "History not found"})
		return
	}

	history.IsModel = !history.IsModel
	err = h.ctx.DB.Save(&history).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "History published updated successfully"})
}
func (h *GeminiHandler) CreateAgentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)

	var companySetting com.CompanySetting
	err := h.ctx.DB.Find(&companySetting, "company_id = ?", member.CompanyID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Company setting is not found"})
		return
	}

	var input models.GeminiAgent
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if input.ApiKey == "" {
		input.ApiKey = companySetting.GeminiAPIKey
	}
	err = h.geminiService.CreateAgent(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent created successfully"})
}

func (h *GeminiHandler) UpdateAgentHandler(c *gin.Context) {
	var input models.GeminiAgent
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")

	err := h.geminiService.UpdateAgent(id, &input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent updated successfully"})
}

func (h *GeminiHandler) DeleteAgentHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.geminiService.DeleteAgent(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent deleted successfully"})
}
