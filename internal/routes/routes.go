package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/handlers"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	router.POST("/register", handlers.WithDB(db, handlers.Register))
	router.POST("/login", handlers.WithDB(db, handlers.Login))
}
