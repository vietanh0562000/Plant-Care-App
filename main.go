package main

import (
	"plant-care-app/controllers"
	"plant-care-app/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	database.Connect()

	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	router.Run(":8080")
}
