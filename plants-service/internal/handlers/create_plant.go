package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"plant-care-app/plants-service/config"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PlantInput struct {
	Name             string `form:"Name" binding:"required"`
	WateringInterval string `form:"WateringInterval" binding:"required"`
	SpeciesID        string `form:"SpeciesID" binding:"required"`
}

func CreatePlant(c *gin.Context) {
	cfg := config.GetInstance()
	var plantInput PlantInput

	// Debug: Print all form data
	form, err := c.MultipartForm()
	if err != nil {
		fmt.Printf("Error getting multipart form: %v\n", err)
	} else {
		fmt.Printf("Form data: %+v\n", form.Value)
	}

	if err := c.ShouldBind(&plantInput); err != nil {
		fmt.Printf("Binding error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debug: Print the bound input
	fmt.Printf("Bound input: %+v\n", plantInput)

	// read image from request
	file, err := c.FormFile("Image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image file is not found " + err.Error()})
		return
	}

	// get container folder
	uploadDir := cfg.GetUploadDir()
	fmt.Println(uploadDir)
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error create upload directory " + err.Error()})
		return
	}

	// Generate file name
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	fullPath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, fullPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file " + err.Error()})
		return
	}

	// Convert string inputs to required types
	wateringInterval, err := strconv.Atoi(plantInput.WateringInterval)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid watering interval"})
		return
	}

	speciesID, err := strconv.Atoi(plantInput.SpeciesID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid species ID"})
		return
	}

	// Verify that the species exists
	var species models.Species
	if err := database.DB.First(&species, speciesID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Species not found"})
		return
	}

	userIDStr, _ := c.Get("user_id")
	userId, err := strconv.Atoi(fmt.Sprintf("%v", userIDStr))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request = fmt.Sprintf("%s/user:%v", cfg.GetUserServiceHost(), userId)
	response, err := http.Get(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	defer response.Body.Close()
	if err != nil || response.StatusCode == http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Create the plant
	newPlant := models.Plant{
		Name:             plantInput.Name,
		ImagePath:        fullPath,
		WateringInterval: wateringInterval,
		SpeciesID:        speciesID,
		UserID:           userId,
	}

	result := database.DB.Create(&newPlant)
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
