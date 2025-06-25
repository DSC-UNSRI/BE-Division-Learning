package main

import (
	"log"
	"net/http"
	"uts/config"
	"uts/database"
	"uts/routes"
)

func main(){
	config.LoadEnv()
	database.Connect()
	database.Migrate()

	mux := routes.SetupRoutes()
	
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server starting on port %s\n", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}