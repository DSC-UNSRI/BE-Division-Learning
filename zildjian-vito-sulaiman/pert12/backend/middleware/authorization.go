
package middleware

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func Admin(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	if user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
	}

	return c.Next()
}

func ProfileUpdate(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid ID"})
	}

	if user.ID != uint(id) && user.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "Forbidden"})
	}

	return c.Next()
}
