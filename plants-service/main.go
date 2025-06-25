package main

import (
	"fmt"
	"plant-care-app/plants-service/config"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/handlers"
	"plant-care-app/plants-service/internal/routes"
	services "plant-care-app/plants-service/pkg"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetInstance()
	database.Connect()
	database.SeedData()
	router := gin.Default()
	routes.SetupRoutes(router)

	var scheduler = services.Scheduler{}
	scheduler.CreateSchedulerAt(2, handlers.SendMailWatering)

	address := fmt.Sprintf(":%s", cfg.GetAppPort())
	router.Run(address)
}
