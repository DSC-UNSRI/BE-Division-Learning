package database

import (
	"fmt"
	"log"
	"UTS_BE/models"
)

func Migrate() {
	schemas := []string{
		models.UserSchema,
		models.QuestionSchema,
		models.AnswerSchema,
	}

	for _, schema := range schemas {
		_, err := DB.Exec(schema)
		if err != nil {
			log.Fatalf("Error migrating schema: %v\nSchema:\n%s", err, schema)
		}
	}

	fmt.Println("✅ Database migration completed successfully.")
}