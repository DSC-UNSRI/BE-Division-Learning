package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"Tugas-Pertemuan-7/database"
	"Tugas-Pertemuan-7/models"
	"Tugas-Pertemuan-7/utils"
)

func GetFilms(w http.ResponseWriter, r *http.Request) {
	directorID, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Internal server error: Director ID not found in context", http.StatusInternalServerError)
		return
	}

	rows, err := database.DB.Query("SELECT id, title, director_id FROM films WHERE director_id = ? AND deleted_at IS NULL", directorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	films := []models.Film{}
	for rows.Next() {
		film := models.Film{}
		if err := rows.Scan(&film.ID, &film.Title, &film.DirectorID); err != nil {
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
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	film := models.Film{}
	err := database.DB.QueryRow("SELECT id, title, director_id FROM films WHERE id = ? AND deleted_at IS NULL", id).Scan(&film.ID, &film.Title, &film.DirectorID)
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
	directorID, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized: Director ID not found in context.", http.StatusUnauthorized)
		return
	}

	film := models.Film{}
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	film.Title = r.FormValue("title")

	if film.Title == "" {
		http.Error(w, "Title is a required field", http.StatusBadRequest)
		return
	}

	film.DirectorID = directorID

	res, err := database.DB.Exec("INSERT INTO films (title, director_id) VALUES (?, ?)", film.Title, film.DirectorID)
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
		"message": "Film successfully created",
		"film":    film,
	})
}

func UpdateFilm(w http.ResponseWriter, r *http.Request, id string) {
	if id == "" {
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	directorIDFromContext, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized: Director ID not found in context.", http.StatusUnauthorized)
		return
	}

	film := models.Film{}
	err := database.DB.QueryRow("SELECT id, title, director_id FROM films WHERE id = ? AND deleted_at IS NULL", id).Scan(&film.ID, &film.Title, &film.DirectorID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if film.DirectorID != directorIDFromContext {
		http.Error(w, "Forbidden: You can only update films you created.", http.StatusForbidden)
		return
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")

	updateFields := make(map[string]interface{})
	if title != "" && title != film.Title {
		updateFields["title"] = title
		film.Title = title
	}

	if len(updateFields) > 0 {
		query := "UPDATE films SET "
		params := []interface{}{}
		first := true
		for key, val := range updateFields {
			if !first {
				query += ", "
			}
			query += key + " = ?"
			params = append(params, val)
			first = false
		}
		query += " WHERE id = ? AND deleted_at IS NULL"
		params = append(params, id)

		_, err = database.DB.Exec(query, params...)
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
		http.Error(w, "ID parameter is missing", http.StatusBadRequest)
		return
	}

	directorIDFromContext, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized: Director ID not found in context.", http.StatusUnauthorized)
		return
	}

	film := models.Film{}
	err := database.DB.QueryRow("SELECT id, title, director_id FROM films WHERE id = ? AND deleted_at IS NULL", id).Scan(&film.ID, &film.Title, &film.DirectorID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Film not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if film.DirectorID != directorIDFromContext {
		http.Error(w, "Forbidden: You can only delete films you created.", http.StatusForbidden)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM films WHERE id = ? AND deleted_at IS NULL)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error checking film existence", http.StatusInternalServerError)
		return
	}

	if !exists {
		http.Error(w, "Film not found or already deleted", http.StatusNotFound)
		return
	}

	_, err = database.DB.Exec("UPDATE films SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete film", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Film deleted successfully",
		"id":      id,
	})
}

func GetFilmsByDirectorID(w http.ResponseWriter, r *http.Request, directorID string) {
	if directorID == "" {
		http.Error(w, "Director ID parameter is missing", http.StatusBadRequest)
		return
	}

	directorIDFromContext, ok := r.Context().Value(utils.DirectorIDKey).(int)
	if !ok {
		http.Error(w, "Internal server error: Director ID not found in context", http.StatusInternalServerError)
		return
	}
	directorRoleFromContext, ok := r.Context().Value(utils.RoleKey).(string)
	if !ok {
		http.Error(w, "Internal server error: Director role not found in context", http.StatusInternalServerError)
		return
	}

	parsedDirectorID, err := strconv.Atoi(directorID)
	if err != nil {
		http.Error(w, "Invalid director ID format", http.StatusBadRequest)
		return
	}

	// Logic for role-based access
	if directorRoleFromContext != "admin" && directorIDFromContext != parsedDirectorID {
		http.Error(w, "Forbidden: You can only view films by your own director ID unless you are an admin.", http.StatusForbidden)
		return
	}

	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM directors WHERE id = ? AND deleted_at IS NULL)", directorID).Scan(&exists)
	if err != nil {
		http.Error(w, "Failed to check director existence", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Director not found", http.StatusNotFound)
		return
	}

	rows, err := database.DB.Query("SELECT id, title, director_id FROM films WHERE director_id = ? AND deleted_at IS NULL", directorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	films := []models.Film{}
	for rows.Next() {
		film := models.Film{}
		if err := rows.Scan(&film.ID, &film.Title, &film.DirectorID); err != nil {
			http.Error(w, "Failed to scan film data for director", http.StatusInternalServerError)
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