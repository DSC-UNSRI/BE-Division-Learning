package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	return godotenv.Load()
}

func AmbilVariabel(kunci string) string {
	return os.Getenv(kunci)
}
