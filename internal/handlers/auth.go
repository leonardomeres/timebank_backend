package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(c *gin.Context, db *gorm.DB) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	input.Password = string(hashedPassword)

	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context, db *gorm.DB) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password credentials", "details": err.Error()})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 1 day expiry
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func WithDB(db *gorm.DB, handler func(*gin.Context, *gorm.DB)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, db)
	}
}
