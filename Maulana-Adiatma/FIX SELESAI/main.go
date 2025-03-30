package main

import (
	"fmt"

	"cobafix.go/autentikasi"
)

func main() {
	fmt.Println("Silakan login terlebih dahulu")
	if autentikasi.Login() {
		fmt.Println("Akses diberikan")
	} else {
		fmt.Println("Akses ditolak")
	}
}
