package handlers

import (
	"mailcast/db"
	"mailcast/models"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	// sub := models.Subscriber{EmailID: "email", SubscribedAt: time.Now()}
	sub := models.Subscriber{}
	tx := db.Db.Create(&sub)
	println(tx)
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
