package main

import (
	"plant-care-app/controllers"
	"plant-care-app/database"
	"plant-care-app/routes"
	"plant-care-app/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database.Connect()
	controllers.CreateFirstAdmin()
	controllers.SeedData()

	router := gin.Default()

	routes.SetupRoutes(router)

	scheduler := &services.Scheduler{}
	scheduler.CreateSchedulerAt(9)

	router.Run(":8080")
}
