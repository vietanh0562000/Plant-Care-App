package main

import (
	"plant-care-app/controllers"
	"plant-care-app/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	router := gin.Default()

	router.POST("/register", controllers.Register)

	router.Run(":8080")
}
