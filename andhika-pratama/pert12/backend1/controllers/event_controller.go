package controllers

import (
	"pert12/database"
	"pert12/models"
	"pert12/utils"

	"strconv"
	"time"

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

func UpdateEvent(c *fiber.Ctx) error {
	var event models.Event

	paramID := c.Params("id")
	ID, err := strconv.Atoi(paramID)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid event ID",
		})
	}

	if err := database.DB.First(&event, ID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Event not found",
		})
	}

	updates := make(map[string]interface{})

	location := c.FormValue("location")
	if location != "" {
		updates["location"] = location
	}
	
	dateStr := c.FormValue("start")
	if dateStr != "" {
		layout := "2006-01-02"
		parsedDate, err := time.Parse(layout, dateStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid date format. Please use YYYY-MM-DD.",
			})
		}
		updates["start"] = parsedDate
	}

	if _, err := c.FormFile("cover"); err == nil {
		filePath, err := utils.SaveFile(c, "cover", "cover")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		updates["cover"] = filePath
	}

	if len(updates) > 0 {
		if err := database.DB.Model(&event).Updates(updates).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Could not update event",
			})
		}
	}	

	return c.Status(200).JSON(fiber.Map{
		"message": "Event updated successfully",
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