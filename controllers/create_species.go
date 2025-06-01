package controllers

import (
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"

	"github.com/gin-gonic/gin"
)

type SpeciesInput struct {
	Name  string `gorm:"primaryKey"`
	Notes string
}

func CreateSpecies(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not exist"})
		return
	}

	var user models.User
	err := database.DB.First(&user).Where("id = ?", userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
		return
	}

	if user.Role != "Admin" {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"error": "Not enough authorization"})
		return
	}

	var speciesInput SpeciesInput

	if err := c.ShouldBindJSON(&speciesInput); err != nil {
		c.JSON(http.StatusBadRequest, g.H{"error": err.Error()})
		return
	}
}
