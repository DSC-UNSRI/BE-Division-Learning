package controllers

import (
	"pert12/database"
	"pert12/models"
	"pert12/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMe(c *fiber.Ctx) error {
	var user models.User 
	userID := c.Locals("user_id").(int)

	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(200).JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	var user models.User 
	userID := c.Locals("user_id").(int)

	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "User not found"})
	}

	updates := make(map[string]interface{})

	name := c.FormValue("name")
	if name != "" {
		updates["name"] = name
	}

	password := c.FormValue("password")
	if password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Failed to hash password"})
		}
		updates["password"] = string(hashedPassword)
	}

	if _, err := c.FormFile("profile_picture"); err == nil {
		filePath, err := utils.SaveFile(c, "profile_picture", "profile")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"message": err.Error()})
		}
		updates["profile_picture"] = filePath
	}

	if len(updates) > 0 {
		if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{"message": "Could not update profile"})
		}
	}	

	return c.Status(200).JSON(fiber.Map{
		"message": "Profile updated successfully",
	})
}

