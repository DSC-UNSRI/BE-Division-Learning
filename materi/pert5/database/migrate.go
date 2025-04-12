	package database

	import (
		"be_pert5/models"
		"fmt"
		"log"
	)

	func Migrate() {
		_, err := DB.Exec(models.BooksQuery)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
		fmt.Println("Migrate Success")
	}