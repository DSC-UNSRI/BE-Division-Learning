package controllers

import (
	"bufio"
	"fmt"
	"os"
)

var daftarBarang []string  

func InputBarang() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\nMasukkan barang yang akan dibawa (ketik 'selesai' untuk mengakhiri):")

	for {
		fmt.Print("- ")
		scanner.Scan()
		item := scanner.Text()

		if item == "selesai" || item == "SELESAI" {
			break
		}

		daftarBarang = append(daftarBarang, item)
	}

	fmt.Println("\nBarang telah ditambahkan!")
}

func HapusBarang() {
	fmt.Print("\nMasukkan nama barang yang ingin dihapus: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	item := scanner.Text()

	index := -1
	for i, v := range daftarBarang {
		if v == item {
			index = i
			break
		}
	}

	if index != -1 {
		daftarBarang = append(daftarBarang[:index], daftarBarang[index+1:]...)
		fmt.Println("\nBarang berhasil dihapus!")
	} else {
		fmt.Println("\nBarang tidak ditemukan.")
	}
}

func ViewBarang() {
	if len(daftarBarang) == 0 {
		fmt.Println("Kamu belum input barang.")
		return
	}

	fmt.Println("\nDaftar barang untuk iftar:")
	for i, item := range daftarBarang {
		fmt.Printf("%d. %s\n", i+1, item)
	}
}