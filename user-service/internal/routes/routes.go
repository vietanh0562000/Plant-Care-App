package routes

import (
	"plant-care-app/user-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.GET("/user:userID", handlers.GetUser)
	r.GET("/user_email/:id", handlers.GetUserEmail)
	r.GET("/users/:id/firebaseToken", handlers.GetFireBaseToken)
}
