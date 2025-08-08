package handlers

import (
	com "ametory-pm/models/company"
	"fmt"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/thirdparty/ai_generator"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/gin-gonic/gin"
)

type AiAgentHandler struct {
	ctx                *context.ERPContext
	aiGeneratorService *ai_generator.AiGeneratorService
}

func NewAiAgentHandler(ctx *context.ERPContext) *AiAgentHandler {

	aiGeneratorService, ok := ctx.ThirdPartyServices["AiGenerator"].(*ai_generator.AiGeneratorService)
	if !ok {
		panic("aiGeneratorService is not instance of cache.CacheManager")
	}

	return &AiAgentHandler{
		ctx:                ctx,
		aiGeneratorService: aiGeneratorService,
	}
}

func (h *AiAgentHandler) GenerateContentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)
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

	var histories []models.AiAgentHistory

	skipHistory := c.Query("skip_history")
	if skipHistory == "false" {
		err = h.ctx.DB.Model(&models.AiAgentHistory{}).Find(&histories, "ai_agent_id = ? and is_model = ?", agentId, true).Error
		if err != nil {
			c.JSON(404, gin.H{"error": "Agent histories is not found"})
			return
		}
	}

	if agentId == "" {
		c.JSON(404, gin.H{"error": "Agent is not found"})
		return
	}
	agent, err := h.aiGeneratorService.GetAgent(agentId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Agent is not found"})
		return
	}

	generator, err := h.aiGeneratorService.GetGeneratorFromID(agentId)
	if err != nil {
		c.JSON(404, gin.H{"error": "Generator is not found"})
		return
	}

	instruction := fmt.Sprintf(`%s
		
%s`, agent.SystemInstruction, `
Tolong jawab dalam format json : 
{
  "response": string,
  "type": string,
  "command": string,
  "params": object
}

jangan menggunakan json markdown 

Keterangan:
response: jawaban bila tipe nya pertanyaan
type: command atau question
command: jika tipe command
params: jika tipe command dibutuhkan parameter

`)
	generator.SetSystemInstruction(instruction)
	var his []ai_generator.AiMessage = []ai_generator.AiMessage{}

	for _, v := range histories {

		his = append(his, ai_generator.AiMessage{
			Role:    "user",
			Content: v.Input,
		})

		his = append(his, ai_generator.AiMessage{
			Role:    "assistant",
			Content: v.Output,
		})

	}
	fmt.Println(instruction)
	// h.aiGeneratorService.SetupModel(companySetting.GeminiAPIKey)
	utils.LogJson(his)
	output, err := generator.Generate(input.Content, nil, his)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	history := models.AiAgentHistory{
		Input:     input.Content,
		Output:    output.Content,
		AiAgentID: &agentId,
	}
	// ADD HISTORY TO DB
	h.aiGeneratorService.CreateHistory(&history)

	c.JSON(200, gin.H{"data": output.Content})
}

func (h *AiAgentHandler) GetAgentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)
	var companySetting com.CompanySetting
	err := h.ctx.DB.Find(&companySetting, "company_id = ?", member.CompanyID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Company setting is not found"})
		return
	}

	// h.aiGeneratorService.SetupAPIKey(companySetting.GeminiAPIKey)
	agents, err := h.aiGeneratorService.GetAgents(c.Request)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": agents})
}

func (h *AiAgentHandler) GetAgentDetailHandler(c *gin.Context) {
	id := c.Param("id")

	agent, err := h.aiGeneratorService.GetAgent(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": agent})

}
func (h *AiAgentHandler) GetAgentHistoriesHandler(c *gin.Context) {
	id := c.Param("id")

	// agent, err := h.aiGeneratorService.GetAgent(id)
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// 	return
	// }

	var histories []models.AiAgentHistory
	err := h.ctx.DB.Model(&models.AiAgentHistory{}).Find(&histories, "ai_agent_id = ?", id).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Agent histories is not found"})
		return
	}
	c.JSON(200, gin.H{"data": histories})

}

func (h *AiAgentHandler) DeleteHistoryHandler(c *gin.Context) {
	historyId := c.Param("historyId")

	err := h.ctx.DB.Delete(&models.AiAgentHistory{}, "id = ?", historyId).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "History deleted successfully"})
}

func (h *AiAgentHandler) UpdateHistoryHandler(c *gin.Context) {
	// id := c.Param("id")
	historyId := c.Param("historyId")

	var input models.AiAgentHistory
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = h.ctx.DB.Model(&models.AiAgentHistory{}).Where("id = ?", historyId).Updates(map[string]any{
		"input":       input.Input,
		"output":      input.Output,
		"ai_agent_id": input.AiAgentID,
	}).Error
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "History updated successfully"})
}

func (h *AiAgentHandler) ToggleModelHistoryHandler(c *gin.Context) {
	historyId := c.Param("historyId")

	var history models.AiAgentHistory
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
func (h *AiAgentHandler) CreateAgentHandler(c *gin.Context) {
	member := c.MustGet("member").(models.MemberModel)

	var companySetting com.CompanySetting
	err := h.ctx.DB.Find(&companySetting, "company_id = ?", member.CompanyID).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "Company setting is not found"})
		return
	}

	var input models.AiAgentModel
	err = c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if input.ApiKey == "" {
		input.ApiKey = companySetting.GeminiAPIKey
	}

	companyID := c.GetHeader("ID-Company")
	input.CompanyID = &companyID

	err = h.aiGeneratorService.CreateAgent(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent created successfully"})
}

func (h *AiAgentHandler) UpdateAgentHandler(c *gin.Context) {
	var input models.AiAgentModel
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	input.ID = id
	err := h.aiGeneratorService.UpdateAgent(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent updated successfully"})
}

func (h *AiAgentHandler) DeleteAgentHandler(c *gin.Context) {
	id := c.Param("id")
	err := h.aiGeneratorService.DeleteAgent(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Agent deleted successfully"})
}
