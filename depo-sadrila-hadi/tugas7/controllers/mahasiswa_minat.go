package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"tugas7/database"
	"tugas7/middleware"
	"tugas7/models"
	"tugas7/utils"
	"strconv"
	"strings"
)

func AddMinatToMahasiswa(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string) {
	currentUser, ok := middleware.GetCurrentUserFromContext(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusForbidden, "Authentication required or context error")
		return
	}

	var targetMahasiswaID int
	var err error

	if strings.ToLower(mahasiswaIDStr) == "me" {
		targetMahasiswaID = currentUser.ID
	} else {
		targetMahasiswaID, err = strconv.Atoi(mahasiswaIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Mahasiswa ID format in path")
			return
		}
	}

	if currentUser.ID != targetMahasiswaID {
		utils.RespondWithError(w, http.StatusForbidden, "Forbidden: You can only add interests to your own profile.")
		return
	}

	var input struct {
		MinatID int `json:"minat_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format, expected {'minat_id': id}")
		return
	}
	defer r.Body.Close()

	if input.MinatID <= 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Minat ID provided in request body")
		return
	}

	var mhsExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM mahasiswa WHERE id = ? AND deleted_at IS NULL)", targetMahasiswaID).Scan(&mhsExists)
	if err != nil || !mhsExists {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found (or check failed)")
		return
	}

	var minatExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM minat WHERE id = ? AND deleted_at IS NULL)", input.MinatID).Scan(&minatExists)
	if err != nil || !minatExists {
		utils.RespondWithError(w, http.StatusNotFound, "Minat not found or is deleted")
		return
	}

	query := "INSERT INTO mahasiswa_minat (mahasiswa_id, minat_id) VALUES (?, ?)"
	_, err = database.DB.Exec(query, targetMahasiswaID, input.MinatID)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "Mahasiswa already has this minat")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to add minat to mahasiswa")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "Minat added to mahasiswa successfully"})
}

func GetMinatByMahasiswaID(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string) {
	currentUser, ok := middleware.GetCurrentUserFromContext(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusForbidden, "Authentication required or context error")
		return
	}

	var targetMahasiswaID int
	var err error

	if strings.ToLower(mahasiswaIDStr) == "me" {
		targetMahasiswaID = currentUser.ID
	} else {
		targetMahasiswaID, err = strconv.Atoi(mahasiswaIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Mahasiswa ID format in path")
			return
		}
	}

	var mhsExists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM mahasiswa WHERE id = ? AND deleted_at IS NULL)", targetMahasiswaID).Scan(&mhsExists)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error checking mahasiswa existence")
		return
	}
	if !mhsExists {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found")
		return
	}

	query := `
		SELECT m.id, m.nama_minat, m.deskripsi, m.created_at, m.updated_at
		FROM minat m
		JOIN mahasiswa_minat mm ON m.id = mm.minat_id
		WHERE mm.mahasiswa_id = ? AND m.deleted_at IS NULL
		ORDER BY m.nama_minat ASC
	`
	rows, err := database.DB.Query(query, targetMahasiswaID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database query error fetching interests for mahasiswa")
		return
	}
	defer rows.Close()

	minats := []models.Minat{}
	for rows.Next() {
		var m models.Minat
		var deskripsi sql.NullString
		if err := rows.Scan(&m.ID, &m.NamaMinat, &deskripsi, &m.CreatedAt, &m.UpdatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan minat row for mahasiswa")
			return
		}
		if deskripsi.Valid {
			m.Deskripsi = deskripsi.String
		} else {
			m.Deskripsi = ""
		}
		minats = append(minats, m)
	}
	if err = rows.Err(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Row iteration error after fetching interests for mahasiswa")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"minat_mahasiswa": minats})
}

func RemoveMinatFromMahasiswa(w http.ResponseWriter, r *http.Request, mahasiswaIDStr string, minatIDStr string) {
	currentUser, ok := middleware.GetCurrentUserFromContext(r.Context())
	if !ok {
		utils.RespondWithError(w, http.StatusForbidden, "Authentication required or context error")
		return
	}

	var targetMahasiswaID int
	var err error

	if strings.ToLower(mahasiswaIDStr) == "me" {
		targetMahasiswaID = currentUser.ID
	} else {
		targetMahasiswaID, err = strconv.Atoi(mahasiswaIDStr)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "Invalid Mahasiswa ID format in path")
			return
		}
	}

	minatID, err := strconv.Atoi(minatIDStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Minat ID format in path")
		return
	}

	if currentUser.ID != targetMahasiswaID {
		utils.RespondWithError(w, http.StatusForbidden, "Forbidden: You can only remove interests from your own profile.")
		return
	}

	query := "DELETE FROM mahasiswa_minat WHERE mahasiswa_id = ? AND minat_id = ?"
	res, err := database.DB.Exec(query, targetMahasiswaID, minatID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to remove minat from mahasiswa")
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check affected rows after removing interest")
		return
	}

	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Minat relationship not found for this mahasiswa, or already removed")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Minat removed from mahasiswa successfully"})
}