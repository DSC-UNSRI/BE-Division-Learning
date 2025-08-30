package controllers

import (
	"fmt"
	"pert12/database"
	"pert12/models"
	"pert12/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMe(c *fiber.Ctx) error {
	userID := c.Locals("user_id") // di-set oleh JWT middleware
	var user models.User

	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"profilePic": user.ProfilePic,
		},
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	userID := c.Locals("user_id")
	if fmt.Sprint(userID) != c.Params("id") {
		return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
	}

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	if name := c.FormValue("name"); name != "" {
		user.Name = name
	}

	if email := c.FormValue("email"); email != "" {
		user.Email = email
	}

	if password := c.FormValue("password"); password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
		}
		user.Password = string(hashedPassword)
	}

	if _, err := c.FormFile("profile_pic"); err == nil {
		profilepic, err := utils.SaveFile(c, "profile_pic", true)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Failed to Update Profile",
			})
		}
		user.ProfilePic = profilepic
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to update user profile",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "User Profile updated successfully",
		"user": fiber.Map{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"profilePic": user.ProfilePic,
		},
	})
}
