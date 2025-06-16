package main

import (
	"fmt"
	"plant-care-app/user-service/config"
	"plant-care-app/user-service/internal/database"
	"plant-care-app/user-service/internal/handlers"
	"plant-care-app/user-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	database.Connect(cfg)
	handlers.CreateFirstAdmin()
	database.SeedData()

	router := gin.Default()

	routes.SetupRoutes(router)

	// scheduler := &services.Scheduler{}
	// scheduler.CreateSchedulerAt(9)

	address := fmt.Sprintf(":%s", cfg.AppHost)
	router.Run(address)
}
