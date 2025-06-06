package controllers

import (
	"be_pert7/database"
	"be_pert7/models"
	"be_pert7/utils"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func GetMenus(w http.ResponseWriter, r *http.Request) {
	ctxChefID := r.Context().Value(utils.ChefIDKey).(string)
	rows, err := database.DB.Query("SELECT menu_id, menu_name, description, price, chef_id, category FROM menus WHERE chef_id = ? AND deleted_at IS NULL", ctxChefID)
	if err != nil {
		http.Error(w, "Failed to retrieve menus: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	menus := []models.Menu{}
	for rows.Next() {
		menu := models.Menu{}
		if err := rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category); err != nil {
			http.Error(w, "Failed to scan menu data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		menus = append(menus, menu)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"menus": menus,
	})
}

func GetMenuByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Menu ID is required", http.StatusBadRequest)
		return
	}

	var menu models.Menu
	err := database.DB.QueryRow("SELECT menu_id, menu_name, description, price, chef_id, category FROM menus WHERE menu_id = ? AND deleted_at IS NULL", id).Scan(
		&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Menu not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"menu": menu,
	})
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	menu := models.Menu{}

	ctxChefIDStr := r.Context().Value(utils.ChefIDKey).(string)
	ctxChefID, err := strconv.Atoi(ctxChefIDStr)
	if err != nil {
		http.Error(w, "Invalid Chef ID in context or not authenticated", http.StatusInternalServerError)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	menu.Name = r.FormValue("menu_name")
	menu.Description = r.FormValue("description")

	priceStr := r.FormValue("price")
	if priceStr == "" {
		http.Error(w, "Price is required", http.StatusBadRequest)
		return
	}
	if menu.Price, err = strconv.Atoi(priceStr); err != nil {
		http.Error(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	menu.ChefID = ctxChefID
	menu.Category = r.FormValue("category")

	// Basic validation for required fields
	if menu.Name == "" || menu.Description == "" || menu.Category == "" {
		http.Error(w, "Menu name, description, and category are required", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO menus (menu_name, description, price, chef_id, category) VALUES (?, ?, ?, ?, ?)", menu.Name, menu.Description, menu.Price, menu.ChefID, menu.Category)
	if err != nil {
		http.Error(w, "Failed to create menu: "+err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	menu.MenuID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) 
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Menu successfully created",
		"menu":    menu,
	})
}

func UpdateMenu(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Menu ID is required", http.StatusBadRequest)
		return
	}

	ctxChefIDStr := r.Context().Value(utils.ChefIDKey).(string)
	ctxChefID, err := strconv.Atoi(ctxChefIDStr)
	if err != nil {
		http.Error(w, "Invalid Chef ID in context or not authenticated", http.StatusInternalServerError)
		return
	}

	menu := models.Menu{}

	if err := database.DB.QueryRow("SELECT menu_id, menu_name, description, price, chef_id, category FROM menus WHERE menu_id = ? AND deleted_at IS NULL", id).Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Menu not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if menu.ChefID != ctxChefID {
		http.Error(w, "Unauthorized: You can only update your own menus", http.StatusForbidden)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if name := r.FormValue("menu_name"); name != "" {
		menu.Name = name
	}
	if description := r.FormValue("description"); description != "" {
		menu.Description = description
	}
	if priceStr := r.FormValue("price"); priceStr != "" {
		if price, err := strconv.Atoi(priceStr); err != nil {
			http.Error(w, "Invalid price format", http.StatusBadRequest)
			return
		} else {
			menu.Price = price
		}
	}

	if category := r.FormValue("category"); category != "" {
		menu.Category = category
	}

	_, err = database.DB.Exec("UPDATE menus SET menu_name = ?, description = ?, price = ?, category = ? WHERE menu_id = ? AND deleted_at IS NULL", menu.Name, menu.Description, menu.Price, menu.Category, id)
	if err != nil {
		http.Error(w, "Failed to update menu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Menu updated successfully",
		"menu":    menu,
	})
}

func DeleteMenu(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "Menu ID is required", http.StatusBadRequest)
		return
	}

	ctxChefIDStr := r.Context().Value(utils.ChefIDKey).(string)
	ctxChefID, err := strconv.Atoi(ctxChefIDStr)
	if err != nil {
		http.Error(w, "Invalid Chef ID in context or not authenticated", http.StatusInternalServerError)
		return
	}

	var menuChefID int
	if err := database.DB.QueryRow("SELECT chef_id FROM menus WHERE menu_id = ? AND deleted_at IS NULL", id).Scan(&menuChefID); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Menu not found or already deleted", http.StatusNotFound)
		} else {
			http.Error(w, "Database error during menu ownership check: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}
	
	if menuChefID != ctxChefID {
		http.Error(w, "Unauthorized: You can only delete your own menus", http.StatusForbidden)
		return
	}
	
	if _, err := database.DB.Exec("UPDATE menus SET deleted_at = ? WHERE menu_id = ?", time.Now(), id); err != nil {
		http.Error(w, "Failed to delete menu: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Menu deleted successfully",
		"menu_id": id,
	})
}

func GetMenusByChef(w http.ResponseWriter, r *http.Request, chefID string) {
	if chefID == "" {
		http.Error(w, "Chef ID is required", http.StatusBadRequest)
		return
	}

	var exists bool
	if err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE chef_id = ? AND deleted_at IS NULL)", chefID).Scan(&exists); err != nil {
		http.Error(w, "Database error while validating chef: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Chef not found", http.StatusNotFound)
		return
	}

	rows, err := database.DB.Query(`
		SELECT menu_id, menu_name, description, price, category, chef_id 
		FROM menus 
		WHERE chef_id = ? AND deleted_at IS NULL`, chefID)

	if err != nil {
		http.Error(w, "Database error while retrieving menus: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	menus := []models.Menu{}
	for rows.Next() {
		menu := models.Menu{}
		if err := rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.Category, &menu.ChefID); err != nil {
			http.Error(w, "Failed to scan menu data: "+err.Error(), http.StatusInternalServerError)
			return
			}
		menus = append(menus, menu)
	}

	if len(menus) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Chef has no assigned menus",
			"chef_id": chefID,
			"menus":   []models.Menu{}, 
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"chef_id": chefID,
		"menus":   menus,
	})
}


func GetMenusByCategory(w http.ResponseWriter, r *http.Request, category string) {
	if category == "" {
		http.Error(w, "Category is required", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT menu_id, menu_name, description, price, chef_id, category 
		FROM menus 
		WHERE category = ? AND deleted_at IS NULL`, category)

	if err != nil {
		http.Error(w, "Database error while retrieving menus: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	menus := []models.Menu{}
	for rows.Next() {
		menu := models.Menu{}
		if err := rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category); err != nil {
			http.Error(w, "Failed to scan menu data: "+err.Error(), http.StatusInternalServerError)
			return
		}
		menus = append(menus, menu)
	}

	if len(menus) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK) 
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":  "No menus found for this category",
			"category": category,
			"menus":    []models.Menu{},
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"category": category,
		"menus":    menus,
	})
}