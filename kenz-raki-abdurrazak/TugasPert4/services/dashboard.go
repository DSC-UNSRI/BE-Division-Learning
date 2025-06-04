package services

import (
	"fmt"
	"github.com/artichys/BE-Division-Learning/models"
	"github.com/artichys/BE-Division-Learning/utils"
)

var (
	selectedVehicle  models.Vehicle
	items           []models.Item
	recommendations []models.Recommendation
	friends         []models.Friend
)

func ChooseVehicle() {
	options := []string{"Kendaraan Pribadi", "Bus Kaleng", "Nebeng", "Travel"}

	fmt.Println("Pilih kendaraan (1-4):")
	for i, option := range options {
		fmt.Printf("%d. %s\n", i+1, option)
	}

	choice := utils.ReadInt()

	if choice < 1 || choice > 4 {
		fmt.Println("Pilihan tidak valid! Silakan coba lagi.")
		return
	}

	selectedVehicle = models.Vehicle{FirstChoice: options[choice-1]}

	fmt.Println("Kendaraan berhasil dipilih:", selectedVehicle.FirstChoice)
}

func AddItem() {
	fmt.Print("Masukkan nama barang: ")
	itemName := utils.ReadInput()
	items = append(items, models.Item{Name: itemName})
	fmt.Println("Barang berhasil ditambahkan:", itemName)
}

func AddRecommendation() {
	fmt.Print("Masukkan kategori rekomendasi: ")
	category := utils.ReadInput()
	fmt.Print("Masukkan rekomendasi: ")
	content := utils.ReadInput()
	recommendations = append(recommendations, models.Recommendation{Category: category, Content: content})
	fmt.Println("Rekomendasi berhasil ditambahkan:", category, "-", content)
}

func AddFriend() {
	fmt.Print("Masukkan nama teman: ")
	name := utils.ReadInput()
	fmt.Print("Masukkan divisi teman: ")
	division := utils.ReadInput()
	friends = append(friends, models.Friend{Name: name, Division: division})
	fmt.Println("Teman berhasil ditambahkan:", name, "(", division, ")")
}

func ShowData() {
	fmt.Println("\n=== Data Dashboard ===")
	fmt.Println("Kendaraan:", selectedVehicle.FirstChoice)

	fmt.Println("\nBarang yang dibawa:")
	if len(items) == 0 {
		fmt.Println("- Tidak ada barang")
	} else {
		for _, item := range items {
			fmt.Println("-", item.Name)
		}
	}

	fmt.Println("\nRekomendasi:")
	if len(recommendations) == 0 {
		fmt.Println("- Tidak ada rekomendasi")
	} else {
		for _, rec := range recommendations {
			fmt.Println(rec.Category+":", rec.Content)
		}
	}

	fmt.Println("\nTeman yang ikut:")
	if len(friends) == 0 {
		fmt.Println("- Tidak ada teman")
	} else {
		for _, friend := range friends {
			fmt.Println(friend.Name, "(", friend.Division, ")")
		}
	}
}
