package models

import (
	"time"

	"github.com/google/uuid"
)

type Reimbursement struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID      uuid.UUID
	Amount      float64
	Description string
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreatedBy   uuid.UUID
	UpdatedBy   uuid.UUID
	IPAddress   string
	RequestID   string
}
