package mq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MqInit(qConn *amqp.Connection, qName string) (queue amqp.Queue, channel *amqp.Channel, err error) {
	ch, err := qConn.Channel()
	if err != nil {
		log.Panicf("Failed to open a channel: %s", err)
	}

	q, err := ch.QueueDeclare(
		qName, // name
		true,  // durable
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
