package database

import (
	"fmt"
	"log"
	"tugas7/models"
)

func Migrate() {
	_, err := DB.Exec(models.MahasiswaQuery)
	if err != nil {
		log.Fatalf("Failed to migrate mahasiswa table: %v", err)
	}
	fmt.Println("Migrated Mahasiswa table successfully")

	_, err = DB.Exec(models.MinatQuery)
	if err != nil {
		log.Fatalf("Failed to migrate minat table: %v", err)
	}
	fmt.Println("Migrated Minat table successfully")

	_, err = DB.Exec(models.MahasiswaMinatQuery)
	if err != nil {
		log.Fatalf("Failed to migrate mahasiswa_minat table: %v", err)
	}
	fmt.Println("Migrated MahasiswaMinat table successfully")

	fmt.Println("Database migration completed successfully!")
}