package utils

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    dbUser := os.Getenv("DB_USERNAME")
    dbPass := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }

    err = db.Ping()
    if err != nil {
        log.Fatalf("Error connecting to database: %s", err)
    }

    DB = db
    log.Println("Successfully connected to the database!")
}

func CloseDB() {
    if DB != nil {
        DB.Close()
        log.Println("Database connection closed.")
    }
}