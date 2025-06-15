package services

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/gomail.v2"
)

type MailService struct {
}

func (mailSrv *MailService) SendNewMail(message string, destination string) {
	// Set host, port, username, password
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := 587
	smtpUser := os.Getenv("SMTP_USER")
	smtpPassword := os.Getenv("SMTP_PASSWORD")

	if smtpHost == "" || smtpUser == "" || smtpPassword == "" {
		log.Println("Invalid email config")
		return
	}

	newDialer := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPassword)

	newMessage := gomail.NewMessage()
	newMessage.SetHeader("From", smtpUser)
	newMessage.SetHeader("To", destination)
	newMessage.SetHeader("Subject", "Plant need your care")

	newMessage.SetBody("text/plain", "Go to app and water your plant")

	if err := newDialer.DialAndSend(newMessage); err != nil {
		fmt.Println("Fail to send message to user")
		return
	}

	fmt.Println("Send email succesfully")
}
