package display

import (
	"fmt"
	"cobafix.go/dashboard/controllers"

)

func LihatSemuaData() {
	selected := controllers.GetKendaraan()

	fmt.Println("\n===== Data Kendaraan =====")
	if selected == "" {
		fmt.Println("Kendaraan kamu belum dipilih.")
	} else {
		fmt.Println("Kendaraan yang kamu pilih adalah:", selected)
	}
}
