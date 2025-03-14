package controllers

import (
	"log"
	"tugas4/models"
)

func Authenticate() models.User {
	user, err := models.LoadUser()
	if err != nil || !user.IsValid() && user.Email == "naufal@gmail.com" && user.Password == "naufal123" {
		log.Fatal("Autentikasi gagal")
	}
	return user
}
