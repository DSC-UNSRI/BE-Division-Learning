package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas05/models"
	"strconv"
)

type MenuController struct {
	db *sql.DB
}

func NewMenuController() *MenuController {
	return &MenuController{db: db}
}

// Create Menu
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

// GetAll Menu
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menus)
}

// Update Menu
func (c *MenuController) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/menus/"):]

	var menu models.Menu
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&menu); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	// Kirim response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menu)
}

// Delete Menu
func (c *MenuController) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/menus/"):]

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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Menu deleted successfully"}`))
}

// SearchMenus
func (c *MenuController) SearchMenus(w http.ResponseWriter, r *http.Request) {
	// Parameter pencarian
	name := r.URL.Query().Get("name")
	category := r.URL.Query().Get("category")
	minPrice := r.URL.Query().Get("min_price")
	maxPrice := r.URL.Query().Get("max_price")

	// Query dasar
	query := `SELECT id, name, description, price, chef_id, category
			  FROM menus WHERE 1=1`

	var args []interface{}

	// Tambahkan kondisi pencarian
	if name != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+name+"%")
	}

	if category != "" {
		query += " AND category = ?"
		args = append(args, category)
	}

	if minPrice != "" {
		minPriceFloat, _ := strconv.ParseFloat(minPrice, 64)
		query += " AND price >= ?"
		args = append(args, minPriceFloat)
	}

	if maxPrice != "" {
		maxPriceFloat, _ := strconv.ParseFloat(maxPrice, 64)
		query += " AND price <= ?"
		args = append(args, maxPriceFloat)
	}

	// Eksekusi query
	rows, err := c.db.Query(query, args...)
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

	// Cek apakah daftar menu kosong
	if len(menus) == 0 {
		http.Error(w, "No menus found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menus)
}

// GetMenusByChef
func (c *MenuController) GetMenusByChef(w http.ResponseWriter, r *http.Request) {
	chefID := r.URL.Path[len("/menus/chef/"):]

	query := `SELECT id, name, description, price, chef_id, category
			  FROM menus WHERE chef_id = ?`

	rows, err := c.db.Query(query, chefID)
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

	// Cek apakah daftar menu kosong
	if len(menus) == 0 {
		http.Error(w, "No menus found for this chef", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(menus)
}
