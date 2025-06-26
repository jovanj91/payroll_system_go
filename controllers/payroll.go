package controllers

import (
	"net/http"
	"time"

	"github.com/jovanj91/payroll_system_go/models"
	"github.com/jovanj91/payroll_system_go/utils"

	"github.com/jovanj91/payroll_system_go/config"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PayrollPeriodRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

func CreatePayrollPeriod(c *gin.Context) {
	var req PayrollPeriodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	start, _ := time.Parse("2006-01-02", req.StartDate)
	end, _ := time.Parse("2006-01-02", req.EndDate)
	period := models.PayrollPeriod{
		StartDate: start,
		EndDate:   end,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: c.MustGet("user_id").(uuid.UUID),
		UpdatedBy: c.MustGet("user_id").(uuid.UUID),
		IPAddress: c.GetString("ip_address"),
		RequestID: c.GetString("request_id"),
	}
	config.DB.Create(&period)
	c.JSON(http.StatusCreated, period)
}

func RunPayroll(c *gin.Context) {
	periodID := c.Param("id")
	db := config.DB
	var period models.PayrollPeriod
	if err := db.First(&period, "id = ?", periodID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payroll period not found"})
		return
	}
	if period.Processed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Payroll already processed"})
		return
	}

	var users []models.User
	db.Where("role = ?", models.EmployeeRole).Find(&users)

	for _, user := range users {
		var attendances []models.Attendance
		db.Where("user_id = ? AND date BETWEEN ? AND ?", user.ID, period.StartDate, period.EndDate).Find(&attendances)

		var overtimes []models.Overtime
		db.Where("user_id = ? AND date BETWEEN ? AND ?", user.ID, period.StartDate, period.EndDate).Find(&overtimes)

		var reimbursements []models.Reimbursement
		db.Where("user_id = ? AND date BETWEEN ? AND ?", user.ID, period.StartDate, period.EndDate).Find(&reimbursements)

		attendedDays := len(attendances)
		workingDays := 0
		for d := period.StartDate; !d.After(period.EndDate); d = d.AddDate(0, 0, 1) {
			if d.Weekday() != time.Saturday && d.Weekday() != time.Sunday {
				workingDays++
			}
		}
		proratedSalary := (float64(attendedDays) / float64(workingDays)) * user.Salary

		overtimeHours := 0.0
		for _, o := range overtimes {
			overtimeHours += o.Hours
		}
		overtimePay := overtimeHours * ((user.Salary / float64(workingDays)) / 8.0) * 2

		reimbursement := 0.0
		for _, r := range reimbursements {
			reimbursement += r.Amount
		}

		total := proratedSalary + overtimePay + reimbursement

		db.Create(&models.Payslip{
			UserID:         user.ID,
			PeriodID:       period.ID,
			BaseSalary:     proratedSalary,
			AttendanceDays: attendedDays,
			OvertimeHours:  overtimeHours,
			OvertimePay:    overtimePay,
			Reimbursement:  reimbursement,
			TotalPay:       total,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			CreatedBy:      c.MustGet("user_id").(uuid.UUID),
			UpdatedBy:      c.MustGet("user_id").(uuid.UUID),
			IPAddress:      c.GetString("ip_address"),
			RequestID:      c.GetString("request_id"),
		})
	}

	db.Model(&period).Update("processed", true)
	utils.CreateAuditLog(db, c, "payroll_periods", "CREATE")
	c.JSON(http.StatusOK, gin.H{"message": "Payroll processed"})
}
