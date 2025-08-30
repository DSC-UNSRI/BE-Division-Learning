package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: Admins only",
			})
		}
		return c.Next()
	}
}
