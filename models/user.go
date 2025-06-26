package models

import (
	"time"

	"github.com/google/uuid"
)

type UserRole string

const (
	AdminRole    UserRole = "admin"
	EmployeeRole UserRole = "employee"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string    `gorm:"uniqueIndex;not null"`
	Password  string    `gorm:"not null"`
	Salary    float64
	Role      UserRole `gorm:"type:varchar(20);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	IPAddress string
	RequestID string
}
