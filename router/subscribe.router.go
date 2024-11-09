package router

import (
	"mailcast/handlers/subscribe"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SubscribeRoutes(app *fiber.App) {
	subscribeGrp := app.Group("/subscribe", logger.New())
	subscribeGrp.Post("/", subscribe.Subscribe)
}
