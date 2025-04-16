package models

import (
	"os"
	"github.com/joho/godotenv"
)

type User struct {
	Name     string
	Email    string
	Password string
}

func LoadUser() (User, error) {
	err := godotenv.Load()
	if err != nil {
		return User{}, err
	}
	return User{
		Name:     os.Getenv("NAMA"),
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
	}, nil
}

func (u User) IsValid() bool {
	return u.Email != "" && u.Password != ""
}
