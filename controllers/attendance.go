package controllers

import (
	"net/http"
	"time"

	"github.com/jovanj91/payroll_system_go/config"
	"github.com/jovanj91/payroll_system_go/models"
	"github.com/jovanj91/payroll_system_go/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SubmitAttendance(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	role := c.MustGet("user_role").(string)

	if role != string(models.EmployeeRole) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only employees can submit attendance"})
		return
	}

	now := time.Now()
	dateOnly := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	weekday := now.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot submit attendance on weekends"})
		return
	}

	db := config.DB
	var existing models.Attendance
	err := db.Where("user_id = ? AND date = ?", userID, now.Format("2006-01-02")).First(&existing).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Attendance already submitted today"})
		return
	}

	a := models.Attendance{
		UserID:    userID,
		Date:      dateOnly,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: userID,
		UpdatedBy: userID,
		IPAddress: c.GetString("ip_address"),
		RequestID: c.GetString("request_id"),
	}
	db.Create(&a)
	utils.CreateAuditLog(db, c, "attendances", "CREATE")
	c.JSON(http.StatusCreated, gin.H{"message": "Attendance submitted"})
}
