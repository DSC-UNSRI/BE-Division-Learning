package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func SaveFile(file *multipart.FileHeader, path string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return "", err
	}

	fileName := fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(file.Filename))
	filePath := filepath.Join(path, fileName)

	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return filePath, nil
}