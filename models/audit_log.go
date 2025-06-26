package models

import (
	"time"

	"github.com/google/uuid"
)

type AuditLog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Table     string
	Action    string
	UserID    uuid.UUID
	RequestID string
	IPAddress string
	CreatedAt time.Time
}
