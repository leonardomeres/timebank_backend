package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/handlers"
	"github.com/leonardomeres/timebank_backend/internal/middleware"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	userGroup := router.Group("/api")
	{
		userGroup.POST("register", handlers.WithDB(db, handlers.Register))
		userGroup.POST("login", handlers.WithDB(db, handlers.Login))

		protected := userGroup.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("profile", handlers.WithDB(db, handlers.GetProfile))
			protected.POST("skills", handlers.WithDB(db, handlers.CreateSkill))
		}

	}
}
