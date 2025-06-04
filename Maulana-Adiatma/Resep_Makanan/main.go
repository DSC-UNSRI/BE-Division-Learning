package main

import (
	"log"
	"net/http"
	"percobaan3/database"
	"percobaan3/routes"
)

func main() {
	database.InitDB()

	router := routes.SetupRoutes()

	log.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
