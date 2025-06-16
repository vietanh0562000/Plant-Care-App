package handlers

import (
	"net/http"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/models"

	"github.com/gin-gonic/gin"
)

type SpeciesInput struct {
	Name  string `json:"name" binding:"required"`
	Notes string
}

func CreateSpecies(c *gin.Context) {
	var speciesInput SpeciesInput

	if err := c.ShouldBindJSON(&speciesInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create new species
	species := models.Species{
		Name:  speciesInput.Name,
		Notes: speciesInput.Notes,
	}

	if err := database.DB.Create(&species).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, species)
}
