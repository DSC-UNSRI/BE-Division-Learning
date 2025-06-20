package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySQLConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database!")

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	return db
}
