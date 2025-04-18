package database

import (
	"be_pert5/models"
	"fmt"
	"log"
)

func Migrate() {
	queries := []string{models.ChefQuery, models.MenuQuery}
	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}
	fmt.Println("Migrate Success")
}