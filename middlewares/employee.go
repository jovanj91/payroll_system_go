package middlewares

import (
	"net/http"

	"github.com/jovanj91/payroll_system_go/models"

	"github.com/gin-gonic/gin"
)

func AuthEmployee() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("user_role")
		if !exists || role != string(models.EmployeeRole) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Only employees can access this route"})
			return
		}
		c.Next()
	}
}
