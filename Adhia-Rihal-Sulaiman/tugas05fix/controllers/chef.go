
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
		rows.Scan(&chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username)
		chefs = append(chefs, chef)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"chefs": chefs,
	})
}

func CreateChef(w http.ResponseWriter, r *http.Request) {
	chef := models.Chef{}
	err := r.ParseForm() //pakai Multipart jika ada file 
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	chef.Name = r.FormValue("name")
	chef.Speciality = r.FormValue("speciality")


	chef.Experience, err = strconv.Atoi(r.FormValue("experience"))
	if err != nil {
		http.Error(w, "invalid experience", http.StatusBadRequest)
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
		"message": "chef successfully created",
		"chef":    chef,
	})
}

func UpdateChef(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	chef := models.Chef{}
	err := database.DB.QueryRow("SELECT id, name, speciality, experience, username FROM chefs WHERE id = ? AND deleted_at IS NULL", id).Scan(&chef.ID, &chef.Name, &chef.Speciality, &chef.Experience, &chef.Username)
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
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	speciality := r.FormValue("speciality")

	// Convert experience to integer
	experience, err := strconv.Atoi(r.FormValue("experience"))
	if err != nil {
		http.Error(w, "invalid experience", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")

	if name != "" {
		chef.Name = name
	}
	if speciality != "" {
		chef.Speciality = speciality
	}
	if experience != 0 {
		chef.Experience = experience
	}
	if username != "" {
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
		http.Error(w, "id is null", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM chefs WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "chef not found", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE chefs SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete chef", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Chef deleted successfully",
		"id":      id,
	})
}
