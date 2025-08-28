package controllers

import (
	"be_pertemuan10/database"
	"be_pertemuan10/model"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	users := []model.User{}
	if err := database.DB.Find(&users).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "tidak dapat mengambil data users",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"users": users,
	})
}

func GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := model.User{}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "tidak dapat mengambil data users",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := model.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	database.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"user": "user terbuat",
	})
}

func PatchUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := model.User{}
	newUser := model.User{}

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "data tidak valid",
		})
	}

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "tidak dapat mengambil data users",
		})
	}

	user.Name = newUser.Name

	database.DB.Save(&user)

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}

func DeleteUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user := model.User{}
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "tidak dapat mengambil data users",
		})
	}

	database.DB.Delete(user)

	return c.Status(200).JSON(fiber.Map{
		"user": user,
	})
}
