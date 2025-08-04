package main

import (
	"log"

	"github.com/leonardomeres/timebank_backend/internal/config"

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
	config.InitDB()

	r := gin.Default()
	//TODO: Setup routes here
	// setupRoutes(r)
	r.Run(":8080")
}
