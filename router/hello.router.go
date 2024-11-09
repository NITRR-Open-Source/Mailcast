package router

import (
	"mailcast/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func HelloRoutes(app *fiber.App) {
	helloGroup := app.Group("/hello", logger.New())
	helloGroup.Get("/", handlers.Hello)
}
