package main

import (
	"backend-iftar-gdgoc/controllers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: Gagal memuat file .env. Pastikan file tersebut ada di root direktori.")
		return err
	}
	return nil
}

func AmbilVariabel(key string) string {
	return os.Getenv(key)
}
func main() {
	
	LoadEnv()

	if !controllers.AuthenticateUser() {
		fmt.Println("Autentikasi gagal. Cek kembali email dan password.")
		return
	}

	controllers.Dashboard()
}
