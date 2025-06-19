package controllers

import (
	"net/http"
	"plant-care-app/notification-service/internal/services"

	"github.com/gin-gonic/gin"
)

type EmailRequestContent struct {
	Destination string `json:"destination" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

func SendEmail(c *gin.Context) {
	var emailRequestContent EmailRequestContent
	if err := c.ShouldBindJSON(&emailRequestContent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailService := services.NewEmailService()
	err := emailService.SendEmail(emailRequestContent.Destination, emailRequestContent.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email was sent"})
}
