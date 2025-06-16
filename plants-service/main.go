package main

import (
	"fmt"
	"plant-care-app/plants-service/config"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetInstance()
	database.Connect()
	router := gin.Default()
	routes.SetupRoutes(router)

	address := fmt.Sprintf(":%s", cfg.GetAppPort())
	router.Run(address)
}
