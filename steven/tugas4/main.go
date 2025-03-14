package main

import (
	"fmt"
	"os"
	"tugas4/controller/authController"
)

func main() {
	fmt.Println("=== WELCOME TO IFTAR GDGoC ===")
	fmt.Println("1. Login")
	fmt.Println("2. EXIT")

	var pilihan int;
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		authController.Auth()
	case 2:
		fmt.Println("Keluar dari program...")
		os.Exit(0)
	}
}
