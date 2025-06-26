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

func SendWateringNotification() {
	var plants []models.Plant
	result := database.DB.Find(&plants).Where("last_time_watering + (watering_interval || ' days) <= NOW()")
	if result.Error != nil {
		log.Printf("Database error: %v\n", result.Error)
		return
	}

	cfg := config.GetInstance()

	for i := 0; i < len(plants); i++ {
		userId := plants[i].UserID

		content := "Your plants need water"

		trySendEmailToUser(userId, content, cfg)

		trySendDeviceNotificationToUser(userId, content, cfg)
	}

	fmt.Println("---------------------------------SE")
}

func trySendEmailToUser(userId int, content string, cfg *config.Config) {
	var request = fmt.Sprintf("%s/user_email/%v", cfg.GetUserServiceHost(), userId)
	response, err := http.Get(request)

	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Printf("Error:%s", err.Error())
		response.Body.Close()
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Printf("Error:%s", err.Error())
		response.Body.Close()
		return
	}
	userEmail := data["email"].(string)
	sendRequestEmailInternal(userEmail, content, cfg)
}

func sendRequestEmailInternal(destination, content string, cfg *config.Config) {
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

func trySendDeviceNotificationToUser(userId int, content string, cfg *config.Config) {
	var request = fmt.Sprintf("%s/users/%v/firebaseToken", cfg.GetUserServiceHost(), userId)
	response, err := http.Get(request)

	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Printf("Error:%s", err.Error())
		response.Body.Close()
		return
	}

	var data map[string]interface{}
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		fmt.Printf("Error:%s", err.Error())
		response.Body.Close()
		return
	}
	firebaseToken := data["Token"].(string)
	sendDeviceNotification(firebaseToken, "Watering", content, cfg)
}

func sendDeviceNotification(destination, title, body string, cfg *config.Config) {
	// Create data to send
	payload := map[string]string{
		"Token": destination,
		"Title": title,
		"Body":  body,
	}

	// Encode as JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	var url = fmt.Sprintf("%s/send_device_notification", cfg.GetNotificationServiceHost())
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
