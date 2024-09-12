package producers

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Produce(ch *amqp.Channel, ctx context.Context, q amqp.Queue, body string) (string, error) {
	err := ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	if err != nil {
		return "", err
	}

	message := fmt.Sprintf("[x] Sent %s", body)
	return message, nil
}
