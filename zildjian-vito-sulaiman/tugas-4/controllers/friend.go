package controllers

import (
	"bufio"
	"fmt"

	"tugas-4/models"
)

func ManageFriends(dashboard *models.Dashboard, scanner *bufio.Scanner) {
	fmt.Println("\n--- Kelola Teman ---")
	fmt.Println("1. Tambah Teman")
	fmt.Println("2. Perbarui Teman")
	fmt.Print("Pilih menu: ")

	var choose int
	fmt.Scanln(&choose)

	switch choose {
	case 1:
		addFriend(dashboard, scanner)
	case 2:
		updateFriend(dashboard, scanner)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func addFriend(dashboard *models.Dashboard, scanner *bufio.Scanner) {
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
	fmt.Println("Teman berhasil ditambahkan!")
}

func updateFriend(dashboard *models.Dashboard, scanner *bufio.Scanner) {
	if len(dashboard.Friends) == 0 {
		fmt.Println("❌ Tidak ada teman untuk diperbarui.")
		return
	}

	fmt.Println("\nDaftar Teman:")
	for i, f := range dashboard.Friends {
		fmt.Printf("%d. %s (Divisi: %s)\n", i+1, f.Name, f.Division)
	}

	fmt.Print("\nMasukkan nomor teman yang ingin diperbarui: ")
	var index int
	fmt.Scanln(&index)

	if index < 1 || index > len(dashboard.Friends) {
		fmt.Println("❌ Pilihan tidak valid.")
		return
	}

	fmt.Println("Masukkan nama baru: ")
	scanner.Scan()
	dashboard.Friends[index-1].Name = scanner.Text()

	fmt.Println("Masukkan divisi baru: ")
	scanner.Scan()
	dashboard.Friends[index-1].Division = scanner.Text()

	fmt.Println("Data teman berhasil diperbarui!")
}
