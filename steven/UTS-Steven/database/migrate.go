package database

import (
	"uts-gdg/models"
	"fmt"
	"log"
)

func Migrate() {
	_, err := DB.Exec(models.UserQuery)
	if err != nil {
		log.Fatalf("Failed to migrate: %v", err)
	}
	fmt.Println("Migrate Success")
}