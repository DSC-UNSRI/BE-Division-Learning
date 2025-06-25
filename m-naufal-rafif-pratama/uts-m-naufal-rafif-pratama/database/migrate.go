package database

import (
	"log"
	"uts/models"
)

func Migrate() {
	if _, err := DB.Exec(models.UserQuery); err != nil {
		log.Fatalf("Failed to migrate users table: %v", err)
	}
	if _, err := DB.Exec(models.TokenQuery); err != nil {
		log.Fatalf("Failed to migrate token table: %v", err)
	}
	if _, err := DB.Exec(models.QuestionQuery); err != nil {
		log.Fatalf("Failed to migrate Question table: %v", err)
	}
	if _, err := DB.Exec(models.AnswerQuery); err != nil {
		log.Fatalf("Failed to migrate Answer table: %v", err)
	}

	log.Println("Migration successful")
}