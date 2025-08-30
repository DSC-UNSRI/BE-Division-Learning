
package middleware

import (
	"backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func Authenticated(c *fiber.Ctx) error {
	sess, err := c.Locals("session").(*session.Store).Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong"})
	}

	userID := sess.Get("userID")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.User
	config.DB.First(&user, userID)

	if user.ID == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	c.Locals("user", user)
	return c.Next()
}
