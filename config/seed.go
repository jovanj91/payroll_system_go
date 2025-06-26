package config

import (
	"log"

	"github.com/jovanj91/payroll_system_go/models"
	"github.com/jovanj91/payroll_system_go/utils"

	"github.com/brianvoe/gofakeit/v6"
)

func SeedInitialData() {
	db := DB
	AutoMigrate(db)

	var count int64

	db.Model(&models.User{}).Count(&count)
	if count > 0 {
		log.Println("ℹUsers already seeded. Skipping seed.")
		return
	}

	adminPassword, _ := utils.HashPassword("admin123")
	db.Create(&models.User{
		Username:  "admin",
		Password:  adminPassword,
		Role:      models.AdminRole,
		Salary:    0,
		IPAddress: "127.0.0.1",
		RequestID: "seed-admin",
	})

	testPassword, _ := utils.HashPassword("test123")
	db.Create(&models.User{
		Username:  "test",
		Password:  testPassword,
		Role:      models.EmployeeRole,
		Salary:    1000000,
		IPAddress: "127.0.0.1",
		RequestID: "seed-user",
	})

	for i := 0; i < 99; i++ {
		password, _ := utils.HashPassword("password")
		db.Create(&models.User{
			Username:  gofakeit.Username(),
			Password:  password,
			Role:      models.EmployeeRole,
			Salary:    float64(gofakeit.Number(3000000, 10000000)),
			IPAddress: "127.0.0.1",
			RequestID: "seed-user",
		})
	}

	log.Println("Seeded 1 admin and 100 employees.")
}
