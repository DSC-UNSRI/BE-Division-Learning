package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas-5/config"
	"tugas-5/database"
	"tugas-5/routes"
	"tugas-5/utils"
)

func main() {
	user, err := config.ENVLoad()
	if err != nil {
		log.Fatalf("Error loading env: %v", err)
	}

	if !utils.Authenticate(user) {
		log.Fatal("User authentication failed: user cannot be nil or invalid")
	}

	database.InitDB()
	defer database.DB.Close()
	database.Migrate()

	mux := http.NewServeMux()
	apiMux := http.NewServeMux()

	db := database.DB
	routes.RegisterRoutes(apiMux, db)

	mux.Handle("/api/", http.StripPrefix("/api", apiMux))

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
