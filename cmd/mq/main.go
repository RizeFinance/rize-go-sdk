package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/rizefinance/rize-go-sdk/internal"
	"github.com/rizefinance/rize-go-sdk/mq"
)

func init() {
	// Load local env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}

func main() {
	config := mq.Config{
		Username:    internal.CheckEnvVariable("mq_username"),
		Password:    internal.CheckEnvVariable("mq_password"),
		ClientID:    internal.CheckEnvVariable("mq_client_id"),
		Environment: internal.CheckEnvVariable("environment"),
		Debug:       true,
	}

	// Create new Rize MQ client
	rc, err := mq.NewClient(&config)
	if err != nil {
		log.Fatal("Error building MQ client\n", err)
	}

	// Create MQ connection
	if err := rc.MessageQueue.Connect(); err != nil {
		log.Fatal("Error creating MQ connection:\n", err)
	}

	// Subscribe to customer MQ topics
	sub, err := rc.MessageQueue.Subscribe("transfer", "transferSubscription")
	if err != nil {
		log.Printf("Subscribe failed: %s\n", err)
	}

	// Unsubscribe from a subscription
	if err := rc.MessageQueue.Unsubscribe(sub); err != nil {
		log.Printf("Unsubscribe failed: %s\n", err)
	}
}
