package controllers

import (
	"fmt"
	"project-root/models"
)

var users []models.User

func AddUser(name, email string) {
	user := models.User{Name: name, Email: email}
	users = append(users, user)
	fmt.Println("User berhasil ditambahkan:", name)
}

func ListUsers() {
	fmt.Println("Daftar Users:")
	for _, user := range users {
		fmt.Printf("Nama: %s, Email: %s\n", user.Name, user.Email)
	}
}
