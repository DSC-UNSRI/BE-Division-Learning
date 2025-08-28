package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role").(string)
	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden - Admin Only",
		})
	}

	return c.Next()
}