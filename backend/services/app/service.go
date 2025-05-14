package app

import (
	"ametory-pm/config"
	"ametory-pm/models/company"

	con "ametory-pm/models/connection"
	"ametory-pm/models/project"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/AMETORY/ametory-erp-modules/utils"
	"github.com/go-redis/redis/v8"
	"gopkg.in/olahol/melody.v1"
)

type AppService struct {
	ctx               *context.ERPContext
	Config            *config.Config
	Redis             *redis.Client
	Websocket         *melody.Melody
	ConnectionService *ConnectionService
}

func NewAppService(erpContext *context.ERPContext, config *config.Config, redis *redis.Client, ws *melody.Melody) *AppService {
	if !erpContext.SkipMigration {
		erpContext.DB.AutoMigrate(
			&project.ProjectPreferenceModel{},
			&company.CompanySetting{},
			&con.ConnectionModel{},
		)
	}
	return &AppService{
		ctx:               erpContext,
		Config:            config,
		Redis:             redis,
		Websocket:         ws,
		ConnectionService: NewConnectionService(erpContext),
	}
}

var (
	cruds    = []string{"create", "read", "update", "delete"}
	services = map[string][]map[string][]string{
		"auth": {{"user": cruds, "admin": cruds, "rbac": cruds}},
		"contact": {
			{"customer": cruds},
		},
		"company": {
			{"company": append(cruds, "approval")},
		},
		"customer_relationship": {
			{"whatsapp": cruds},
			{"whatsapp_template": cruds},
			{"form_template": cruds},
			{"form": cruds},
		},
		"inventory": {
			{"product": cruds},
			{"product_category": cruds},
		},
		"whatsapp": {
			{"broadcast": append(cruds, "send")},
			{"message": append(cruds, "send")},
		},
		"misc": {
			{"tag": cruds},
		},
		"project_management": {
			{"project": cruds},
			{"member": append(cruds, "approval", "invite")},
			{"task": cruds},
		},
	}
)

func (AppService) GenerateDefaultPermissions() []models.PermissionModel {
	var permissions []models.PermissionModel

	for service, modules := range services {
		for _, module := range modules {
			for key, actions := range module {
				for _, action := range actions {
					permissions = append(permissions, models.PermissionModel{Name: service + ":" + key + ":" + action})
				}
			}
		}
	}

	return permissions
}

func (a AppService) GenerateDefaultRoles(companyID string) []models.RoleModel {
	var roles []models.RoleModel
	roles = append(roles, models.RoleModel{Name: "SUPER ADMIN", IsSuperAdmin: true, IsOwner: true, CompanyID: &companyID})
	cruds = []string{"create", "read", "update", "delete"}
	services = map[string][]map[string][]string{
		"contact": {
			{"customer": cruds},
		},
		"customer_relationship": {
			{"whatsapp": cruds},
			{"form_template": cruds},
			{"whatsapp_template": cruds},
			{"form": cruds},
		},
		"misc": {
			{"tag": cruds},
		},
		"whatsapp": {
			{"broadcast": append(cruds, "send")},
			{"message": append(cruds, "send")},
		},
		"project_management": {
			{"project": cruds},
			{"member": append(cruds, "approval", "invite")},
			{"task": cruds},
		},
	}
	permissionNames := []string{}
	for service, modules := range services {
		for _, module := range modules {
			for key, actions := range module {
				for _, action := range actions {
					permissionNames = append(permissionNames, service+":"+key+":"+action)
				}
			}
		}
	}

	permissions := []models.PermissionModel{}

	a.ctx.DB.Where("name in (?)", permissionNames).Find(&permissions)
	roles = append(roles, models.RoleModel{Name: "ADMIN", Permissions: permissions, CompanyID: &companyID})

	services = map[string][]map[string][]string{
		"contact": {
			{"customer": cruds},
		},
		"customer_relationship": {
			{"whatsapp": []string{"read"}},
			{"form_template": []string{"read"}},
			{"whatsapp_template": cruds},
			{"form": []string{"read"}},
		},
		"whatsapp": {
			{"broadcast": append(cruds, "send")},
			{"message": append(cruds, "send")},
		},
		"misc": {
			{"tag": cruds},
		},
		"project_management": {
			{"project": []string{"read"}},
			{"member": append([]string{"read"}, "invite")},
			{"task": cruds},
		},
	}
	permissionNames = []string{}
	for service, modules := range services {
		for _, module := range modules {
			for key, actions := range module {
				for _, action := range actions {
					permissionNames = append(permissionNames, service+":"+key+":"+action)
				}
			}
		}
	}

	permissions = []models.PermissionModel{}

	a.ctx.DB.Where("name in (?)", permissionNames).Find(&permissions)
	roles = append(roles, models.RoleModel{Name: "MEMBER", Permissions: permissions, CompanyID: &companyID})
	utils.LogJson(roles)

	for i, v := range roles {
		var role models.RoleModel
		a.ctx.DB.Where("name = ? and company_id = ?", v.Name, companyID).Find(&role)
		if role.ID == "" {
			a.ctx.DB.Create(&v)
		} else {
			a.ctx.DB.Model(&role).Association("Permissions").Replace(v.Permissions)
		}
		roles[i] = v
	}

	return roles
}
