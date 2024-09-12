package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MqInit(qName string) (queue amqp.Queue, channel *amqp.Channel, err error) {
	qconn, err := amqp.Dial("amqp://user:password@localhost:5672/")
	if err != nil {
		return amqp.Queue{}, nil, err
	}
	defer qconn.Close()

	ch, err := qconn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		qName, // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return amqp.Queue{}, nil, err
	}

	return q, ch, nil
}
