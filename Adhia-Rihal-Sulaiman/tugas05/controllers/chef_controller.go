package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas05/config"
	"tugas05/models"
)

type ChefController struct {
	db *sql.DB
}

func NewChefController() *ChefController {
	return &ChefController{db: db}
}

// Create Chef
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
	err := config.DB.QueryRow("SELECT COUNT(*) FROM chefs WHERE username = ?", chef.Username).Scan(&count)
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

	result, err := config.DB.Exec(query,
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

// Login Chef - Without Token (simple username and password check)
func (c *ChefController) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Decode request body
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&loginRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check credentials directly without using JWT or token
	var chef models.Chef
	query := `SELECT id, name, username, speciality, experience, password_hash
			  FROM chefs WHERE username = ? AND password_hash = ?`

	err := config.DB.QueryRow(query, loginRequest.Username, loginRequest.Password).Scan(
		&chef.ID,
		&chef.Name,
		&chef.Username,
		&chef.Speciality,
		&chef.Experience,
		&chef.PasswordHash,
	)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Remove password before sending response
	chef.PasswordHash = ""

	// Send response (return chef details after login)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chef)
}

// GetAll Chef
func (c *ChefController) GetAll(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, speciality, experience, username FROM chefs"
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var chefs []models.Chef
	for rows.Next() {
		var chef models.Chef
		err := rows.Scan(
			&chef.ID,
			&chef.Name,
			&chef.Speciality,
			&chef.Experience,
			&chef.Username,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		chefs = append(chefs, chef)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chefs)
}

// GetByID Chef
func (c *ChefController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chefs/"):]

	var chef models.Chef
	query := `SELECT id, name, speciality, experience, username
			  FROM chefs WHERE id = ?`

	err := config.DB.QueryRow(query, id).Scan(
		&chef.ID,
		&chef.Name,
		&chef.Speciality,
		&chef.Experience,
		&chef.Username,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Chef not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chef)
}

// Update Chef
func (c *ChefController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chefs/"):]

	var chef models.Chef
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&chef); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Query update
	query := `UPDATE chefs
			  SET name = ?, speciality = ?, experience = ?
			  WHERE id = ?`

	_, err := config.DB.Exec(query,
		chef.Name,
		chef.Speciality,
		chef.Experience,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chef)
}

// Delete Chef
func (c *ChefController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/chefs/"):]

	// Query delete
	query := "DELETE FROM chefs WHERE id = ?"

	result, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Periksa apakah ada baris yang terpengaruh
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Chef not found", http.StatusNotFound)
		return
	}

	// Kirim response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Chef deleted successfully"}`))
}
