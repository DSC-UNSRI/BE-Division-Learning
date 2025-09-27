package controllers

import (
	"oauth-2_redis/database"
	"oauth-2_redis/models"
	"oauth-2_redis/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateEvent(c *fiber.Ctx) error {
	var event models.Event

	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "harap masukan lokasi dengan tepat",
		})
	}

	if err := database.DB.Create(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal membuat event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "event berhasil dibuat",
	})
}

func UpdateEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "event tidak ditemukan",
		})
	}

	if location := c.FormValue("location"); location != "" {
		event.Location = location
	}
	if start := c.FormValue("start"); start != "" {
		var t time.Time
		var err error

		t, err = time.Parse(time.RFC3339, start)
		if err != nil {
			t, err = time.Parse("2006-01-02", start)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"error": "Invalid start format. Use RFC3339 (e.g. 2025-08-26T00:00:00Z) or YYYY-MM-DD",
				})
			}
		}

		event.Start = t
	}

	if _, err := c.FormFile("cover"); err == nil {
		file, err := utils.SaveFile(c, "cover")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "harap masukkan gambar valid",
			})
		}
		event.Cover = file
	}

	if err := database.DB.Save(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mengupdate event",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "event berhasil diupdate",
	})
}

func DeleteEvent(c *fiber.Ctx) error {
	id := c.Params("id")
	var event models.Event

	if err := database.DB.Where("id = ?", id).First(&event).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "event tidak ditemukan",
		})
	}

	if err := database.DB.Delete(&event).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal menghapus event",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "event berhasil dihapus",
	})
}

func GetEvent(c *fiber.Ctx) error {
	var events []models.Event
	if err := database.DB.Find(&events).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "belum ada jadwal nonton bareng",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"events": events,
	})
}
