package config

import (
	"fmt"
	"log"

	"github.com/jovanj91/payroll_system_go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=localhost user=postgres password=root dbname=db_payroll_system_go port=5432 sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	DB = db
	return db
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Attendance{},
		&models.Overtime{},
		&models.Reimbursement{},
		&models.PayrollPeriod{},
		&models.Payslip{},
		&models.AuditLog{},
	)
}
