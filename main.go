package main

import (
	"plant-care-app/controllers"
	"plant-care-app/database"
	"plant-care-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database.Connect()
	controllers.CreateFirstAdmin()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
