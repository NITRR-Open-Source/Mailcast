package main

import (
	"context"
	mq "email_app/conn"
	"email_app/helpers"
	"email_app/producers"
	"email_app/router"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.FailOnError(err, "Error loading .env file")

	// loading the url from .env
	qUrl := os.Getenv("RABBITMQ_URL")

	// connecting to the rabbitmq
	qConn, err := mq.QueueConn(qUrl)
	helpers.FailOnError(err, "Error connecting to the rabbit mq")
	defer qConn.Close()

	// initializing the message queue
	broadcastQ, broadcastCh, err := mq.MqInit(qConn, "broadcast_emails")
	helpers.FailOnError(err, "Failed to initialize the message queue")
	defer broadcastCh.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// example of sending message to the queue
	message, err := producers.Produce(broadcastCh, ctx, broadcastQ, "Message")
	if err != nil {
		log.Println("Error pushing the message:", err)
		return
	}
	log.Println(message)

	app := fiber.New()

	router.HelloRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
