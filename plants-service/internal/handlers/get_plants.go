package handlers

import (
	"net/http"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPlants(c *gin.Context) {
	var plants []models.Plant

	result := database.DB.Find(&plants)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, plants)
}

func GetPlantsNeedWatering(c *gin.Context) {
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

	var plants []models.Plant
	now := time.Now()

	// Find plants where:
	// 1. Belongs to the user
	// 2. Last watering time + watering interval is before or equal to now
	query := `
		SELECT * FROM plants 
		WHERE user_id = ? 
		AND last_time_watering + (watering_interval || ' days')::interval <= ?
		ORDER BY last_time_watering ASC
	`
	result := database.DB.Raw(query, int(uid), now).Scan(&plants)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch plants"})
		return
	}

	// Add next watering date to response
	type PlantResponse struct {
		models.Plant
		NextWateringDate  time.Time `json:"next_watering_date"`
		DaysUntilWatering int       `json:"days_until_watering"`
	}

	var response []PlantResponse
	for _, plant := range plants {
		nextWatering := plant.LastTimeWatering.AddDate(0, 0, plant.WateringInterval)
		daysUntil := int(nextWatering.Sub(now).Hours() / 24)

		response = append(response, PlantResponse{
			Plant:             plant,
			NextWateringDate:  nextWatering,
			DaysUntilWatering: daysUntil,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"plants": response,
	})
}
