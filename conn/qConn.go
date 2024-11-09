package conn

import (
	"mailcast/helpers"

	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueConn(qUrl string) *amqp.Connection {
	qConn, err := amqp.Dial(qUrl)
	helpers.FailOnError(err, "Error connecting to the rabbit mq")
	return qConn
}
