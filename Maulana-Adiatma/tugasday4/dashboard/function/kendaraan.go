package dashboard

import "fmt"

func PilihKendaraan() {
	fmt.Println("==== Pilih Kendaraan ====")
	fmt.Println("= 1. Kendaraan Pribadi  =")
	fmt.Println("= 2. Bus Kaleng         =")
	fmt.Println("= 3. Nebeng             =")
	fmt.Println("= 4. Travel             =")
	fmt.Println("=========================")

	var opsi1, opsi2 int
	fmt.Print("Masukkan pilihan kendaraan pertama: ")
	fmt.Scan(&opsi1)
	fmt.Print("Masukkan pilihan kendaraan kedua: ")
	fmt.Scan(&opsi2)
	cetakPilihan := func(pilihan int) {
		switch pilihan {
		case 1:
			fmt.Println("save!! kamu menggunakan kendaraan pribadi")
		case 2:
			fmt.Println("save!! kamu menggunakan Bus kaleng")
		case 3:
			fmt.Println("save!! kamu Nebeng bersama kawanmu")
		case 4:
			fmt.Println("save!! kamu menggunakan jasa Travel")
		default:
			fmt.Println("Pilihan mu tidak sesuai silahkan input 1 - 4")
		}
	}

	fmt.Println("=== Pilihan kendaraanmu ===")
	cetakPilihan(opsi1)
	cetakPilihan(opsi2)
}
