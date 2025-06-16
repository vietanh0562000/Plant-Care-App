package routes

import (
	"plant-care-app/plants-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	r.POST("/plants", handlers.CreatePlant)
	r.GET("/plants", handlers.GetPlants)
	r.GET("/plants/need-watering", handlers.GetPlantsNeedWatering)

	// Species routes (read-only for regular users)
	r.GET("/species", handlers.GetSpecies)
	// Species management (admin only)
	r.POST("/species", handlers.CreateSpeciesAdmin)
	r.PUT("/species/:id", handlers.UpdateSpeciesAdmin)
	r.DELETE("/species/:id", handlers.DeleteSpeciesAdmin)
}
