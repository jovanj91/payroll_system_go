package models

import (
	"time"

	"github.com/google/uuid"
)

type PayrollPeriod struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`
	Processed bool      `gorm:"default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	IPAddress string
	RequestID string
}
