package service

import (
	"context"
	"go-rabbitmq/src/config"
	"go-rabbitmq/src/error"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ProducerMsg(bodyMessage []byte) {
	ch, err := config.RmqConnection()

	if err != nil {
		error.FailOnError(err, "Failed to connect to RabbitMQ")
		return
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"monitoring",
		true,
		false,
		false,
		false,
		nil,
	)

	error.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(
		ctx,
		"amq.direct",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bodyMessage,
		},
	)

	error.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s to RMQ", bodyMessage)
}
