package controllers

import (
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"

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
