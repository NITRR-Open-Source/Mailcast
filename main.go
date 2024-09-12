package main

import (
	"context"
	mq "email_app/conn"
	"email_app/helpers"
	"email_app/producers"
	"log"
	"time"
)

func main() {
	// initializing the message queue
	broadcastQ, broadcastCh, err := mq.MqInit("broadcast_emails")
	helpers.FailOnError(err, "Failed to initialize the message queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// example of sending message to the queue
	message, err := producers.Produce(broadcastCh, ctx, broadcastQ, "Message")
	if err != nil {
		log.Println("Error pushing the message:", err)
		return
	}
	log.Println(message)
}
