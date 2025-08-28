package controllers

import (
	"backend2/database"
	"backend2/models"
	"backend2/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetEvent(c *fiber.Ctx) error {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to fetch events"})
	}
	return c.Status(200).JSON(fiber.Map{"events": events})
}

func PostEvent(c *fiber.Ctx) error {
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
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	location := c.FormValue("location")

	if location != "" {
		event.Location = location
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
		event.Start = parsedDate
	}
	
	if _, err := c.FormFile("cover"); err == nil {
        filePath, err := utils.SaveFile(c, "cover", "true")
        if err != nil {
            return c.Status(400).JSON(fiber.Map{
                "message": err.Error(),
            })
        }
        event.Cover = filePath
    }

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to update event"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Event updated successfully"})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "Event not found"})
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete event"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Event deleted successfully"})
}
