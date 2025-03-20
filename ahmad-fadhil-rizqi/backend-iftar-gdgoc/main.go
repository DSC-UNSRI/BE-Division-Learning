package main

import (
	"backend-iftar-gdgoc/config"
	"backend-iftar-gdgoc/controllers"
	"fmt"
	"log"
)

func main() {
	err := config.LoadEnv()
	if err != nil {
		log.Fatal("Gagal memuat konfigurasi")
	}

	if !controllers.AuthenticateUser() {
		fmt.Println("Autentikasi gagal. Cek kembali email dan password.")
		return
	}

	controllers.Dashboard()
}
