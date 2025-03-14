package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Task_4/models"
)

func SelectVehicle(data *models.Data) {
	for {
		fmt.Println("Pilih kendaraan:")
		fmt.Println("1. Kendaraan Pribadi")
		fmt.Println("2. Bus Kaleng")
		fmt.Println("3. Nebeng")
		fmt.Println("4. Travel")
		fmt.Print("Masukkan pilihan (1/2/3/4): ")

		var choice int
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Input tidak valid! Masukkan angka 1-4.")
			continue
		}
		switch choice {
		case 1:
			data.VehicleChoice = "Kendaraan Pribadi"
		case 2:
			data.VehicleChoice = "Bus Kaleng"
		case 3:
			data.VehicleChoice = "Nebeng"
		case 4:
			data.VehicleChoice = "Travel"
		default:
			fmt.Println("Pilihan tidak valid! Silakan pilih kembali.")
			continue
		}
		break
	}
	fmt.Println("Kendaraan yang dipilih:", data.VehicleChoice)
}
func AddItem(data *models.Data) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Masukkan barang yang dibawa (atau ketik 'selesai' untuk berhenti): ")
		scanner.Scan()
		item := scanner.Text()
		if strings.ToLower(item) == "selesai" {
			break
		}
		data.Items = append(data.Items, item)
	}
}

func AddRecommendation(data *models.Data) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan kategori rekomendasi (misal: Film, Makanan): ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Print("Masukkan isi rekomendasi: ")
	scanner.Scan()
	content := scanner.Text()

	if data.Recommendations == nil {
		data.Recommendations = make(map[string]string)
	}
	data.Recommendations[category] = content
}

func AddFriend(data *models.Data) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama teman: ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Print("Masukkan divisi teman: ")
	scanner.Scan()
	division := scanner.Text()

	friend := models.Friend{Name: name, Division: division}
	data.Friends = append(data.Friends, friend)
}

func ShowData(data models.Data) {
	fmt.Println("\nDATA PENGGUNA")
	fmt.Println("Kendaraan yang dipilih:", data.VehicleChoice)
	fmt.Println("\nBarang yang dibawa:")
	for _, item := range data.Items {
		fmt.Println("-", item)
	}
	fmt.Println("\nRekomendasi Iftar:")
	for category, content := range data.Recommendations {
		fmt.Printf("- %s: %s\n", category, content)
	}
	fmt.Println("\nTeman yang ikut:")
	for _, friend := range data.Friends {
		fmt.Printf("- %s (Divisi: %s)\n", friend.Name, friend.Division)
	}
}
