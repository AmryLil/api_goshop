package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"google.golang.org/api/option"
)

func InitFirebase() *storage.Client {
	opt := option.WithCredentialsFile("goshop-be94c-firebase-adminsdk-xyt39-4a1194ba92.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase app: %v", err)
	}
	client, err := app.Storage(context.Background())
	if err != nil {
		log.Fatalf("Error initializing Firebase storage client: %v", err)
	}
	return client
}
