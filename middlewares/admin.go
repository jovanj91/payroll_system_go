package middlewares

import (
	"net/http"

	"github.com/jovanj91/payroll_system_go/models"

	"github.com/gin-gonic/gin"
)

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != string(models.AdminRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Admin access only"})
			return
		}
		c.Next()
	}
}
