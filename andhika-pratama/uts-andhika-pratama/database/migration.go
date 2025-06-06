package database

import (
	"uts/models"

	"log"
)

func Migrate() {
	queries := []string{models.UserQuery, models.TokenQuery, models.ChallengeQuery, models.QuestionQuery, models.AnswerQuery, models.VoteQuery}
	for _, q := range queries {
		_, err := DB.Exec(q)
		if err != nil {
			log.Fatalf("Failed to migrate: %v", err)
		}
	}
}