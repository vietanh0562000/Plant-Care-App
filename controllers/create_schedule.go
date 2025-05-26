package controllers

import (
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ScheduleInput struct {
	Time time.Time `json:"time" binding:"required"`
	Note string    `json:"note"`
}

func CreatePlantSchedule(c *gin.Context) {
	// TODO: Get userID, plantID, check if that user has that plant
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not exist"})
		return
	}

	// Convert userID to int
	uid, ok := userID.(float64)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	plantID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var plant models.Plant
	result := database.DB.First(&plant, plantID).Where("user_id = ?", int(uid))
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	var scheduleInput ScheduleInput
	if err := c.ShouldBindJSON(&scheduleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSchedule := models.Schedule{
		Time:    scheduleInput.Time,
		Note:    scheduleInput.Note,
		PlantID: uint(plantID),
	}

	if err := database.DB.Create(&newSchedule).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Create schedule success",
		"schedule": newSchedule,
	})
}
