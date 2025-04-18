package database

import (
	"fmt"
	"log"
	"tugas-5/models"
)

func Migrate() {
	queries := []string{
		models.UsersQuery,
		models.ProgrammerQuery,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}

	fmt.Println("Migration success for all tables.")
}
