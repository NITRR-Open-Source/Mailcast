package conn

import (
	"mailcast/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func MqInit(qConn *amqp.Connection, qName string) (amqp.Queue, *amqp.Channel) {
	ch, err := qConn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")

	q, err := ch.QueueDeclare(
		qName, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	helpers.FailOnError(err, "Failed to initialize the message queue")

	return q, ch
}
