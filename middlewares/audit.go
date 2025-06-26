package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuditMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		ip := c.ClientIP()
		c.Set("request_id", requestID)
		c.Set("ip_address", ip)
		c.Next()
	}
}
