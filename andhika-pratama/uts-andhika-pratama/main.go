package main

import (
	"uts/config"
	"uts/database"

	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	fmt.Println("Database Connected")
	defer database.DB.Close()

	database.Migrate()
	fmt.Println("Migration Success")

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}