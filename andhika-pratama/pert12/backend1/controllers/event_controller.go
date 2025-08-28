package controllers

import (
	"pert12/database"
	"pert12/models"

	"github.com/gofiber/fiber/v2"
)

func GetEvents(c *fiber.Ctx) error {
	var events []models.Event

	if err := database.DB.Find(&events).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to retrieve events",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"events": events,
	})
}