package models

import (
	"time"

	"github.com/google/uuid"
)

type Overtime struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID
	Date      time.Time `gorm:"index;not null"`
	Hours     float64   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	IPAddress string
	RequestID string
}
