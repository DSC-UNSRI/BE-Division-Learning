package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"tugas05/models"
)

type MenuController struct {
	db *sql.DB
}

func NewMenuController(db *sql.DB) *MenuController {
	return &MenuController{db: db}
}

// Create menu (menggunakan http.Request dan http.ResponseWriter)
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

// GetAll menu (menggunakan http.Request dan http.ResponseWriter)
func (c *MenuController) GetAll(w http.ResponseWriter, r *http.Request) {
	query := `SELECT id, name, description, price, chef_id, category FROM menus`
	rows, err := c.db.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var menus []models.Menu
	for rows.Next() {
		var menu models.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Description,
			&menu.Price,
			&menu.ChefID,
			&menu.Category,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		menus = append(menus, menu)
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menus)
}

// GetByID menu (menggunakan http.Request dan http.ResponseWriter)
func (c *MenuController) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var menu models.Menu
	query := `SELECT id, name, description, price, chef_id, category
			  FROM menus WHERE id = ?`

	err := c.db.QueryRow(query, id).Scan(
		&menu.ID,
		&menu.Name,
		&menu.Description,
		&menu.Price,
		&menu.ChefID,
		&menu.Category,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Menu not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

// Update menu (menggunakan http.Request dan http.ResponseWriter)
func (c *MenuController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	var menu models.Menu
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validasi input
	if menu.Name == "" || menu.Price <= 0 {
		http.Error(w, "Invalid menu details", http.StatusBadRequest)
		return
	}

	// Query update
	query := `UPDATE menus
			  SET name = ?, description = ?, price = ?, chef_id = ?, category = ?
			  WHERE id = ?`

	_, err := c.db.Exec(query,
		menu.Name,
		menu.Description,
		menu.Price,
		menu.ChefID,
		menu.Category,
		id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set ID dari parameter
	menu.ID, _ = strconv.Atoi(id)

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

// Delete menu (menggunakan http.Request dan http.ResponseWriter)
func (c *MenuController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// Query delete
	query := "DELETE FROM menus WHERE id = ?"

	result, err := c.db.Exec(query, id)
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
		http.Error(w, "Menu not found", http.StatusNotFound)
		return
	}

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	http.Error(w, "Menu deleted suscessfully", http.StatusOK)
}
