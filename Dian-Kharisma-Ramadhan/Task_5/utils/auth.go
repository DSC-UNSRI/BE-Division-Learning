package utils

import (
	"database/sql"
)

func SimpleAuth(db *sql.DB, nama string, password string) bool {
	var storedPassword string
	err := db.QueryRow("SELECT password FROM nasabah WHERE nama = ?", nama).Scan(&storedPassword)
	if err != nil || storedPassword != password {
		return false
	}
	return true
}
