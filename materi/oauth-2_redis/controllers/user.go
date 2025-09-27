package controllers

import (
	"log"
	"oauth-2_redis/database"
	"oauth-2_redis/models"
	"oauth-2_redis/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	user := models.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Ada data yang kosong, mohon diperiksa kembali",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat user",
		})
	}
	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Gagal membuat user",
		})
	}

	user.Password = ""

	return c.Status(201).JSON(fiber.Map{
		"message": "Akun berhasil dibuat, silahkan menunggu verifikasi admin dan coba login beberapa saat lagi",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Status(400).JSON(fiber.Map{"message": "Email dan password wajib diisi"})
	}

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Kredensial tidak valid"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"message": "Kredensial tidak valid"})
	}

	_, err := utils.GenerateAccessToken(c, user.ID, user.Name, user.Role)
	if err != nil {
		log.Println(err)
		return c.Status(500).JSON(fiber.Map{"message": "Gagal buat access token"})
	}

	signedRefresh, err := utils.GenerateRefreshToken(c, user.ID, user.Name, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Gagal buat refresh token"})
	}

	refreshToken := models.Token{
		RefreshToken: signedRefresh,
		ParentToken:  "",
		UserID:       user.ID,
		Exp:          time.Now().Add(7 * 24 * time.Hour).Unix(),
	}

	if err := utils.SaveRefreshToken(c, refreshToken); err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Redis error"})
	}

	return c.JSON(fiber.Map{
		"message": "Login berhasil",
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
		"message": "Logout sukses",
	})
}

func Me(c *fiber.Ctx) error {
	userID := c.Locals("id").(int)

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User tidak ditemukan",
		})
	}

	return c.JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user tidak ditemukan",
		})
	}

	if name := c.FormValue("name"); name != "" {
		user.Name = name
	}

	if password := c.FormValue("password"); password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Gagal membuat user",
			})
		}
		user.Password = string(hashedPassword)
	}

	if _, err := c.FormFile("profile_picture"); err == nil {
		file, err := utils.SaveFile(c, "profile_picture")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "harap masukkan gambar valid",
			})
		}
		user.ProfilePicture = file
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "gagal mengupdate user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user berhasil diupdate",
	})
}
