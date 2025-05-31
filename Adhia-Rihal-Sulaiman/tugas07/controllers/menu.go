package controllers

import (
	"be_pert7/database"
	"be_pert7/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetMenus(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, description, price, chef_id, category FROM menus WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	menus := []models.Menu{}
	for rows.Next() {
		menu := models.Menu{}
		rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category)
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
    err := database.DB.QueryRow("SELECT id, name, description, price, chef_id, category FROM menus WHERE id = ? AND deleted_at IS NULL", id).Scan(
        &menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category)

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Menu not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
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
	err := r.ParseForm() //pakai Multipart jika ada file
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	menu.Name = r.FormValue("name")
	menu.Description = r.FormValue("description")
	menu.Price, err = strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}
	menu.ChefID, err = strconv.Atoi(r.FormValue("chef_id"))
	if err != nil {
		http.Error(w, "invalid chef_id", http.StatusBadRequest)
		return
	}
	menu.Category = r.FormValue("category")

	res, err := database.DB.Exec("INSERT INTO menus (name, description, price, chef_id, category) VALUES (?, ?, ?, ?, ?)", menu.Name, menu.Description, menu.Price, menu.ChefID, menu.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	menu.MenuID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "menu successfully created",
		"menu":    menu,
	})
}

func UpdateMenu(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	menu := models.Menu{}
	err := database.DB.QueryRow("SELECT id, name, description, price, chef_id, category FROM menus WHERE id = ? AND deleted_at IS NULL", id).Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Menu not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price, err := strconv.Atoi(r.FormValue("price"))
	if err != nil {
		http.Error(w, "invalid price", http.StatusBadRequest)
		return
	}
	chefID, err := strconv.Atoi(r.FormValue("chef_id"))
	if err != nil {
		http.Error(w, "invalid chef_id", http.StatusBadRequest)
		return
	}
	category := r.FormValue("category")

	if name != "" {
		menu.Name = name
	}
	if description != "" {
		menu.Description = description
	}
	if price != 0 {
		menu.Price = price
	}
	if chefID != 0 {
		menu.ChefID = chefID
	}
	if category != "" {
		menu.Category = category
	}

	_, err = database.DB.Exec("UPDATE menus SET name = ?, description = ?, price = ?, chef_id = ?, category = ? WHERE id = ? AND deleted_at IS NULL", menu.Name, menu.Description, menu.Price, menu.ChefID, menu.Category, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM menus WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "menu not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE menus SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete menu", http.StatusInternalServerError)
		return
	}
	

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Menu deleted successfully",
		"id":      id,
	})
}

func GetMenusByChef(w http.ResponseWriter, r *http.Request, chefID string) {

    var exists bool
    err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE id = ? AND deleted_at IS NULL)", chefID).Scan(&exists)
    if err != nil {
        http.Error(w, "Database error while validating chef", http.StatusInternalServerError)
        return
    }
    
    if !exists {
        http.Error(w, "Chef not found", http.StatusNotFound)
        return
    }

    rows, err := database.DB.Query(`
        SELECT id, name, description, price, category, chef_id 
        FROM menus 
        WHERE chef_id = ? AND deleted_at IS NULL`, chefID)
    
    if err != nil {
        http.Error(w, "Database error while retrieving menus", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    menus := []models.Menu{}
    for rows.Next() {
        menu := models.Menu{}
        err := rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.Category, &menu.ChefID)
        if err != nil {
            http.Error(w, "Failed to scan menu data", http.StatusInternalServerError)
            return
        }
        menus = append(menus, menu)
    }

    if len(menus) == 0 {
        http.Error(w, "Chef has no assigned menus", http.StatusOK)
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
        SELECT id, name, description, price, chef_id, category 
        FROM menus 
        WHERE category = ? AND deleted_at IS NULL`, category)

    if err != nil {
        http.Error(w, "Database error while retrieving menus", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    menus := []models.Menu{}
    for rows.Next() {
        menu := models.Menu{}
        err := rows.Scan(&menu.MenuID, &menu.Name, &menu.Description, &menu.Price, &menu.ChefID, &menu.Category)
        if err != nil {
            http.Error(w, "Failed to scan menu data", http.StatusInternalServerError)
            return
        }
        menus = append(menus, menu)
    }

    if len(menus) == 0 {
        http.Error(w, "No menus found for this category", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "category": category,
        "menus":    menus,
    })
}



