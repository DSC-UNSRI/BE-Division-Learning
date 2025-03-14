package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"tugas-4/models"
)

func StartDashboard(user models.User) {
	fmt.Println("Login berhasil! Selamat datang,", user.Name)

	var dashboard models.Dashboard
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Pilih Kendaraan")
		fmt.Println("2. Tambahkan barang")
		fmt.Println("3. Tambahkan Rekomendasi")
		fmt.Println("4. Tambah teman yang ikut iftar")
		fmt.Println("5. Lihat data")
		fmt.Println("6. Exit")
		fmt.Print("Pilih Opsi: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Pilih kendaraan (Kendaraan Pribadi, Bus Kaleng, Nebeng): ")
			scanner.Scan()
			dashboard.VehicleChoice = scanner.Text()

		case 2:
			fmt.Println("Masukkan barang untuk dibawa: ")
			scanner.Scan()
			dashboard.Items = append(dashboard.Items, scanner.Text())

		case 3:
			fmt.Println("Masukkan kategori: ")
			scanner.Scan()
			category := scanner.Text()

			if dashboard.Recommendations == nil {
				dashboard.Recommendations = make(map[string]string)
			}

			fmt.Println("Masukkan rekomendasi: ")
			scanner.Scan()
			dashboard.Recommendations[category] = scanner.Text()

		case 4:

			fmt.Println("Masukkan nama teman: ")
			scanner.Scan()
			friendName := scanner.Text()

			fmt.Println("Masukkan divisi teman: ")
			scanner.Scan()
			friendDivision := scanner.Text()

			newFriend := models.Friend{
				Name:     friendName,
				Division: friendDivision,
			}
			dashboard.Friends = append(dashboard.Friends, newFriend)

		case 5:
			fmt.Println("\n--- Dashboard Data ---")
			fmt.Println("Kendaraan:", dashboard.VehicleChoice)
			fmt.Println("Item:", strings.Join(dashboard.Items, ", "))
			fmt.Println("Rekomendasi:")
			for cat, rec := range dashboard.Recommendations {
				fmt.Println("-", cat, ":", rec)
			}
			fmt.Println("Teman:")
			for _, f := range dashboard.Friends {
				fmt.Printf("- %s (Divisi: %s)\n", f.Name, f.Division)
			}

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, coba lagi.")
		}
	}
}
