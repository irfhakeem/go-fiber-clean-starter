package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func PublishEmail(queueName string, body []byte) {
	q := DeclareQueue(queueName)

	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish a message: %v", err)
	}
	log.Printf("email sent!")
}
