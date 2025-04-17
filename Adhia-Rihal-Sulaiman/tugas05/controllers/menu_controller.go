package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas05/models"

)

type MenuController struct {
	db *sql.DB
}

func NewMenuController(db *sql.DB) *MenuController {
	return &MenuController{db: db}
}

// Create Menu (menggunakan http.Request dan http.ResponseWriter)
func (c *MenuController) Create(w http.ResponseWriter, r *http.Request) {
	var menu models.Menu
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi input
	if menu.Name == "" || menu.Price <= 0 || menu.ChefID <= 0 {
		http.Error(w, "Invalid menu details", http.StatusBadRequest)
		return
	}

	// Cek apakah chef tersedia
	var chefExists int
	err := c.db.QueryRow("SELECT COUNT(*) FROM chefs WHERE id = ?", menu.ChefID).Scan(&chefExists)
	if err != nil || chefExists == 0 {
		http.Error(w, "Chef not found", http.StatusBadRequest)
		return
	}

	// Query insert
	query := `INSERT INTO menus
		(name, description, price, chef_id, category)
		VALUES (?, ?, ?, ?, ?)`

	result, err := c.db.Exec(query,
		menu.Name,
		menu.Description,
		menu.Price,
		menu.ChefID,
		menu.Category,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Dapatkan ID yang baru dibuat
	id, _ := result.LastInsertId()
	menu.ID = int(id)

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(menu)
}
