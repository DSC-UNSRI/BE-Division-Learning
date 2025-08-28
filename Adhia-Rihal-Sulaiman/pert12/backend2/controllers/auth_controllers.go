package controllers

import (
	"backend2/database"
	"backend2/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *fiber.Ctx) error {
	var existingUser models.User

	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")
	profilePicture := c.FormValue("profile_picture")

	if name == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Name is required"})
	}
	if email == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Email is required"})
	}
	if password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Password is required"})
	}
	if role == "" {
		role = "user"
	}
	if role != "user" && role != "admin" {
		return c.Status(400).JSON(fiber.Map{"message": "Role must be user or admin"})
	}
	if profilePicture == "" {
		profilePicture = "default.png"
	}

	if err := database.DB.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{"message": "Email is already taken"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to hash password"})
	}

	newUser := models.User{
		Name:           name,
		Email:          email,
		Password:       string(hashedPassword),
		Role:           role,
		ProfilePicture: profilePicture,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create user"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User created successfully",
		"user": fiber.Map{
			"id":              newUser.ID,
			"name":            newUser.Name,
			"email":           newUser.Email,
			"role":            newUser.Role,
			"profile_picture": newUser.ProfilePicture,
		},
	})
}

func Login(c *fiber.Ctx) error {
	var user models.User

	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Email and password are required"})
	}

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(3 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to create token"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	userLogin := models.UserLogin{Role: user.Role}
	userLogin.ID = user.ID

	return c.Status(200).JSON(fiber.Map{
		"message": "Login successful",
		"user":    userLogin,
		"token":   tokenString,
	})
}

func Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.Status(200).JSON(fiber.Map{"message": "Logout successful"})
}
