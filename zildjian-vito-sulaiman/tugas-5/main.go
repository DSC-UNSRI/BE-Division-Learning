package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas-5/config"
	"tugas-5/database"
	"tugas-5/routes"
)

func main() {
	config.ENVLoad()

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
