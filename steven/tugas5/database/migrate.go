package database

import (
	"tugas5/models"
	"fmt"
	"log"
)

func Migrate() {
	_, err := DB.Exec(models.ProductsQuery)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	fmt.Println("Migrate Success")
}