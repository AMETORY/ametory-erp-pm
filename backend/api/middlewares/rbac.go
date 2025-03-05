package middlewares

import (
	"fmt"
	"strings"

	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/AMETORY/ametory-erp-modules/shared/models"
	"github.com/gin-gonic/gin"
)

func RbacUserMiddleware(erpContext *context.ERPContext, permissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		member := c.MustGet("member").(models.MemberModel)

		erpContext.DB.Preload("Role.Permissions").Find(&member)

		user := c.MustGet("user").(models.UserModel)
		user.Roles = []models.RoleModel{*member.Role}

		ok, err := CheckPermission(user, permissions)
		if !ok {
			c.JSON(403, gin.H{"message": err.Error(), "error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckPermission(user models.UserModel, permissionNames []string) (bool, error) {

	// Periksa izin
	for _, roleName := range permissionNames {
		for _, role := range user.Roles {
			if role.IsSuperAdmin {
				return true, nil
			}

			for _, permission := range role.Permissions {
				if permission.Name == roleName {
					return true, nil
				}
			}
		}
	}

	return false, fmt.Errorf("permissions %s not found", strings.Join(permissionNames, ", "))
}
