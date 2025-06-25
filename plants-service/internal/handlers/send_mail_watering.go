package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"plant-care-app/plants-service/config"
	"plant-care-app/plants-service/internal/database"
	"plant-care-app/plants-service/internal/models"
)

func SendMailWatering() {
	var plants []models.Plant
	result := database.DB.Find(&plants).Where("last_time_watering + (watering_interval || ' days) <= NOW()")
	if result.Error != nil {
		log.Printf("Database error: %v\n", result.Error)
		return
	}

	cfg := config.GetInstance()

	for i := 0; i < len(plants); i++ {
		userId := plants[i].UserID
		var request = fmt.Sprintf("%s/user_email/%v", cfg.GetUserServiceHost(), userId)
		response, err := http.Get(request)
		fmt.Println(response)
		if err != nil || response.StatusCode != http.StatusOK {
			response.Body.Close()
			continue
		}

		fmt.Println(response)

		var data map[string]interface{}
		if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
			response.Body.Close()
			continue
		}
		userEmail := data["email"].(string)
		content := "Let's water your plant"
		SendRequestEmailInternal(userEmail, content, cfg)
	}
}

func SendRequestEmailInternal(destination, content string, cfg *config.Config) {
	fmt.Println(destination)
	// Create data to send
	payload := map[string]string{
		"Destination": destination,
		"Content":     content,
	}

	// Encode as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	var url = fmt.Sprintf("%s/send_email", cfg.GetNotificationServiceHost())
	fmt.Println(url)
	// Send POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check response
	fmt.Println("Status:", resp.Status)
}
