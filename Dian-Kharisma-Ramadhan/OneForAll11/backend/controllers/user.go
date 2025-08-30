package controllers

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"backend/config"
	"backend/models"
	"backend/utils"
)

func GetMe(c *fiber.Ctx) error {
	// Dapatkan nilai userID dari JWT claims
	userIDFloat64 := c.Locals("userID").(float64)

	// Konversi dari float64 ke uint
	userID := uint(userIDFloat64)

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	return c.JSON(user)
}

func UpdateProfile(c *fiber.Ctx) error {
	log.Println("Memulai UpdateProfile...")
	
	// Dapatkan ID pengguna dari token
	userIDFloat64 := c.Locals("userID").(float64)
	authUserID := uint(userIDFloat64)

	// Dapatkan ID pengguna dari URL
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Error: Invalid user ID di URL.")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid user ID"})
	}

	// Cek otorisasi: pastikan pengguna hanya bisa mengupdate profilnya sendiri
	if authUserID != uint(id) {
		log.Println("Error: Pengguna tidak diotorisasi untuk update profil ini.")
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You are not authorized to update this profile"})
	}

	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		log.Printf("Error: Pengguna dengan ID %d tidak ditemukan.\n", id)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	// Perbarui nama jika ada
	if name := c.FormValue("name"); name != "" {
		user.Name = name
		log.Printf("Nama baru diterima: %s\n", name)
	}

	// Perbarui password jika ada
	if password := c.FormValue("password"); password != "" {
		log.Println("Password baru diterima.")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error: Gagal hash password baru.")
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to hash password"})
		}
		user.Password = string(hashedPassword)
	}

	// Unggah file foto profil
	log.Println("Mencoba mengunggah foto profil...")
	file, err := c.FormFile("profile_picture")
	if err == nil {
		log.Println("File foto profil diterima.")
		filePath, err := utils.SaveFile(file, "./assets/images/profiles")
		if err != nil {
			log.Println("Error: Gagal menyimpan file.", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to upload file"})
		}
		user.ProfilePicture = filePath
		log.Printf("File foto profil berhasil disimpan di: %s\n", filePath)
	} else if err.Error() != "no file header" {
		log.Println("Error: Gagal mendapatkan file dari form.", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to get file from form"})
	}

	// Simpan perubahan ke database
	log.Println("Mencoba menyimpan perubahan ke database...")
	if err := config.DB.Save(&user).Error; err != nil {
		log.Println("Error: Gagal menyimpan ke database.", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to save profile"})
	}
	
	log.Println("Profil berhasil diperbarui.")
	return c.JSON(fiber.Map{"message": "Profile updated successfully", "profile_picture_url": user.ProfilePicture})
}
