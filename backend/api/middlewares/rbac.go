package middlewares

import (
	"github.com/AMETORY/ametory-erp-modules/auth"
	"github.com/AMETORY/ametory-erp-modules/context"
	"github.com/gin-gonic/gin"
)

func RbacUserMiddleware(erpContext *context.ERPContext, permissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.JSON(401, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}
		rbacSrv, ok := erpContext.RBACService.(*auth.RBACService)
		if !ok {
			c.JSON(500, gin.H{"message": "RBAC service is not available"})
			c.Abort()
			return
		}
		ok, err := rbacSrv.CheckPermission(userID.(string), permissions)
		if !ok {
			c.JSON(403, gin.H{"message": err.Error(), "error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
