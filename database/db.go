package database

import (
	"fmt"
	"log"
	"plant-care-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user=vanh dbname=plant_care_app sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Plant{}, &models.Schedule{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database migrated successfully")

	DB = db
}
