package main

import (
	"bytes"
	"log"
	"mailcast/helpers"
	"os"
	"time"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	err := godotenv.Load()
	helpers.FailOnError(err, "Error loading .env file")

	qUrl := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(qUrl)
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	helpers.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"broadcast_emails", // queue
		"email",            // consumer
		false,              // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	helpers.FailOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
		dotCount := bytes.Count(d.Body, []byte("."))
		t := time.Duration(dotCount)
		time.Sleep(t * time.Second)

		// Manually acknowledge the message
		if err := d.Ack(false); err != nil {
			log.Printf("Failed to acknowledge message: %v", err)
		}

		log.Printf("Done")
	}

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
}
