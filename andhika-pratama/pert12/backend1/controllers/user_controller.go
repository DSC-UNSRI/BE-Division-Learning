package controllers

import (
	"pert12/database"
	"pert12/models"

	"github.com/gofiber/fiber/v2"
)

func GetMe(c *fiber.Ctx) error {
	var user models.User 
	userID := c.Locals("user_id").(int)

	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(user)
}