package controllers

import (
	"backend/database"
	"backend/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
    password := c.FormValue("password")
	confirmPassword := c.FormValue("confirmPassword")

	if name == "" || email == "" || password == "" || confirmPassword == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "name, email, password, and confirm password are required",
        })
    }

	if password != confirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "password and confirm password do not match",
		})
	}

	var existingUser models.User
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "email already registered",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to hash password"})
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
		ProfilePicture: "https://i.pravatar.cc/150",
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	user.Password = ""
	return c.Status(201).JSON(fiber.Map{
		"message": "Register success",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error{
	email := c.FormValue("email")
    password := c.FormValue("password")


	if email == "" || password == "" {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "email and password are required",
        })
    }

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"Message": "Kredensial tidak valid"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"Message": "Kredensial tidak valid"})
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.Name,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not login"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    signed,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Login success",
		"user": fiber.Map{
			"id":   user.ID,
			"role": user.Role,
		},
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
	})

	return c.JSON(fiber.Map{
		"message": "Logout success",
	})
}