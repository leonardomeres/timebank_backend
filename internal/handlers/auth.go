package handlers

import (
	"net/http"

	_ "github.com/leonardomeres/timebank_backend/doc"
	"github.com/leonardomeres/timebank_backend/internal/auth"
	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register godoc
// @Summary      Register new user
// @Description  Creates a new user account with hashed password
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user  body  models.User  true  "User registration input"
// @Success      201  {object}  map[string]string
// @Failure      400  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /register [post]
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
	input.IsActive = true // Default to active
	input.Balance = 4.0   // Default balance
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

	token, err := auth.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func WithDB(db *gorm.DB, handler func(*gin.Context, *gorm.DB)) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(c, db)
	}
}
