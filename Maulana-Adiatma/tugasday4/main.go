package main

import(
	"fmt"
	"tugasday4/autentikasi"
)

func main(){
	if autentikasi.Login() {
		dashboard.MenuDashboard()
	} else {
		fmt.Println("Gagal login. Program berhenti.")
	}
}
