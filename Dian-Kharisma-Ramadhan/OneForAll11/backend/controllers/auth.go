package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"backend/config"
	"backend/models"
)

func Login(c *fiber.Ctx) error {
	// Ambil data langsung dari form value, karena frontend mengirim multipart/form-data
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Lakukan validasi dasar
	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Email and password are required"})
	}

	var user models.User
	// Cari user berdasarkan email
	if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		// Mengembalikan pesan generik untuk alasan keamanan
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	// Bandingkan password yang di-hash dari database dengan password plain-text dari form
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Mengembalikan pesan generik yang sama
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid email or password"})
	}

	// Buat token JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret")) // Ganti dengan secret key yang kuat
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not generate token"})
	}

	// Mengembalikan token dan data user (tanpa password)
	return c.JSON(fiber.Map{
		"token": tokenString,
		"user":  user,
	})
}

func Signup(c *fiber.Ctx) error {
	// Ambil data langsung dari form value
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	role := c.FormValue("role")

	// Lakukan validasi dasar
	if name == "" || email == "" || password == "" || role == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "All fields are required"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not hash password"})
	}

	// Buat objek user
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     role,
	}

	// Simpan pengguna ke DB
	if err := config.DB.Create(&user).Error; err != nil {
		// Error mungkin karena email sudah ada (unique constraint)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Could not create user"})
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

func Logout(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "Logged out successfully"})
}