package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leonardomeres/timebank_backend/internal/models"
	"gorm.io/gorm"
)

// GetProfile godoc
// @Summary      Get user profile
// @Description  Retrieves the profile of the authenticated user
// @Tags         user
// @Produce      json
// @Success      200  {object}  models.GetProfileResponse
// @Failure      401  {object}  models.GenericUnauthorizedResponse
// @Failure      500  {object}  models.GenericErrorResponse
// @Router       /profile [get]
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

// CreateSkill godoc
// @Summary      Create a new skill
// @Description  Allows authenticated users to create a new skill
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        skill  body  models.Skill  true  "Skill creation input"
// @Success      201  {object}  models.SkillCreationResponse
// @Failure      401  {object}  models.GenericUnauthorizedResponse
// @Failure      409  {object}  models.SkillExistsResponse
// @Failure      500  {object}  models.SkillFailCreationResponse
// @Router       /skills [post]
func CreateSkill(c *gin.Context, db *gorm.DB) {
	var input models.Skill
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDRaw, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userIDFloat, ok := userIDRaw.(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID type"})
		return
	}

	userID := uint(userIDFloat)

	//TODO: Validate AreaID if necessary
	/* var are models.Area
	if err := db.First(&are, input.AreaID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Area not found"})
		return
	}
	*/
	skill := models.Skill{
		Name:        input.Name,
		AreaID:      input.AreaID,
		Description: input.Description,
		CreatedByID: userID,
		CreatedBy:   models.User{ID: userID},
	}

	if err := db.Create(&skill).Error; err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"uni_skills_name\" (SQLSTATE 23505)" {
			c.JSON(http.StatusConflict, gin.H{"error": "Skill already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create skill"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Skill created successfully"})

}
