package config

import "github.com/joho/godotenv"

func ENVLoad() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}
