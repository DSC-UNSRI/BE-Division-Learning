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
	defer database.DB.Close()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}