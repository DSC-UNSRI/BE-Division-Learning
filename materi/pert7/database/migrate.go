package database

import (
	"be_pert5/models"
	"fmt"
	"log"
)

func Migrate() {
	query := []string{models.BooksQuery, models.UsersQuery}
	for _, q := range query {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}

	fmt.Println("Migrate Success")
}