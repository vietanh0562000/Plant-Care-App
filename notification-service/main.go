package main

import (
	"fmt"
	"plant-care-app/notification-service/config"
	"plant-care-app/notification-service/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetInstance()

	router := gin.Default()
	routes.SetupRoutes(router)

	address := fmt.Sprintf(":%s", cfg.GetAppPort())
	router.Run(address)
}
