package config

import (
	"fmt"
	"log"
	"os"

	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB = db

	autoMigrate()
}

func autoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Skill{},
		&models.Offer{},
		&models.Transaction{},
	)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("✅ Database migrated successfully")
}
