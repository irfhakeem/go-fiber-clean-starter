package rabbitmq

import (
	"log"

	"github.com/irfhakeem/go-fiber-clean-starter/rabbitmq/consumers"
)

func ConsumeAll() {
	log.Println("Starting all consumers...")

	ConsumeEmail("email_queue", consumers.HandleVerificationEmail)
}

func ConsumeEmail(queueName string, handler func([]byte)) {
	q := DeclareQueue(queueName)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	for msg := range msgs {
		handler(msg.Body)
	}
}
