package services

import (
	"fmt"
	"os"
	"plant-care-app/models"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
}

func NewEmailService() *EmailService {
	host := os.Getenv("SMTP_HOST")
	port := 587 // default SMTP port
	username := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")

	dialer := gomail.NewDialer(host, port, username, password)

	return &EmailService{
		dialer: dialer,
	}
}

func (s *EmailService) SendWateringReminder(user models.User, plants []models.Plant) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_FROM"))
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "🌱 Plants Need Watering!")

	// Create email body
	body := fmt.Sprintf(`
		<h2>Hello %s!</h2>
		<p>The following plants need watering:</p>
		<ul>
	`, user.Name)

	for _, plant := range plants {
		body += fmt.Sprintf(`
			<li>
				<strong>%s</strong><br>
				Last watered: %s<br>
				Watering interval: %d days
			</li>
		`, plant.Name, plant.LastTimeWatering.Format("January 2, 2006"), plant.WateringInterval)
	}

	body += `
		</ul>
		<p>Don't forget to water your plants to keep them healthy! 🌿</p>
	`

	m.SetBody("text/html", body)

	if err := s.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
