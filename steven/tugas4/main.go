package main

import (
	"fmt"
	"os"
	"tugas4/controller"
)

func main() {
	fmt.Println("=== WELCOME TO IFTAR GDGoC ===")
	fmt.Println("1. Login")
	fmt.Println("2. EXIT")

	var pilihan int;
	fmt.Print("MASUKKAN PILIHAN :")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		controller.Auth()
	case 2:
		fmt.Println("Keluar dari program...")
		os.Exit(0)
	default:
        fmt.Println("Pilihan tidak tersedia")
		main()
    }
}
