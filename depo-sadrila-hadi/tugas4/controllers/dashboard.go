package controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"tugas4/models"
)

func ShowMenu(data *models.Data) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nIftar GDGoC Dashboard")
		fmt.Println("1. Choose Vehicle")
		fmt.Println("2. Enter Items to Bring")
		fmt.Println("3. Enter Recommendation")
		fmt.Println("4. Enter Friend Data")
		fmt.Println("5. View All Data")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			fmt.Println("Choose a Vehicle:")
			fmt.Println("1. Private Vehicle")
			fmt.Println("2. Tin Bus")
			fmt.Println("3. Hitchhiking")
			fmt.Println("4. Travel")
			fmt.Print("Enter your choice (1-4): ")
			scanner.Scan()
			choice := strings.TrimSpace(scanner.Text())
			var vehicle string
			switch choice {
			case "1":
				vehicle = "Private Vehicle"
			case "2":
				vehicle = "Tin Bus"
			case "3":
				vehicle = "Hitchhiking"
			case "4":
				vehicle = "Travel"
			default:
				fmt.Println("Invalid option")
				continue
			}
			data.Kendaraan = []string{vehicle}
		case "2":
			fmt.Println("Enter items (separated by commas if more than one): ")
			scanner.Scan()
			items := strings.Split(scanner.Text(), ",")
			data.Barang = items
		case "3":
			fmt.Println("Recommendation category: ")
			scanner.Scan()
			category := scanner.Text()
			fmt.Println("Recommendation content (separated by commas if more than one): ")
			scanner.Scan()
			content := strings.Split(scanner.Text(), ",")
			data.Rekomendasi[category] = content
		case "4":
			fmt.Println("Friend's name: ")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println("Friend's division: ")
			scanner.Scan()
			division := scanner.Text()
			data.Teman = append(data.Teman, models.Teman{Nama: name, Divisi: division})
		case "5":
			fmt.Printf("Vehicle: %v\n", data.Kendaraan)
			fmt.Printf("Items: %v\n", data.Barang)
			fmt.Printf("Recommendations: %v\n", data.Rekomendasi)
			fmt.Printf("Friends: %v\n", data.Teman)
		case "6":
			fmt.Println("Exiting Dashboard")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}