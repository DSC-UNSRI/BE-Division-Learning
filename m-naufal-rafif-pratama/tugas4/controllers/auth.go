package controllers

import (
	"log"
	"tugas4/models"
)

func Authenticate() models.User {
	user, err := models.LoadUser()
	if err != nil || !user.IsValid() {
		log.Fatal("Autentikasi gagal")
	}
	return user
}
