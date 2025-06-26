package models

import (
	"time"

	"github.com/google/uuid"
)

type Payslip struct {
	ID             uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID         uuid.UUID
	PeriodID       uuid.UUID
	BaseSalary     float64
	AttendanceDays int
	OvertimeHours  float64
	OvertimePay    float64
	Reimbursement  float64
	TotalPay       float64
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreatedBy      uuid.UUID
	UpdatedBy      uuid.UUID
	IPAddress      string
	RequestID      string
}
