package database

import (
	"fmt"
	"log"
	"tugas7/models"
)

func Migrate() {
	_, err := DB.Exec(models.OrganizationQuery)
	if err != nil {
		log.Fatalf("Failed to create organizations table: %v", err)
	}

	_, err = DB.Exec(models.StudentQuery)
	if err != nil {
		log.Fatalf("Failed to create students table: %v", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS auth_tokens (
			auth_token VARCHAR(255) PRIMARY KEY
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create auth_tokens table: %v", err)
	}

	fmt.Println("Database migration completed successfully")
}