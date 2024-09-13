package router

import (
	"email_app/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HelloRoutes(app *fiber.App) {
	api := app.Group("/hello", logger.New())
	api.Get("/", handler.Hello)
}
