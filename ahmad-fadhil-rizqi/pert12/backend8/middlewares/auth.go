package middlewares

import (
	"backend8/utils"
	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthRequired() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		claims, err := utils.ParseJWT(token, os.Getenv("JWT_SECRET"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
		}
		c.Locals("user_id", claims.UserID)
		c.Locals("role", claims.Role)
		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok || role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "forbidden"})
		}
		return c.Next()
	}
}
