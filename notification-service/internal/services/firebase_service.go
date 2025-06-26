package services

import (
	"context"
	"fmt"
	"log"
	"plant-care-app/notification-service/config"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func GetMessagingClient() (*messaging.Client, error) {
	if firebaseApp == nil {
		var err error
		firebaseApp, err = InitFirebaseApp()
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}

	return firebaseApp.Messaging(context.Background())
}

func InitFirebaseApp() (*firebase.App, error) {
	cfg := config.GetInstance()
	opt := option.WithCredentialsFile(cfg.GetGGAppCredentailsPath()) // Path to your file
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func SendPushNotification(client *messaging.Client, token, title, body string) error {
	message := &messaging.Message{
		Token: token,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		// Optional: Add custom data
		// Data: map[string]string{"key1": "value1"},
	}

	response, err := client.Send(context.Background(), message)
	if err != nil {
		return err
	}

	log.Println("Successfully sent message:", response)
	return nil
}
