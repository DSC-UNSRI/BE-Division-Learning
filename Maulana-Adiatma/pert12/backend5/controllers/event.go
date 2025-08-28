package controllers

import (
	"pert12/database"
	"pert12/models"
	"pert12/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostEvent(c *fiber.Ctx) error {
	var event models.Event

	cover, err := utils.SaveFile(c, "cover", true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to save cover image",
		})
	}
	location := c.FormValue("location")

	event.Cover = cover

	event.Location = location

	if startPostStr := c.FormValue("startpost"); startPostStr != "" {
		t, err := time.Parse(time.RFC3339Nano, startPostStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid startpost format. Use ISO 8601 (RFC3339Nano)",
			})
		}
		event.Start = t
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to post event",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Success to post event",
	})
}

func GetEvent(c *fiber.Ctx) error {
	var event []models.Event

	limit := c.Query("limit")
	query := database.DB.Model(&models.Event{}).Where("start <= ?", time.Now()).Order("start DESC")

	if limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(limitInt)
		}
	}

	if err := query.Find(&event).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to get event",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"event": event,
	})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "event not found",
		})
	}

	if _, err := c.FormFile("cover"); err == nil {
		cover, err := utils.SaveFile(c, "cover", true)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to Update cover image",
			})
		}
		event.Cover = cover
	}

	if location := c.FormValue("location"); location != "" {
		event.Location = location
	}

	if startPostStr := c.FormValue("startpost"); startPostStr != "" {
		t, err := time.Parse(time.RFC3339Nano, startPostStr)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error": "Invalid startpost format. Use ISO 8601 (RFC3339Nano)",
			})
		}
		event.Start = t
	}

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update event",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "event updated successfully",
		"event":  event,
	})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.First(&event, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "event not found",
		})
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to delete event",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "event deleted successfully",
	})
}
