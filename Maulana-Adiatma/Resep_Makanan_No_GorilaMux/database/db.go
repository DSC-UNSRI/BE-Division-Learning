package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Gagal load .env")
	}

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatal("Gagal koneksi DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal ping DB:", err)
	}

	fmt.Println("Database connected")
}
