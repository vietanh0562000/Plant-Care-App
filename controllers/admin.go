package controllers

import (
	"log"
	"os"
	"plant-care-app/database"
	"plant-care-app/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateFirstAdmin() {
	// Find admin account in data base
	var adminCount int64
	database.DB.Model(&models.User{}).Where("is_admin = ?", true).Count(&adminCount)

	if adminCount > 0 {
		log.Println("Admin already exist")
		return
	}

	// IF NOT create the first admin account

	// * Get admin credential from .env
	adminName := os.Getenv("ADMIN_NAME")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	adminEmail := os.Getenv("ADMIN_EMAIL")

	if adminName == "" || adminPassword == "" || adminEmail == "" {
		log.Println("Admin credentials not set in env")
		return
	}

	// * Create admin account
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(adminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing admin password: %v", err.Error())
		return
	}

	newAdmin := models.User{
		Name:     adminName,
		Password: string(hashPassword),
		Email:    adminEmail,
		IsAdmin:  true,
	}

	if err := database.DB.Create(&newAdmin).Error; err != nil {
		log.Printf("Error creating admin: %v", err)
		return
	}

	log.Println("Admin account created successfully")

}
