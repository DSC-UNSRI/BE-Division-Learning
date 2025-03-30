package controllers

var kendaraanTerpilih string

func Kendaraan(pilihan int) {
	switch pilihan {
	case 1:
		kendaraanTerpilih = "Kendaraan Pribadi"
	case 2:
		kendaraanTerpilih = "Bus Kaleng"
	case 3:
		kendaraanTerpilih = "Nebeng"
	case 4:
		kendaraanTerpilih = "Travel"
	default:
		kendaraanTerpilih = ""
	}
}


func GetKendaraan() string {
	return kendaraanTerpilih
}
