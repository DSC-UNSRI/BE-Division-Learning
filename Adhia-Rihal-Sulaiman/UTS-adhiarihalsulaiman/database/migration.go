package database

import (
	"uts_adhia/models"

	"log"
)

func Migrate() {
	queries := []string{models.UserTableSchema, models.QuestionTableSchema, models.AnswerTableSchema, models.ForgetTableSchema, models.HighlightTableSchema, models.TokenTableSchema, models.UserTableSchema}
	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}
}
