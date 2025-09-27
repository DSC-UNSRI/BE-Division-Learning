package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string) (string, error) {
	file, err := c.FormFile(fieldName)
	if err != nil {
		return "", err
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExt := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	if !allowedExt[ext] {
		return "", fmt.Errorf("invalid file type")
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	saveDir := "./assets/"
	publicPath := "/assets/"

	savePath := filepath.Join(saveDir, newFilename)
	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	fullURL := fmt.Sprintf("%s%s%s", baseURL, publicPath, newFilename)

	return fullURL, nil
}