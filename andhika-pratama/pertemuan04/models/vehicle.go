package models

import "fmt"

func ChooseVehicle() {
	vehicles := []string{"Kendaraan Pribadi", "Bus Kaleng", "Nebeng", "Travel"}
	var choices [2]int

	fmt.Println("Pilih 2 kendaraan untuk menuju iftar:")
	for i, v := range vehicles {
		fmt.Printf("%d. %s\n", i+1, v)
	}

	for i := 0; i < 2; i++ {
		fmt.Printf("Pilihan %d: ", i+1)
		fmt.Scan(&choices[i])
		if choices[i] < 1 || choices[i] > len(vehicles) {
			fmt.Println("Pilihan tidak valid, silakan ulangi.")
			i--
			continue
		}
		if i == 1 && choices[0] == choices[1] {
			fmt.Println("Kedua pilihan tidak boleh sama, silakan pilih ulang pilihan kedua.")
			i--
		}
	}

	selectedVehicles := []string{vehicles[choices[0]-1], vehicles[choices[1]-1]}
	fmt.Println("Kendaraan terpilih:", selectedVehicles[0], "dan", selectedVehicles[1])
}
