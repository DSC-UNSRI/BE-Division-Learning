package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string, required bool) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		if required {
			return "", err
		}
		// kalau nggak required → return kosong tanpa error
		return "", nil
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}
	if !required { // kalau bukan cover (misal lampiran dokumen) boleh pdf
		allowedExt[".pdf"] = true
	}

	if !allowedExt[ext] {
		return "", fmt.Errorf("invalid file type")
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := "./assets/profile/"
	publicPath := "/assets/profile/"
	if required {
		saveDir = "./assets/cover/"
		publicPath = "/assets/cover/"
	}

	savePath := filepath.Join(saveDir, newFilename)
	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	fullURL := fmt.Sprintf("%s%s%s", baseURL, publicPath, newFilename)

	return fullURL, nil
}
