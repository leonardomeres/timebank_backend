package main

import (
	"log"

	"github.com/leonardomeres/timebank_backend/internal/config"
	"github.com/leonardomeres/timebank_backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// Initialize database connection
	config.InitDB()

	// Create Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, config.DB)
	// Run the server
	r.Run(":8080")
}
