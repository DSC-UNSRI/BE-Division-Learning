package controllers

import (
	"pert12/database"
	"pert12/models"

	"os"
	"time"
	
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *fiber.Ctx) error {
	var existingUser models.User

	name := c.FormValue("name")
	if name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Name is required",
		})
	}

	email := c.FormValue("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is required",
		})
	}
	if err := database.DB.Where("email = ?", email).First(&existingUser).Error;  err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Email is already taken",
		})
	}

	password := c.FormValue("password")
	if password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password is required",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to hash password",
		})
	}

	role := c.FormValue("role")
	if role == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Role is required",
		})
	}
	if role != "user" && role != "admin" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Role must be user or admin",
		})
	}

	newUser := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Failed to create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user": newUser,
	})
}

func Login(c *fiber.Ctx) error {
	var user models.User

	email := c.FormValue("email")
	if email == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Email is required",
		})
	}

	password := c.FormValue("password")
	if password == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Password is required",
		})
	}

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Wrong Credential",
		})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Could not create token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict", 
	})

	userLogin := models.UserLogin{
		Role: user.Role,
	}
	userLogin.ID = user.ID

	return c.Status(200).JSON(fiber.Map{
		"message": "login Successful",
		"user":    userLogin,
	})
}