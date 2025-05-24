package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"pmm/config"
	"pmm/database"
	"pmm/routes"
)

func main() {
	config.ENVLoad()

	database.InitDB()
	if database.DB != nil {
		defer database.DB.Close()
	} else {
		log.Fatal("Database connection is nil after InitDB")
	}

	database.Migrate()

	routes.InitRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running at http://localhost:%s/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}