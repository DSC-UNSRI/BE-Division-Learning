package controllers

import (
	"pert12/database"
	"pert12/models"
	"pert12/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	res := models.User{
		Name:       user.Name,
		Email:      user.Email,
		Role:       user.Role,
		ProfilePic: user.ProfilePic,
	}

	return c.Status(200).JSON(fiber.Map{
		"user": res,
	})
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

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
	})
}
