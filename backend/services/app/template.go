package app

import "encoding/json"

type ColumnTemplate struct {
	Industry string `json:"industry"`
	Columns  []struct {
		Name  string `json:"name"`
		Icon  string `json:"icon"`
		Color string `json:"color"`
	} `json:"columns"`
}

var templates = `[
    {
      "industry": "IT & Software Development",
      "columns": [
        { "name": "Backlog", "icon": "📝", "color": "#FFDDC1" },
        { "name": "In Progress", "icon": "🚧", "color": "#FFABAB" },
        { "name": "Done", "icon": "✅", "color": "#D4E157" },
        { "name": "Bugs", "icon": "🐞", "color": "#FF8A80" },
        { "name": "Under QA", "icon": "🔍", "color": "#B3E5FC" },
        { "name": "Finished", "icon": "🏁", "color": "#C5E1A5" },
        { "name": "Released", "icon": "🚀", "color": "#B39DDB" }
      ]
    },
    {
      "industry": "Digital Marketing & Content Creation",
      "columns": [
        { "name": "Strategy & Planning", "icon": "🎯", "color": "#FFCC80" },
        { "name": "Content Drafting", "icon": "✍️", "color": "#FFE082" },
        { "name": "Design & Media Creation", "icon": "🎨", "color": "#CE93D8" },
        { "name": "Scheduled", "icon": "📢", "color": "#F48FB1" },
        { "name": "Live & Monitoring", "icon": "📊", "color": "#81D4FA" },
        { "name": "Completed", "icon": "✅", "color": "#A5D6A7" }
      ]
    },
    {
      "industry": "Sales & CRM",
      "columns": [
        { "name": "Lead Generation", "icon": "📋", "color": "#FFECB3" },
        { "name": "Contacted", "icon": "📞", "color": "#F8BBD0" },
        { "name": "Negotiation", "icon": "🤝", "color": "#E1BEE7" },
        { "name": "Closed - Won", "icon": "🏁", "color": "#A5D6A7" },
        { "name": "Closed - Lost", "icon": "❌", "color": "#FFAB91" },
        { "name": "Follow-up", "icon": "🔄", "color": "#81D4FA" }
      ]
    },
    {
      "industry": "Manufacturing & Supply Chain",
      "columns": [
        { "name": "Raw Materials Ordered", "icon": "📦", "color": "#FFCCBC" },
        { "name": "Production In Progress", "icon": "🏗️", "color": "#A5D6A7" },
        { "name": "Quality Check", "icon": "🔍", "color": "#B3E5FC" },
        { "name": "Packaging", "icon": "📦", "color": "#FFE082" },
        { "name": "Shipped", "icon": "🚚", "color": "#F8BBD0" },
        { "name": "Delivered", "icon": "✅", "color": "#D4E157" }
      ]
    },
    {
      "industry": "Event Management",
      "columns": [
        { "name": "Planning", "icon": "📝", "color": "#FFDDC1" },
        { "name": "Scheduling", "icon": "📅", "color": "#B3E5FC" },
        { "name": "Vendor Coordination", "icon": "🎭", "color": "#F8BBD0" },
        { "name": "Event Live", "icon": "🔥", "color": "#FFAB91" },
        { "name": "Wrap-up & Feedback", "icon": "🛑", "color": "#FF8A80" },
        { "name": "Completed", "icon": "✅", "color": "#A5D6A7" }
      ]
    },
    {
      "industry": "Customer Support & Service Desk",
      "columns": [
        { "name": "New Tickets", "icon": "📥", "color": "#FFECB3" },
        { "name": "Assigned", "icon": "🎯", "color": "#FFCC80" },
        { "name": "In Progress", "icon": "🔄", "color": "#FFAB91" },
        { "name": "Escalated", "icon": "🛠️", "color": "#B39DDB" },
        { "name": "Resolved", "icon": "✅", "color": "#D4E157" },
        { "name": "Closed", "icon": "🎉", "color": "#81D4FA" }
      ]
    },
    {
      "industry": "Education & Course Development",
      "columns": [
        { "name": "Curriculum Planning", "icon": "📝", "color": "#FFDDC1" },
        { "name": "Content Creation", "icon": "✍️", "color": "#F8BBD0" },
        { "name": "Media Production", "icon": "🎥", "color": "#B3E5FC" },
        { "name": "Reviewed", "icon": "✅", "color": "#A5D6A7" },
        { "name": "Published", "icon": "📚", "color": "#CE93D8" },
        { "name": "Promotion", "icon": "📢", "color": "#FFCC80" }
      ]
    }
  ]`

func (a AppService) CreateDefaultColumnsFromTemplate() []ColumnTemplate {

	columns := []ColumnTemplate{}
	err := json.Unmarshal([]byte(templates), &columns)
	if err != nil {

	}
	return columns
}
