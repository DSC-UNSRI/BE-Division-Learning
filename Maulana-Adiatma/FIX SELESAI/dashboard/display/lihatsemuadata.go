package display

import (
	"fmt"
	"cobafix.go/dashboard/controllers"

)

func LihatSemuaData() {
	fmt.Println("\n========= Data Kendaraan ========")
	controllers.ViewKendaraan()

	fmt.Println("\n========= Data Barang ===========")
	controllers.ViewBarang()

	fmt.Println("\n======= Data Rekomendasi ========")
	controllers.ViewRekomendasi(daftarRekomendasi)

	fmt.Println("\n========== Data Teman ===========")
	controllers.ViewTeman(daftarTeman)
}
