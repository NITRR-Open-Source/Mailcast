package main

import (
	"log"
	conn "mailcast/conn"
	"mailcast/db"
	"mailcast/helpers"
	"mailcast/router"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.FailOnError(err, "Error loading .env file")

	// loading the url from .env
	qUrl := os.Getenv("RABBITMQ_URL")
	dbUrl := os.Getenv("DB_URL")

	// connecting to the rabbitmq
	qConn := conn.QueueConn(qUrl)
	defer qConn.Close()
	log.Println("Connected to rabbitmq")

	// connecting to the db
	db.DBConn(dbUrl)
	log.Println("Connected to db")

	// applying migration
	// log.Println("Applying migration to the db")
	// db.Migrate()

	// initializing the message queue
	// broadcastQ, broadcastCh := conn.MqInit(qConn, "broadcast_emails")
	// defer broadcastCh.Close()

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// // example of sending message to the queue
	// message, err := producers.Produce(broadcastCh, ctx, broadcastQ, "Message")
	// if err != nil {
	// 	log.Println("Error pushing the message:", err)
	// 	return
	// }
	// log.Println(message)

	app := fiber.New()

	router.HelloRoutes(app)
	router.SubscribeRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
