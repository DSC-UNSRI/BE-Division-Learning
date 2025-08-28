package database

import (
	"UTS-Ahmad-Fadhil-Rizqi/models"
	"log"
)

func Migrate() {
	_, err := DB.Exec(models.CreateUsersTableQuery)
	if err != nil { log.Fatalf("Gagal migrasi tabel users: %v", err) }

	_, err = DB.Exec(models.CreateQuestionsTableQuery)
	if err != nil { log.Fatalf("Gagal migrasi tabel questions: %v", err) }

	_, err = DB.Exec(models.CreateAnswersTableQuery)
	if err != nil { log.Fatalf("Gagal migrasi tabel answers: %v", err) }

	_, err = DB.Exec(models.CreateSecurityQuestionsTableQuery)
	if err != nil { log.Fatalf("Gagal migrasi tabel security_questions: %v", err) }

	
	log.Println("Semua tabel berhasil dimigrasi.")
}