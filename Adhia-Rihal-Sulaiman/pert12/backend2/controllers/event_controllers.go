package controllers

import (
	"backend2/database"
	"backend2/models"
	"backend2/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GetProgramStudi handles fetching a list of program studi
func GetProgramStudi(c *fiber.Ctx) error {
	jurusanQuery := c.Query("jurusan")
	limit := c.Query("limit")

	cacheKey := "programstudi_all"
	if jurusanQuery != "" {
		cacheKey += "_" + jurusanQuery
	}
	if limit != "" {
		cacheKey += "_" + limit
	}

	// Anda perlu menyesuaikan metode cache di sini
	if cached, found := cache.GetProgramStudiCache(cacheKey); found {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"programstudi": cached,
			"cache":        true,
		})
	}

	var programStudi []models.ProgramStudi
	query := database.DB.Model(&models.ProgramStudi{}).Where("posted_at <= ?", time.Now()).Order("posted_at DESC")

	if jurusanQuery != "" {
		query = query.Where("jurusan ILIKE ?", "%"+jurusanQuery+"%")
	}

	if limit != "" {
		if limitInt, err := strconv.Atoi(limit); err == nil {
			query = query.Limit(limitInt)
		}
	}

	if err := query.Find(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Program studi not found",
		})
	}

	// Anda perlu menyesuaikan metode cache di sini
	cache.SetProgramStudiCache(cacheKey, programStudi)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"programstudi": programStudi,
		"cache":        false,
	})
}

// CreateProgramStudi handles creating a new program studi
func CreateProgramStudi(c *fiber.Ctx) error {
	var programStudi models.ProgramStudi

	jurusan := c.FormValue("jurusan")
	deskripsi := c.FormValue("deskripsi")
	akreditasi := c.FormValue("akreditasi")
	jenjang := c.FormValue("jenjang")

	// **PERBAIKAN DI SINI:** Menambahkan argumen boolean keempat (isCover)
	// true for isProgramStudi, true for isCover (as it's a logo for programstudi)
	logo, err := utils.SaveFile(c, "logo", true, true)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	if jurusan == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Jurusan is required",
		})
	}

	if postedAtStr := c.FormValue("posted_at"); postedAtStr != "" {
		t, err := time.Parse(time.RFC3339, postedAtStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid posted_at format. Use ISO 8601 (RFC3339)",
			})
		}
		programStudi.PostedAt = t
	} else {
		programStudi.PostedAt = time.Now()
	}

	programStudi.Jurusan = jurusan
	programStudi.Deskripsi = deskripsi
	programStudi.Akreditasi = akreditasi
	programStudi.Jenjang = jenjang
	programStudi.Logo = logo

	if err := database.DB.Create(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to save program studi",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "Program studi created successfully",
		"programstudi": programStudi,
	})
}

// UpdateProgramStudi handles updating an existing program studi
func UpdateProgramStudi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID program studi is required",
		})
	}
	var programStudi models.ProgramStudi

	if err := database.DB.Where("id =?", id).First(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Program studi not found",
		})
	}

	if jurusan := c.FormValue("jurusan"); jurusan != "" {
		programStudi.Jurusan = jurusan
	}
	if deskripsi := c.FormValue("deskripsi"); deskripsi != "" {
		programStudi.Deskripsi = deskripsi
	}
	if akreditasi := c.FormValue("akreditasi"); akreditasi != "" {
		programStudi.Akreditasi = akreditasi
	}
	if jenjang := c.FormValue("jenjang"); jenjang != "" {
		programStudi.Jenjang = jenjang
	}

	if _, err := c.FormFile("logo"); err == nil {
		// **PERBAIKAN DI SINI:** Menambahkan argumen boolean keempat (isCover)
		// true for isProgramStudi, true for isCover (as it's a logo for programstudi)
		logo, err := utils.SaveFile(c, "logo", true, true)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		programStudi.Logo = logo
	}

	if postedAtStr := c.FormValue("posted_at"); postedAtStr != "" {
		t, err := time.Parse(time.RFC3339, postedAtStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid posted_at format. Use ISO 8601 (RFC3339)",
			})
		}
		programStudi.PostedAt = t
	}

	if err := database.DB.Save(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update program studi",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Program studi updated successfully",
		"programstudi": programStudi,
	})
}

// DeleteProgramStudi handles deleting a program studi
func DeleteProgramStudi(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID program studi is required",
		})
	}
	var programStudi models.ProgramStudi

	if err := database.DB.Where("id =?", id).First(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Program studi not found",
		})
	}

	if err := database.DB.Delete(&programStudi).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete program studi",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":      "Program studi deleted successfully",
		"programstudi": programStudi,
	})
}