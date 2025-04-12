package models

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
    godotenv.Load()
}

func Authenticate(email, password string) bool {
    return email == os.Getenv("EMAIL") && password == os.Getenv("PASSWORD")
}