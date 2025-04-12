package main

import (
	"fmt"

	"Task_4/controllers"
	"Task_4/models"
)

func main() {
	user := controllers.LoadEnv()

	if !controllers.Autentikasi(user) {
		fmt.Println("Email atau password salah!")
		return
	}
	data := models.Data{}

	for {
		controllers.ShowMenu()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			controllers.SelectVehicle(&data)
		case 2:
			controllers.AddItem(&data)
		case 3:
			controllers.AddRecommendation(&data)
		case 4:
			controllers.AddFriend(&data)
		case 5:
			controllers.ShowData(data)
		case 6:
			fmt.Println("Terima kasih! Sampai jumpa di iftar.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Coba lagi.")
		}
	}
}
