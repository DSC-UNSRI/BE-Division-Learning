package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func ENVinit() {
	if err := godotenv.Load(); err != nil {
		panic("failed load .env")
	}
	fmt.Println("success load .env")
}