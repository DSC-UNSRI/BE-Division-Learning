package models

import "fmt"

func ChooseVehicle() {
	vehicles := []string{"Private Vehicle", "Budget Bus", "Hitch a Ride", "Travel Car"}
	fmt.Println("Pilih 1 kendaraan untuk menuju iftar:")
	for i, v := range vehicles {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	var choice int
	fmt.Print("Pilihan: ")
	fmt.Scan(&choice)
	if choice < 1 || choice > len(vehicles) {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	selectedVehicle := vehicles[choice-1]
	fmt.Println("Kendaraan terpilih:", selectedVehicle)
}
