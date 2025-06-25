package database

import (
	"log"

	"Tugas-Pertemuan-7/models"
)

func Migrate() {
	var err error

	_, err = DB.Exec(models.DirectorsQuery)
	if err != nil {
		log.Fatalf("Failed to migrate directors table: %v", err)
	}

	_, err = DB.Exec(models.FilmsQuery)
	if err != nil {
		log.Fatalf("Failed to migrate films table: %v", err)
	}

	log.Println("Database migration complete.")
}