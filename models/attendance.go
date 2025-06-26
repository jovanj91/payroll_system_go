package models

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	UserID    uuid.UUID
	Date      time.Time `gorm:"type:date;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	IPAddress string
	RequestID string
}
