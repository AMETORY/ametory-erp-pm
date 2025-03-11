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
    },
    {
      "industry": "HR Recruitment",
      "columns": [
        { "name": "Job Posting", "icon": "📢", "color": "#FFECB3" },
        { "name": "Applications Received", "icon": "📄", "color": "#FFDDC1" },
        { "name": "Screening", "icon": "🧐", "color": "#F8BBD0" },
        { "name": "Interview Scheduled", "icon": "📅", "color": "#B3E5FC" },
        { "name": "Interviewed", "icon": "🤝", "color": "#A5D6A7" },
        { "name": "Offer Extended", "icon": "💼", "color": "#FFE082" },
        { "name": "Hired", "icon": "✅", "color": "#D4E157" },
        { "name": "Onboarding", "icon": "📋", "color": "#81D4FA" }
      ]
    },
    {
      "industry": "Procurement",
      "columns": [
        { "name": "Request for Quotation", "icon": "📋", "color": "#FFDDC1" },
        { "name": "Quotations Received", "icon": "📩", "color": "#FFABAB" },
        { "name": "Vendor Evaluation", "icon": "🧐", "color": "#B3E5FC" },
        { "name": "Purchase Order Created", "icon": "📝", "color": "#A5D6A7" },
        { "name": "Order Shipped", "icon": "🚚", "color": "#FFCC80" },
        { "name": "Received & Inspected", "icon": "🔎", "color": "#D4E157" },
        { "name": "Payment Processed", "icon": "💰", "color": "#F48FB1" },
        { "name": "Completed", "icon": "✅", "color": "#81D4FA" }
      ]
    },
    {
      "industry": "Supply Chain",
      "columns": [
        { "name": "Demand Planning", "icon": "📊", "color": "#FFECB3" },
        { "name": "Supplier Selection", "icon": "🏷️", "color": "#FFDDC1" },
        { "name": "Purchase Order", "icon": "📝", "color": "#FFE082" },
        { "name": "Inbound Logistics", "icon": "🚚", "color": "#B3E5FC" },
        { "name": "Inventory Management", "icon": "📦", "color": "#A5D6A7" },
        { "name": "Production", "icon": "🏗️", "color": "#FFCC80" },
        { "name": "Quality Control", "icon": "🔎", "color": "#CE93D8" },
        { "name": "Outbound Logistics", "icon": "🚛", "color": "#F8BBD0" },
        { "name": "Delivered", "icon": "✅", "color": "#D4E157" }
      ]
    },
    {
      "industry": "Production Process",
      "columns": [
        { "name": "Material Preparation", "icon": "🧱", "color": "#FFDDC1" },
        { "name": "Work Order Created", "icon": "📝", "color": "#FFE082" },
        { "name": "Production Planning", "icon": "📊", "color": "#FFCC80" },
        { "name": "In Progress", "icon": "⚙️", "color": "#B3E5FC" },
        { "name": "Quality Check", "icon": "🔎", "color": "#CE93D8" },
        { "name": "Rework / Adjustment", "icon": "🔧", "color": "#FFABAB" },
        { "name": "Packaging", "icon": "📦", "color": "#A5D6A7" },
        { "name": "Ready for Delivery", "icon": "🚚", "color": "#F8BBD0" },
        { "name": "Delivered", "icon": "✅", "color": "#D4E157" },
        { "name": "Post-Production Review", "icon": "📝", "color": "#81D4FA" }
      ]
    },
    {
      "industry": "Credit Application",
      "columns": [
        { "name": "Application Received", "icon": "📥", "color": "#FFECB3" },
        { "name": "Document Verification", "icon": "📑", "color": "#FFDDC1" },
        { "name": "Credit Assessment", "icon": "💰", "color": "#FFE082" },
        { "name": "Background Check", "icon": "🕵️", "color": "#B3E5FC" },
        { "name": "Approval Process", "icon": "✅", "color": "#A5D6A7" },
        { "name": "Contract Signing", "icon": "📝", "color": "#FFCC80" },
        { "name": "Fund Disbursement", "icon": "💸", "color": "#F8BBD0" },
        { "name": "Repayment Monitoring", "icon": "📊", "color": "#81D4FA" },
        { "name": "Closed", "icon": "🔒", "color": "#D4E157" }
      ]
    },
    {
      "industry": "Repair Service",
      "columns": [
        { "name": "Service Request Received", "icon": "📥", "color": "#FFECB3" },
        { "name": "Initial Assessment", "icon": "🔍", "color": "#FFDDC1" },
        { "name": "Cost Estimation", "icon": "💰", "color": "#FFE082" },
        { "name": "Customer Approval", "icon": "📝", "color": "#B3E5FC" },
        { "name": "Parts Ordering", "icon": "🛒", "color": "#A5D6A7" },
        { "name": "In Progress", "icon": "⚙️", "color": "#FFCC80" },
        { "name": "Quality Control", "icon": "✅", "color": "#CE93D8" },
        { "name": "Ready for Pickup", "icon": "🚗", "color": "#F8BBD0" },
        { "name": "Delivered", "icon": "📦", "color": "#81D4FA" },
        { "name": "Follow-up", "icon": "📞", "color": "#D4E157" }
      ]
    },
    {
      "industry": "Financial/Tax Audit",
      "columns": [
        { "name": "Engagement Planning", "icon": "📋", "color": "#FFECB3" },
        { "name": "Document Submission", "icon": "📄", "color": "#FFDDC1" },
        { "name": "Preliminary Review", "icon": "🔍", "color": "#FFE082" },
        { "name": "Risk Assessment", "icon": "⚠️", "color": "#B3E5FC" },
        { "name": "Fieldwork", "icon": "🏗️", "color": "#A5D6A7" },
        { "name": "Issue Identification", "icon": "❗", "color": "#FFABAB" },
        { "name": "Management Discussion", "icon": "💬", "color": "#FFCC80" },
        { "name": "Report Drafting", "icon": "📝", "color": "#CE93D8" },
        { "name": "Client Review", "icon": "🧐", "color": "#81D4FA" },
        { "name": "Final Report Submission", "icon": "📁", "color": "#D4E157" },
        { "name": "Follow-up & Advisory", "icon": "📞", "color": "#F8BBD0" }
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
