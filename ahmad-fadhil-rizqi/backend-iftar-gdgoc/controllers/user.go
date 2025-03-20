package controllers

import "backend-iftar-gdgoc/config"

type User struct {
	Nama     string
	Email    string
	Password string
}

func AmbilUser() User {
	return User{
		Nama:     config.AmbilVariabel("NAMA"),
		Email:    config.AmbilVariabel("EMAIL"),
		Password: config.AmbilVariabel("PASSWORD"),
	}
}
