package controllers

import (
	"fmt"
	"net/http"
	"plant-care-app/notification-service/internal/services"

	"github.com/gin-gonic/gin"
)

type RequestNotificationInput struct {
	Token string
	Title string
	Body  string
}

func SendDeviceNotification(c *gin.Context) {
	var requestNotificationInput RequestNotificationInput

	if err := c.ShouldBindJSON(&requestNotificationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	client, err := services.GetMessagingClient()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	fmt.Println("Messaging Client is Ready")

	//TODO: Get user token from user

	token := requestNotificationInput.Token

	services.SendPushNotification(client, token, requestNotificationInput.Title, requestNotificationInput.Body)

}
