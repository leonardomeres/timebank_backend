package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/gorm"
)

func GetProfile(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}

	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      user.ID,
		"name":    user.Name,
		"email":   user.Email,
		"skills":  user.Skills,
		"balance": user.Balance,
	})
}
