package rabbitmq

import (
	"log"

	"github.com/irfhakeem/go-fiber-clean-starter/config"
	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var ch *amqp.Channel

func InitRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	conn, ch = config.ConnectRabbitMQ()
	if conn == nil || ch == nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}

	return conn, ch
}

func DeclareQueue(queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}
	return q
}
