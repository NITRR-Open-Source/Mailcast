package mq

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func QueueConn(qUrl string) (qConn *amqp.Connection, err error) {
	qConn, err = amqp.Dial(qUrl)
	if err != nil {
		return nil, err
	}
	return qConn, nil
}
