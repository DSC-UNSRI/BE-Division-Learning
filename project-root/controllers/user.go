package controllers

import (
	"bufio"
	"fmt"
	"os"
	"project-root/models"
	"strings"
)

var UserData = models.User{}

func InitUserData(nama, email string) {
	UserData.Nama = nama
	UserData.Email = email
}

func PilihKendaraan() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPilih Kendaraan:")
		fmt.Println("1. Kendaraan Pribadi")
		fmt.Println("2. Bus Kaleng")
		fmt.Println("3. Nebeng")
		fmt.Println("4. Travel")
		fmt.Println("5. Kembali")
		fmt.Print("Pilihan (1-5): ")

		pilihan, _ := reader.ReadString('\n')
		pilihan = strings.TrimSpace(pilihan)

		switch pilihan {
		case "1":
			UserData.Kendaraan = "Kendaraan Pribadi"
			return
		case "2":
			UserData.Kendaraan = "Bus Kaleng"
			return
		case "3":
			UserData.Kendaraan = "Nebeng"
			return
		case "4":
			UserData.Kendaraan = "Travel"
			return
		case "5":
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}

func TambahBarang() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nMasukkan barang yang akan dibawa (atau 'selesai' untuk keluar): ")
		barang, _ := reader.ReadString('\n')
		barang = strings.TrimSpace(barang)

		if strings.ToLower(barang) == "selesai" {
			break
		}

		if barang != "" {
			UserData.Barang = append(UserData.Barang, barang)
			fmt.Printf("Barang '%s' ditambahkan\n", barang)
		} else {
			fmt.Println("Nama barang tidak boleh kosong")
		}
	}
}

func TambahRekomendasi() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nKategori rekomendasi (atau 'selesai' untuk keluar): ")
		kategori, _ := reader.ReadString('\n')
		kategori = strings.TrimSpace(kategori)

		if strings.ToLower(kategori) == "selesai" {
			break
		}

		if kategori == "" {
			fmt.Println("Kategori tidak boleh kosong")
			continue
		}

		fmt.Print("Isinya: ")
		isi, _ := reader.ReadString('\n')
		isi = strings.TrimSpace(isi)

		if isi == "" {
			fmt.Println("Isi rekomendasi tidak boleh kosong")
			continue
		}

		UserData.Rekomendasi = append(UserData.Rekomendasi, models.Rekomendasi{
			Kategori: kategori,
			Isi:      isi,
		})
		fmt.Println("Rekomendasi ditambahkan")
	}
}

func TambahTeman() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nNama Teman (atau 'selesai' untuk keluar): ")
		nama, _ := reader.ReadString('\n')
		nama = strings.TrimSpace(nama)

		if strings.ToLower(nama) == "selesai" {
			break
		}

		if nama == "" {
			fmt.Println("Nama teman tidak boleh kosong")
			continue
		}

		fmt.Print("Divisi: ")
		divisi, _ := reader.ReadString('\n')
		divisi = strings.TrimSpace(divisi)

		if divisi == "" {
			fmt.Println("Divisi tidak boleh kosong")
			continue
		}

		UserData.Teman = append(UserData.Teman, models.Teman{
			Nama:   nama,
			Divisi: divisi,
		})
		fmt.Printf("Teman '%s' ditambahkan\n", nama)
	}
}

func LihatData() {
	fmt.Println("\n=== DATA USER ===")
	fmt.Println("Nama:", UserData.Nama)
	fmt.Println("Email:", UserData.Email)
	fmt.Println("Kendaraan:", UserData.Kendaraan)

	fmt.Println("\nBarang yang dibawa:")
	if len(UserData.Barang) == 0 {
		fmt.Println("- Belum ada barang")
	} else {
		for i, b := range UserData.Barang {
			fmt.Printf("%d. %s\n", i+1, b)
		}
	}

	fmt.Println("\nRekomendasi:")
	if len(UserData.Rekomendasi) == 0 {
		fmt.Println("- Belum ada rekomendasi")
	} else {
		for i, r := range UserData.Rekomendasi {
			fmt.Printf("%d. %s: %s\n", i+1, r.Kategori, r.Isi)
		}
	}

	fmt.Println("\nTeman yang ikut:")
	if len(UserData.Teman) == 0 {
		fmt.Println("- Belum ada teman")
	} else {
		for i, t := range UserData.Teman {
			fmt.Printf("%d. %s (%s)\n", i+1, t.Nama, t.Divisi)
		}
	}
}
