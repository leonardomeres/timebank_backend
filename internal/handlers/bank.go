package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/services"
	"gorm.io/gorm"
)

func TimeBankHandler(c *gin.Context, db *gorm.DB) {
	switch c.Request.Method {
	case http.MethodPost:
		services.CreateBank(c, db)
	case http.MethodGet:
		services.GetBankByUserID(c, db)
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
}
