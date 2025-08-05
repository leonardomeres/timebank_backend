package main

import (
	"log"

	_ "github.com/leonardomeres/timebank_backend/docs"
	"github.com/leonardomeres/timebank_backend/internal/config"
	"github.com/leonardomeres/timebank_backend/internal/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// @title       TimeBank API
// @version     1.0
// @description API for user registration, login, and skill management in TimeBank.
// @host        localhost:8080
// @BasePath    /api
// @schemes     http
func main() {
	// Initialize database connection
	db := config.InitDB()

	// Create Gin router
	r := gin.Default()

	//Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.SetupRoutes(r, db)
	// Run the server
	r.Run(":8080")
}
