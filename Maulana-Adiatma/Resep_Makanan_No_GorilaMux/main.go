package main

import (
	"log"
	"net/http"
	"resepku/database"
	"resepku/routes"
)

func main() {
	database.InitDB()

	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)

	log.Println("Server running on port 8080...")
}
