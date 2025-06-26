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
	runRule := fmt.Sprintf("0 */1 * * * *")
	scheduler.CreateSchedulerAt(runRule, handlers.SendWateringNotification)

	address := fmt.Sprintf(":%s", cfg.GetAppPort())
	router.Run(address)
}
