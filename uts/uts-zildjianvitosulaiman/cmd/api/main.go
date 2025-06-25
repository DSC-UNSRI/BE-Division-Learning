package main

import (
	"log"
	"net/http"
	"uts-zildjianvitosulaiman/migrations"
	"uts-zildjianvitosulaiman/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := database.NewMySQLConnection()
	defer db.Close()

	migrations.RunMigrations(db)

	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))
	RegisterRoutes(apiMux, db)

	port := "8080"
	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
