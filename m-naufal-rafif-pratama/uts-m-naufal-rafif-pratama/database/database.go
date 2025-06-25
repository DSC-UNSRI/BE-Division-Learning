package database

import (
	"database/sql"
	"fmt"
	"uts/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	user := config.GetEnv("DB_USER", "root")
	pass := config.GetEnv("DB_PASSWORD", "")
	host := config.GetEnv("DB_HOST", "127.0.0.1")
	port := config.GetEnv("DB_PORT", "3306")
	dbname := config.GetEnv("DB_NAME", "utsgdg")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	if err != nil {
		log.Fatalf("Could not create database: %v", err)
	}
	DB.Close()
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to database after creation: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	log.Println("Database connection successful.")
} 