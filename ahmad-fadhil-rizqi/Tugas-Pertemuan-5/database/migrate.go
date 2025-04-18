package database

import (
	"fmt"
	"log"

	"Tugas-Pertemuan-5/models"
)

func Migrate() {
	_, err := DB.Exec(models.FilmsQuery)
	if err != nil {
		log.Fatalf("Gagal migrasi tabel films: %v", err)
	}

	_, err = DB.Exec(models.UsersQuery)
	if err != nil {
		log.Fatalf("Gagal migrasi tabel users: %v", err)
	}

	fmt.Println("Migrasi Sukses")
}