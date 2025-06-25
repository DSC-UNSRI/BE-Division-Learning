package database

import (
	"pertemuan05/models"

	"fmt"
	"log"
)

func Migrate() {
	queries := []string{models.LecturerQuery, models.CourseQuery, models.AddressQuery}
	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}
	fmt.Println("Migrate Success")
}