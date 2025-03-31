package display

import (
	"fmt"
	"cobafix.go/dashboard/controllers"

)

func LihatSemuaData() {
	fmt.Println("\n===== Data Kendaraan =====")
	controllers.PrintKendaraan()

	fmt.Println("\n====== Data Barang =======")
	controllers.ViewBarang()
}
