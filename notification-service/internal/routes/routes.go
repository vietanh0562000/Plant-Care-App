package routes

import (
	controllers "plant-care-app/notification-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/send_email", controllers.SendEmail)
}
