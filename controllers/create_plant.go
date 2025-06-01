package controllers

import (
	"fmt"
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"

	"github.com/gin-gonic/gin"
)

type PlantInput struct {
	Name             string `json:"Name" binding:"required"`
	ImageURL         string `json:"ImageUrl" binding:"required"`
	WateringInterval int    `json:"WateringInterval" binding:"required"`
	SpeciesID        int    `json:"SpeciesID" binding:"required"`
}

func CreatePlant(c *gin.Context) {
	var plantInput PlantInput

	if err := c.ShouldBindJSON(&plantInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Convert userID to int
	uid, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	var user models.User
	result := database.DB.First(&user, int(uid))
	if result.Error != nil {
		fmt.Printf("Error finding user with ID %v: %v\n", uid, result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	fmt.Printf("Found user: ID=%d, Name=%s\n", user.ID, user.Name)

	// Verify that the species exists
	var species models.Species
	if err := database.DB.First(&species, plantInput.SpeciesID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Species not found"})
		return
	}

	// Create the plant with the existing species ID
	newPlant := models.Plant{
		Name:             plantInput.Name,
		ImageURL:         plantInput.ImageURL,
		WateringInterval: plantInput.WateringInterval,
		UserID:           user.ID,
		SpeciesID:        plantInput.SpeciesID,
	}

	result = database.DB.Create(&newPlant)
	if result.Error != nil {
		fmt.Printf("Error creating plant: %v\n", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create plant",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Create new plant successfully",
		"plant":   newPlant,
	})
}
