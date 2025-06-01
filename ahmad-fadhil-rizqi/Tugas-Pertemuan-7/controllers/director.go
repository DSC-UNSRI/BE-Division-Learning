package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Tugas-Pertemuan-7/database"
	"Tugas-Pertemuan-7/models"
)

func GetDirectors(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, name FROM directors WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	directors := []models.Director{}
	for rows.Next() {
		director := models.Director{}
		if err := rows.Scan(&director.ID, &director.Name); err != nil {
			http.Error(w, "Failed to scan director data", http.StatusInternalServerError)
			return
		}
		directors = append(directors, director)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"directors": directors,
	})
}

func GetDirectorByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	director := models.Director{}
	err := database.DB.QueryRow("SELECT id, name FROM directors WHERE id = ? AND deleted_at IS NULL", id).Scan(&director.ID, &director.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	director := models.Director{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	director.Name = r.FormValue("name")

	if director.Name == "" {
		http.Error(w, "name is a required field", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO directors (name) VALUES (?)", director.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}
	director.ID = int(resID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "director successfully created",
		"director": director,
	})
}

func UpdateDirector(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	director := models.Director{}
	err := database.DB.QueryRow("SELECT id, name FROM directors WHERE id = ? AND deleted_at IS NULL", id).Scan(&director.ID, &director.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director not found", http.StatusNotFound)
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
	updated := false
	if name != "" && name != director.Name {
		director.Name = name
		updated = true
	}

	if updated {
		_, err = database.DB.Exec("UPDATE directors SET name = ? WHERE id = ? AND deleted_at IS NULL", director.Name, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Director updated successfully",
		"director": director,
	})
}

func DeleteDirector(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM directors WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error checking director existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "director not found or already deleted", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE directors SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Director deleted successfully",
		"id":      id,
	})
}

func GetDirectorByFilmID(w http.ResponseWriter, r *http.Request, filmID string) {
	if filmID == "" {
		http.Error(w, "film_id parameter is missing", http.StatusBadRequest)
		return
	}

	var directorID int
	err := database.DB.QueryRow("SELECT director_id FROM films WHERE id = ? AND deleted_at IS NULL", filmID).Scan(&directorID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	director := models.Director{}
	err = database.DB.QueryRow("SELECT id, name FROM directors WHERE id = ? AND deleted_at IS NULL", directorID).Scan(&director.ID, &director.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Director associated with this film not found or has been deleted", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(director)
}