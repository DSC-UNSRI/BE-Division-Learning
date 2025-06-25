package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"UTS_BE/config"
	"UTS_BE/database"
	"UTS_BE/routes"
)

func main() {
	// Load environment variables
	config.ENVLoad()

	// Init DB and migrate schema
	database.InitDB()
	database.Migrate()

	// Register routes
	routes.AuthRoutes()
	routes.QuestionRoutes()
	routes.AnswerRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
