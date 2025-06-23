package main

import (
	"uts-gdg/config"
	"uts-gdg/database"
	"uts-gdg/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	database.Migrate()
	routes.UserRoutes()
	routes.QuestionRoutes()

	fmt.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}