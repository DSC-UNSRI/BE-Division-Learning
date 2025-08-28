package controllers

import (
	"backend/models"
	"backend/database"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.Status(200).JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("id")

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	name := c.FormValue("name")
	password := c.FormValue("password")
    profilePicture, err := c.FormFile("profilePicture")

	if name != "" {
        user.Name = name
    }

	if password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = string(hashed)
	}

	if err == nil && profilePicture != nil {
		if profilePicture.Size > 1*1024*1024 {
			return c.Status(400).JSON(fiber.Map{"error": "Ukuran foto maksimal 1MB"})
		}

		ext := filepath.Ext(profilePicture.Filename)
		allowedExt := map[string]bool{
			".png":  true,
			".jpg":  true,
			".jpeg": true,
		}
		if !allowedExt[ext] {
			return c.Status(400).JSON(fiber.Map{"error": "Foto harus berformat PNG, JPG, atau JPEG"})
		}

		if user.ProfilePicture != "" && !strings.Contains(user.ProfilePicture, "pravatar.cc") {
			oldPath := "." + strings.TrimPrefix(user.ProfilePicture, "http://127.0.0.1:3000")
			_ = os.Remove(oldPath)
		}
		
        filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
        savePath := "./assets/" + filename
        if err := c.SaveFile(profilePicture, savePath); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Failed to save photo"})
        }
        user.ProfilePicture = "http://127.0.0.1:3000/assets/" + filename
    }

	if err := database.DB.Save(&user).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

	return c.Status(201).JSON(fiber.Map{
		"message": "Profile berhasil diupdate",
		"user":    user,
	})
}