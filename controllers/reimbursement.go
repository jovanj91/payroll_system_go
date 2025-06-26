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

type ReimbursementRequest struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

func SubmitReimbursement(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	var req ReimbursementRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reimbursement request"})
		return
	}

	db := config.DB

	r := models.Reimbursement{
		UserID:      userID,
		Amount:      req.Amount,
		Description: req.Description,
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   userID,
		UpdatedBy:   userID,
		IPAddress:   c.GetString("ip_address"),
		RequestID:   c.GetString("request_id"),
	}
	db.Create(&r)
	utils.CreateAuditLog(db, c, "reimbursements", "CREATE")
	c.JSON(http.StatusCreated, gin.H{"message": "Reimbursement submitted"})
}
