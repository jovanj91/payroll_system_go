package controllers

import (
	"net/http"

	"github.com/jovanj91/payroll_system_go/config"
	"github.com/jovanj91/payroll_system_go/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPayslip(c *gin.Context) {
	userID := c.MustGet("user_id").(uuid.UUID)
	periodID := c.Param("id")

	var payslip models.Payslip
	db := config.DB
	err := db.Where("user_id = ? AND period_id = ?", userID, periodID).First(&payslip).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payslip not found"})
		return
	}
	c.JSON(http.StatusOK, payslip)
}

func GetPayslipSummary(c *gin.Context) {
	periodID := c.Param("id")
	db := config.DB

	var payslips []models.Payslip
	db.Where("period_id = ?", periodID).Find(&payslips)

	summary := []gin.H{}
	total := 0.0
	for _, p := range payslips {
		summary = append(summary, gin.H{
			"user_id":   p.UserID,
			"total_pay": p.TotalPay,
		})
		total += p.TotalPay
	}

	c.JSON(http.StatusOK, gin.H{
		"summary":         summary,
		"total_disbursed": total,
	})
}
