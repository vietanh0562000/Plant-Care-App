package controllers

import (
	"plant-care-app/database"
	"plant-care-app/models"
	"time"
)

func SeedData() {
	// Create some species
	species := []models.Species{
		{Name: "Monstera Deliciosa", Notes: "Popular houseplant with distinctive leaf holes"},
		{Name: "Snake Plant", Notes: "Low maintenance, air-purifying plant"},
		{Name: "Peace Lily", Notes: "Beautiful flowering plant that helps clean indoor air"},
		{Name: "Fiddle Leaf Fig", Notes: "Trendy indoor tree with large, glossy leaves"},
		{Name: "ZZ Plant", Notes: "Drought-tolerant plant with glossy leaves"},
	}

	for _, s := range species {
		database.DB.FirstOrCreate(&s, models.Species{Name: s.Name})
	}

	// Create a test user if not exists
	user := models.User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "password123", // In real app, this should be hashed
		IsAdmin:  true,
	}
	database.DB.FirstOrCreate(&user, models.User{Email: user.Email})

	// Create plants with different watering schedules
	now := time.Now()
	plants := []models.Plant{
		{
			Name:             "Big Monstera",
			ImagePath:        "uploads/plants/monstera1.jpg",
			WateringInterval: 7,
			LastTimeWatering: now.AddDate(0, 0, -8), // Needs watering
			UserID:           user.ID,
			SpeciesID:        1,
		},
		{
			Name:             "Snake Plant Corner",
			ImagePath:        "uploads/plants/snake1.jpg",
			WateringInterval: 14,
			LastTimeWatering: now.AddDate(0, 0, -10), // Needs watering
			UserID:           user.ID,
			SpeciesID:        2,
		},
		{
			Name:             "Peace Lily Office",
			ImagePath:        "uploads/plants/peace1.jpg",
			WateringInterval: 5,
			LastTimeWatering: now.AddDate(0, 0, -3), // Recently watered
			UserID:           user.ID,
			SpeciesID:        3,
		},
		{
			Name:             "Fiddle Leaf Living Room",
			ImagePath:        "uploads/plants/fiddle1.jpg",
			WateringInterval: 10,
			LastTimeWatering: now.AddDate(0, 0, -5), // Needs watering soon
			UserID:           user.ID,
			SpeciesID:        4,
		},
		{
			Name:             "ZZ Plant Bedroom",
			ImagePath:        "uploads/plants/zz1.jpg",
			WateringInterval: 14,
			LastTimeWatering: now.AddDate(0, 0, -7), // Recently watered
			UserID:           user.ID,
			SpeciesID:        5,
		},
	}

	for _, p := range plants {
		database.DB.FirstOrCreate(&p, models.Plant{Name: p.Name, UserID: user.ID})
	}
}
