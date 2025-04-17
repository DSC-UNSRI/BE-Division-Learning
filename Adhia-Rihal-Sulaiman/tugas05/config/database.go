package config

import (  
	"database/sql"  
	"fmt"  
	"log"  
	"os"  

	_ "github.com/go-sql-driver/mysql"  
	"github.com/joho/godotenv"  
)  

func InitDatabase() *sql.DB {  
	err := godotenv.Load()  
	if err != nil {  
		log.Fatal("Error loading .env file")  
	}  

	dbUser := os.Getenv("DB_USER")  
	dbPass := os.Getenv("DB_PASS")  
	dbHost := os.Getenv("DB_HOST")  
	dbPort := os.Getenv("DB_PORT")  
	dbName := os.Getenv("DB_NAME")  

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",   
		dbUser, dbPass, dbHost, dbPort, dbName)  

	db, err := sql.Open("mysql", connectionString)  
	if err != nil {  
		log.Fatal(err)  
	}  

	err = db.Ping()  
	if err != nil {  
		log.Fatal(err)  
	}  

	return db  
}  