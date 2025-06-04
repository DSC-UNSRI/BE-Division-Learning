package main

import (
	"log"
	"net/http"
	"Task_5/config"
	"Task_5/database"
	"Task_5/routes"
	"fmt"
	"os"
)

func main() {
	config.ENVLoad()
	database.InitDB()
	defer database.DB.Close()
	routes.NasabahRoutes()
	routes.TabunganRoutes()
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server error:", err)
	}
	fmt.Println("DB User:", os.Getenv("DB_USER"))
}
