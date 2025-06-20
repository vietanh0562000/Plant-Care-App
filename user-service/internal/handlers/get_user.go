package handlers

import (
	"net/http"
	"plant-care-app/user-service/internal/database"
	"plant-care-app/user-service/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userID := c.Param("userID")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": user})
}
