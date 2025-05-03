package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	url := os.Getenv("RABBITMQ_URL")
	if url == "" {
		url = "amqp://guest:guest@localhost:5672/"
	}

	log.Printf("Connecting to RabbitMQ at %s", url)

	conn, err := amqp.Dial(url)
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
