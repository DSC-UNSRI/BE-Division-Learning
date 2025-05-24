package main

import (
	"fmt"
	"log"
	"net/http"
	"tugas/todolist/config"
	"tugas/todolist/database"
	"tugas/todolist/routes"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.LoadEnv()

	fmt.Println("Starting the application...")

	db := database.Connect()
	defer db.Close()

	router := routes.SetupRoutes(db)

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
