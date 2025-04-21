package database

import (
	"log"

	"Tugas-Pertemuan-5/models"
)

func Migrate() {
	var err error

	_, err = DB.Exec(models.DirectorsQuery)
	if err != nil {
		log.Fatalf("Gagal migrasi tabel directors: %v", err)
	}

	_, err = DB.Exec(models.FilmsQuery)
	if err != nil {
		log.Fatalf("Gagal migrasi tabel films: %v", err)
	}

	_, err = DB.Exec(models.UsersQuery)
	if err != nil {
		log.Fatalf("Gagal migrasi tabel users: %v", err)
	}

	log.Println("Migrasi Database Selesai.")
}