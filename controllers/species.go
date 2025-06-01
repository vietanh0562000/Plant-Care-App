package controllers

import (
	"net/http"
	"plant-care-app/database"
	"plant-care-app/models"

	"github.com/gin-gonic/gin"
)

// CreateSpeciesAdmin creates a new plant species (admin only)
func CreateSpeciesAdmin(c *gin.Context) {
	var input SpeciesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	species := models.Species{
		Name:  input.Name,
		Notes: input.Notes,
	}

	if err := database.DB.Create(&species).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create species",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Species created successfully",
		"species": species,
	})
}

// GetSpecies returns all plant species
func GetSpecies(c *gin.Context) {
	var species []models.Species
	if err := database.DB.Find(&species).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch species",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"species": species,
	})
}

// UpdateSpeciesAdmin updates an existing species (admin only)
func UpdateSpeciesAdmin(c *gin.Context) {
	id := c.Param("id")
	var input SpeciesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var species models.Species
	if err := database.DB.First(&species, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Species not found"})
		return
	}

	species.Name = input.Name
	species.Notes = input.Notes

	if err := database.DB.Save(&species).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update species",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Species updated successfully",
		"species": species,
	})
}

// DeleteSpeciesAdmin deletes a species (admin only)
func DeleteSpeciesAdmin(c *gin.Context) {
	id := c.Param("id")

	// Check if any plants are using this species
	var count int64
	database.DB.Model(&models.Plant{}).Where("species_id = ?", id).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Cannot delete species that is being used by plants",
		})
		return
	}

	if err := database.DB.Delete(&models.Species{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete species",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Species deleted successfully",
	})
}
