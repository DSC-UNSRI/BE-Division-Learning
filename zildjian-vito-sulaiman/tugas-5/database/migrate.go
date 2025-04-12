package database

import (
	"fmt"
	"log"
	"tugas-5/models"
)

func Migrate() {
	_, err := DB.Exec(models.UsersQuery)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	fmt.Println("Migrate Success")
}
