package services

import (
	"fmt"
	"plant-care-app/notification-service/config"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	dialer *gomail.Dialer
}

func NewEmailService() *EmailService {
	cfg := config.GetInstance()
	host := cfg.GetSMTPHost()
	port := cfg.GetSMTPPort()
	username := cfg.GetSMTPUser()
	password := cfg.GetSMTPPassword()

	dialer := gomail.NewDialer(host, port, username, password)

	return &EmailService{
		dialer: dialer,
	}
}

func (s *EmailService) SendEmail(toEmail string, content string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.dialer.Username)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", "🌱 Plants Need Watering!")

	// Create email body
	body := fmt.Sprintf(`
		<h2>Hello %s!</h2>
		<p>The following plants need watering:</p>
		<ul>
	`, toEmail)

	body += fmt.Sprintf(`
		</ul>
		<p>Don't forget to water your plants to keep them healthy! 🌿</p>
		<p>%s</p>
	`, content)

	m.SetBody("text/html", body)

	if err := s.dialer.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	fmt.Println(m)
	fmt.Println(s.dialer.Host)
	fmt.Println(s.dialer.Password)
	fmt.Println(s.dialer.Port)

	return nil
}
