package main

import (
	"log"
	"net/http"
	"uts-zildjianvitosulaiman/pkg/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	db := database.NewMySQLConnection()
	defer db.Close()

	router := RegisterRoutes(db)

	port := "8080"
	log.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
