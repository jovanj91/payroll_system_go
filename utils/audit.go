package utils

import (
	"time"

	"github.com/jovanj91/payroll_system_go/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateAuditLog(db *gorm.DB, c *gin.Context, table, action string) {

	userID := c.MustGet("user_id").(uuid.UUID)
	ip := c.GetString("ip_address")
	requestID := c.GetString("request_id")

	log := models.AuditLog{
		Table:     table,
		Action:    action,
		UserID:    userID,
		IPAddress: ip,
		RequestID: requestID,
		CreatedAt: time.Now(),
	}
	db.Create(&log)
}
