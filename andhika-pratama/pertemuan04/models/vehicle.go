package models

import "fmt"

var selectedVehicle string

func ChooseVehicle() {
	vehicles := []string{"Private Vehicle", "Budget Bus", "Carpool", "Travel"}
	fmt.Println("Choose 1 vehicle to go to iftar:")
	for i, v := range vehicles {
		fmt.Printf("%d. %s\n", i+1, v)
	}
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)
	if choice < 1 || choice > len(vehicles) {
		fmt.Println("Invalid choice.")
		return
	}
	selectedVehicle = vehicles[choice-1]
	fmt.Println("Selected vehicle:", selectedVehicle)
}
