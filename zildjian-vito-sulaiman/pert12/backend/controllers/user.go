package controllers

import (
	"backend/config"
	"backend/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func GetMe(c *fiber.Ctx) error {
	sess, err := c.Locals("session").(*session.Store).Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Something went wrong"})
	}

	userID := sess.Get("userID")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	var user models.User
	config.DB.First(&user, userID)

	return c.JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if config.DB.First(&user, id).Error != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid request"})
	}

	config.DB.Save(&user)

	return c.JSON(fiber.Map{"message": "User updated successfully", "user": user})
}
