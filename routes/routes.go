package routes

import (
	"plant-care-app/controllers"
	"plant-care-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected routes
	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		// User routes
		authorized.GET("/users/:id", controllers.GetUser)

		// Plant routes
		authorized.POST("/plants", controllers.CreatePlant)
		authorized.GET("/plants", controllers.GetPlants)

		// Species routes (read-only for regular users)
		authorized.GET("/species", controllers.GetSpecies)

		// Admin routes
		admin := authorized.Group("/")
		admin.Use(middlewares.RequireRole("admin"))
		{
			// Species management (admin only)
			admin.POST("/species", controllers.CreateSpeciesAdmin)
			admin.PUT("/species/:id", controllers.UpdateSpeciesAdmin)
			admin.DELETE("/species/:id", controllers.DeleteSpeciesAdmin)
		}
	}
}
