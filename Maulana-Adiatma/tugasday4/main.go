package main

import(
	"fmt"
	"tugasday4/autentikasi"
	"tugasday4/dashboard"
)

func main(){
	if autentikasi.Login() {
		dashboard.MenuDashboard()
	} else {
		fmt.Println("Gagal login. Program berhenti.")
	}
}
