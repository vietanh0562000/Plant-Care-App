package main

import (
	"plant-care-app/controllers"
	"plant-care-app/database"
	"plant-care-app/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database.Connect()

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/create_plant", middlewares.AuthMiddleware(), controllers.CreatePlant)
	router.GET("/plants", controllers.GetPlants)
	router.GET("/users/:id", controllers.GetUser)

	router.Run(":8080")
}
