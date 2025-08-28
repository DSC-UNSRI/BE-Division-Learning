package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string, isProgramStudi bool, isCover bool) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		fmt.Println("Error saat mengambil file dari form:", err) // Tambahan
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	if !isProgramStudi {
		allowedExt[".pdf"] = true
	}

	if !allowedExt[ext] {
		err := fmt.Errorf("invalid file type")
		fmt.Println("Error: Tipe file tidak valid:", err) // Tambahan
		return "", err
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := "./assets/files/"
	publicPath := "/assets/files/"
	if isProgramStudi {
		saveDir = "./assets/logo/"
		publicPath = "/assets/logo/"
	} else if isCover {
		saveDir = "./assets/cover/"
		publicPath = "/assets/cover/"

	}

	savePath := filepath.Join(saveDir, newFilename)
	if err := c.SaveFile(file, savePath); err != nil {
		fmt.Println("Error saat menyimpan file:", err) // Tambahan
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	fullURL := fmt.Sprintf("%s%s%s", baseURL, publicPath, newFilename)

	return fullURL, nil
}
