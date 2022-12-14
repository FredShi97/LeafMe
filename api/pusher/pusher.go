package pusher

import (
	"errors"
	"github.com/2J/LeafMe/api/models"
	"github.com/oliveroneill/exponent-server-sdk-golang/sdk"
	log "github.com/sirupsen/logrus"
)

// Notification is for push notifications
type Notification struct {
	To    string
	Title string
	Body  string
	Data  map[string]string
}

func push(notification Notification) error {
	// To check the token is valid
	pushToken, err := expo.NewExponentPushToken(notification.To)
	if err != nil {
		return err
	}

	log.Info("Sending notification...")
	log.Info("To: " + notification.To)
	log.Info("Title: " + notification.Title)
	log.Info("Body: " + notification.Body)
	log.Info(notification.Data)

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)

	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       pushToken,
			Title:    notification.Title,
			Body:     notification.Body,
			Data:     notification.Data,
			Sound:    "default",
			Priority: expo.DefaultPriority,
		},
	)

	// Check errors
	if err != nil {
		return err
	}

	// Validate responses
	if response.ValidateResponse() != nil {
		log.Error(response.PushMessage.To, "failed")
		return errors.New("Push notification failed")
	}

	return nil
}

// PushNotification pushes a notification based on plant ID
func PushNotification(plantID int64, notification Notification) error {
	// Get plant
	var plant models.Plant
	err := plant.GetByID(plantID)
	if err != nil {
		return err
	}
	if plant.PushToken == "" {
		return errors.New("Push token not set for plant")
	}

	notification.To = plant.PushToken
	return push(notification)
}
