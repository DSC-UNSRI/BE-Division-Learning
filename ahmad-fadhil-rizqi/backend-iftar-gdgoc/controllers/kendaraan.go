package controllers

import (
	"backend-iftar-gdgoc/models"
	"fmt"
)

func SelectVehicle(vehicle *[]models.Vehicle) {
	fmt.Println("\n===== Pilih Kendaraan =====")
	fmt.Println("1. Bus Kaleng")
	fmt.Println("2. Kendaraan Pribadi")
	fmt.Println("3. Travel")
	fmt.Println("4. Nebeng")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	var selectedVehicle string
	switch choice {
	case 1:
		selectedVehicle = "Bus Kaleng"
	case 2:
		selectedVehicle = "Kendaraan Pribadi"
	case 3:
		selectedVehicle = "Travel"
	case 4:
		selectedVehicle = "Nebeng"
	default:
		fmt.Println("Invalid .")
		return
	}

	*vehicle = []models.Vehicle{{Name: selectedVehicle}}
	fmt.Println("Vehicle successfully selected:", selectedVehicle)
}
