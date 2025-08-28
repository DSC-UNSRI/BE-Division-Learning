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

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event

	if err := c.BodyParser(&event); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if event.Location == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Location is required",
		})
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create event",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Event created successfully",
	})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "ID event is required",
		})
	}

	var event models.Event
	if err := database.DB.Where("id =?", id).First(&event).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to delete event",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Event deleted successfully",
	})
}