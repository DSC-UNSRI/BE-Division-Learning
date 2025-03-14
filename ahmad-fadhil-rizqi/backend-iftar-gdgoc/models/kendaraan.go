package models

import (
	"backend-iftar-gdgoc/data"
	"fmt"
)

type Vehicle struct {
	Name string
}

func SelectVehicle(vehicle *[]Vehicle) {
	fmt.Println("\n===== Choose Your Vehicle =====")
	fmt.Println("1. Bus Kaleng")
	fmt.Println("2. Private Car")
	fmt.Println("3. Travel")
	fmt.Println("4. Hitchhiking (Nebeng)")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	var selectedVehicle string
	switch choice {
	case 1:
		selectedVehicle = "Bus Kaleng"
	case 2:
		selectedVehicle = "Private Car"
	case 3:
		selectedVehicle = "Travel"
	case 4:
		selectedVehicle = "Hitchhiking"
	default:
		fmt.Println("Invalid choice.")
		return
	}

	// Hanya boleh memilih satu kendaraan (overwrite data lama)
	*vehicle = []Vehicle{{Name: selectedVehicle}}
	data.CatatLog("Selected vehicle: " + selectedVehicle)
	fmt.Println("Vehicle successfully selected:", selectedVehicle)
}
