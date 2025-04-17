package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"tugas05/routes"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Menghubungkan ke database
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/restaurant_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Setup routing
	routes.SetupRoutes(db)

	// Jalankan server HTTP
	port := ":8080"
	fmt.Println("Server running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
