package controllers

import (
	"be_pert5/database"
	"be_pert5/models"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetChefs(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name, speciality, experience, username FROM chefs WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	chefs := []models.Chef{}
	for rows.Next() {
		chef := models.Chef{}
		if err := rows.Scan(&chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		chefs = append(chefs, chef)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"chefs": chefs})
}

func GetChefByID(w http.ResponseWriter, r *http.Request, id string) {
    if id == "" {
        http.Error(w, "Chef ID is required", http.StatusBadRequest)
        return
    }

    var chef models.Chef
    err := database.DB.QueryRow("SELECT id, name, speciality, experience, username FROM chefs WHERE id = ? AND deleted_at IS NULL", id).Scan(
        &chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username)

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Chef not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "chef": chef,
    })
}


func CreateChef(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	chef := models.Chef{
		Name:       r.FormValue("name"),
		Speciality: r.FormValue("speciality"),
	}

	chef.Experience, err = strconv.Atoi(r.FormValue("experience"))
	if err != nil {
		http.Error(w, "Invalid experience", http.StatusBadRequest)
		return
	}

	chef.Username = r.FormValue("username")
	chef.Password = r.FormValue("password")

	res, err := database.DB.Exec("INSERT INTO chefs (name, speciality, experience, username, password) VALUES (?, ?, ?, ?, ?)", chef.Name, chef.Speciality, chef.Experience, chef.Username, chef.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	chef.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Chef successfully created",
		"chef":    chef,
	})
}

func UpdateChef(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID is null", http.StatusBadRequest)
		return
	}

	chef := models.Chef{}
	err := database.DB.QueryRow("SELECT id, name, speciality, experience, username FROM chefs WHERE id = ? AND deleted_at IS NULL", id).
		Scan(&chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Chef not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	if name := r.FormValue("name"); name != "" {
		chef.Name = name
	}
	if speciality := r.FormValue("speciality"); speciality != "" {
		chef.Speciality = speciality
	}
	if experienceStr := r.FormValue("experience"); experienceStr != "" {
		chef.Experience, err = strconv.Atoi(experienceStr)
		if err != nil {
			http.Error(w, "Invalid experience", http.StatusBadRequest)
			return
		}
	}
	if username := r.FormValue("username"); username != "" {
		chef.Username = username
	}

	_, err = database.DB.Exec("UPDATE chefs SET name = ?, speciality = ?, experience = ?, username = ? WHERE id = ? AND deleted_at IS NULL", chef.Name, chef.Speciality, chef.Experience, chef.Username, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Chef updated successfully",
		"chef":    chef,
	})
}

func DeleteChef(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Chef not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE chefs SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete chef", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Chef deleted successfully",
		"id":      id,
	})
}
