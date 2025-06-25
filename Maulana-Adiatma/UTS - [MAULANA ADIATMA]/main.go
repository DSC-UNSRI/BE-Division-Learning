package main

import (
	"log"
	"net/http"
	
	"utsquora/database"
	"utsquora/routes"
)

func main() {
	database.InitDB()

	routes.SetupRoutes()
	http.ListenAndServe(":8080", nil)

	log.Println("Server running on port 8080...")
}
