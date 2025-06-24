package database

import (
	"uts-gdg/models"
	"fmt"
	"log"
)

func Migrate() {
	queries := []string{
		models.UserQuery,
		models.QuestionQuery,
		models.AnswerQuery,
	}

	for _, query := range queries {
		_, err := DB.Exec(query)
		if err != nil {
			log.Fatalf("Failed to execute migration: %v", err)
		}
	}
	fmt.Println("Migrate Success")
}