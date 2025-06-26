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

type OvertimeRequest struct {
	Hours float64 `json:"hours"`
}

func SubmitOvertime(c *gin.Context) {

	userID := c.MustGet("user_id").(uuid.UUID)
	var req OvertimeRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Hours <= 0 || req.Hours > 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid overtime hours (1-3 allowed)"})
		return
	}

	now := time.Now()
	db := config.DB
	o := models.Overtime{
		UserID:    userID,
		Date:      now,
		Hours:     req.Hours,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: userID,
		UpdatedBy: userID,
		IPAddress: c.GetString("ip_address"),
		RequestID: c.GetString("request_id"),
	}
	db.Create(&o)
	utils.CreateAuditLog(db, c, "overtimes", "CREATE")
	c.JSON(http.StatusCreated, gin.H{"message": "Overtime submitted"})
}
