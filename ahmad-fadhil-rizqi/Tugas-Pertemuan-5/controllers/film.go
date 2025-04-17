package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"Tugas-Pertemuan-5/database"
	"Tugas-Pertemuan-5/models"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, title, director FROM films WHERE deleted_at IS NULL")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	films := []models.Film{}
	for rows.Next() {
		film := models.Film{}
		if err := rows.Scan(&film.ID, &film.Title, &film.Director); err != nil {
			http.Error(w, "Failed to scan film data", http.StatusInternalServerError)
			return
		}
		films = append(films, film)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"films": films,
	})
}

func GetFilmByID(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	film := models.Film{}
	err := database.DB.QueryRow("SELECT id, title, director FROM films WHERE id = ? AND deleted_at IS NULL", id).Scan(&film.ID, &film.Title, &film.Director)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(film)
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	film := models.Film{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	film.Title = r.FormValue("title")
	film.Director = r.FormValue("director")

	if film.Title == "" || film.Director == "" {
		http.Error(w, "title and director are required fields", http.StatusBadRequest)
		return
	}

	res, err := database.DB.Exec("INSERT INTO films (title, director) VALUES (?, ?)", film.Title, film.Director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID", http.StatusInternalServerError)
		return
	}
	film.ID = int(resID)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "film successfully created",
		"film":    film,
	})
}

func UpdateFilm(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	film := models.Film{}
	err := database.DB.QueryRow("SELECT id, title, director FROM films WHERE id = ? AND deleted_at IS NULL", id).Scan(&film.ID, &film.Title, &film.Director)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
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

	title := r.FormValue("title")
	director := r.FormValue("director")
	updated := false
	if title != "" && title != film.Title {
		film.Title = title
		updated = true
	}
	if director != "" && director != film.Director {
		film.Director = director
		updated = true
	}

	if updated {
		_, err = database.DB.Exec("UPDATE films SET title = ?, director = ? WHERE id = ? AND deleted_at IS NULL", film.Title, film.Director, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Film updated successfully",
		"film":    film,
	})
}

func DeleteFilm(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "id parameter is missing", http.StatusBadRequest)
		return
	}

	var exists bool
	err := database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM films WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "database error checking film existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "film not found or already deleted", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE films SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "failed to delete film", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Film deleted successfully",
		"id":      id,
	})
}