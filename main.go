package main

import (
	"log"

	"github.com/jovanj91/payroll_system_go/config"
	"github.com/jovanj91/payroll_system_go/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	config.AutoMigrate(db)
	config.SeedInitialData()

	r := gin.Default()
	routes.SetupRoutes(r)

	log.Println("Server running at :8080")
	r.Run(":8080")
}
