package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/gorm"
)

func CreateBank(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context"})
		return
	}
	var bank models.TimeBank
	if err := c.ShouldBindJSON(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if bank.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank name"})
		return
	}
	var existingBank models.TimeBank
	if err := db.Select("id").Where("name = ?", bank.Name).First(&existingBank).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Bank already exists"})
		return
	}
	userIDFloat, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return
	}
	idUint := uint(userIDFloat)
	bank.CreatedByID = idUint
	if err := db.Create(&bank).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bank", "details": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Bank created successfully"})

}

func GetBankByUserID(c *gin.Context, db *gorm.DB) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in the HTTP context"})
		return
	}
	var banks []models.TimeBank
	var user models.User
	if err := db.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User not found"})
		return
	}

	if err := db.Where("created_by_id = ?", userID).Find(&banks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve banks", "details": err.Error()})
		return
	}
	if len(banks) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No banks found for this user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"banks": banks})
}
