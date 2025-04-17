package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas05/models"

)

type ChefController struct {
	db *sql.DB
}

func NewChefController(db *sql.DB) *ChefController {
	return &ChefController{db: db}
}

// Create Chef (menggunakan http.Request dan http.ResponseWriter)
func (c *ChefController) Create(w http.ResponseWriter, r *http.Request) {
	var chef models.Chef
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chef); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi input
	if chef.Name == "" || chef.Username == "" || chef.PasswordHash == "" {
		http.Error(w, "Name, username, and password are required", http.StatusBadRequest)
		return
	}

	// Cek apakah username sudah ada
	var count int
	err := c.db.QueryRow("SELECT COUNT(*) FROM chefs WHERE username = ?", chef.Username).Scan(&count)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Username already exists", http.StatusConflict)
		return
	}

	// Query insert
	query := `INSERT INTO chefs
		(name, speciality, experience, username, password_hash)
		VALUES (?, ?, ?, ?, ?)`

	result, err := c.db.Exec(query,
		chef.Name,
		chef.Speciality,
		chef.Experience,
		chef.Username,
		chef.PasswordHash,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Dapatkan ID yang baru dibuat
	id, _ := result.LastInsertId()
	chef.ID = int(id)

	// Hapus password sebelum mengirim response
	chef.PasswordHash = ""

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(chef)
}
