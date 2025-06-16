package database

import (
	"plant-care-app/user-service/internal/models"
)

func SeedData() {

	// Create a test user if not exists
	user := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123", // In real app, this should be hashed
		IsAdmin:  true,
	}
	DB.FirstOrCreate(&user, models.User{Email: user.Email})
}
