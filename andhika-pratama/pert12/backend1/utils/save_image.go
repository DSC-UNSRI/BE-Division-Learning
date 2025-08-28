package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SaveFile(c *fiber.Ctx, fieldName string, fileType string) (string, error) {
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
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	var saveDir string
	var publicPath string

	switch fileType {
	case "profile":
		saveDir = "./assets/profile_pictures/"
		publicPath = "/assets/profile_pictures/"
	case "cover":
		saveDir = "./assets/covers/"
		publicPath = "/assets/covers/"
	default:
		return "", fmt.Errorf("unknown file type specified")
	}

	savePath := filepath.Join(saveDir, newFilename)
	if err := c.SaveFile(file, savePath); err != nil {
		return "", err
	}

	baseURL := os.Getenv("BASE_URL")
	fullURL := fmt.Sprintf("%s%s%s", baseURL, publicPath, newFilename)

	return fullURL, nil
}
