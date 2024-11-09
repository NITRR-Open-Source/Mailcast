package subscribe

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func Subscribe(c *fiber.Ctx) error {
	body := new(Body)
	c.BodyParser(body)
	fmt.Println(body)
	return c.JSON(fiber.Map{"status": "success", "message": "Hello i'm ok!", "data": nil})
}
