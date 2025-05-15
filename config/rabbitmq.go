package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

type RabbitMQConfig struct {
	Username string
	Password string
}

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("RABBITMQ_URL is not set")
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return conn, ch
}

func CloseConnectionRabbitMQ(conn *amqp.Connection, ch *amqp.Channel) {
	if ch != nil {
		if err := ch.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ channel: %v", err)
		}
	}

	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Printf("Failed to close RabbitMQ connection: %v", err)
		}
	}
}
