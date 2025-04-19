package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"pmm/database"
	"pmm/models"
	"pmm/utils"
)

func GetAllMinat(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	rows, err := database.DB.Query("SELECT id, nama_minat, deskripsi, created_at, updated_at FROM minat WHERE deleted_at IS NULL ORDER BY id ASC")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database query error fetching minat list")
		return
	}
	defer rows.Close()

	minats := []models.Minat{}
	for rows.Next() {
		var m models.Minat
		var deskripsi sql.NullString
		if err := rows.Scan(&m.ID, &m.NamaMinat, &deskripsi, &m.CreatedAt, &m.UpdatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan minat row")
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
		utils.RespondWithError(w, http.StatusInternalServerError, "Row iteration error after fetching minat list")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"minat": minats})
}

func GetMinatByID(w http.ResponseWriter, r *http.Request, idStr string) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	var m models.Minat
	var deskripsi sql.NullString
	query := "SELECT id, nama_minat, deskripsi, created_at, updated_at FROM minat WHERE id = ? AND deleted_at IS NULL"
	err = database.DB.QueryRow(query, id).Scan(&m.ID, &m.NamaMinat, &deskripsi, &m.CreatedAt, &m.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Minat not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Database query error fetching minat by ID")
		}
		return
	}
	if deskripsi.Valid {
		m.Deskripsi = deskripsi.String
	} else {
		m.Deskripsi = ""
	}

	utils.RespondWithJSON(w, http.StatusOK, m)
}

func CreateMinat(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	var m models.Minat
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}
	defer r.Body.Close()

	namaMinatTrimmed := strings.TrimSpace(m.NamaMinat)
	if namaMinatTrimmed == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Nama Minat is required and cannot be empty")
		return
	}
	deskripsiTrimmed := strings.TrimSpace(m.Deskripsi)

	query := "INSERT INTO minat (nama_minat, deskripsi) VALUES (?, ?)"
	res, err := database.DB.Exec(query, namaMinatTrimmed, deskripsiTrimmed)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "Nama Minat already exists")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create minat")
		}
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get last insert ID")
		return
	}

	var createdMinat models.Minat
	fetchQuery := "SELECT id, nama_minat, deskripsi, created_at, updated_at FROM minat WHERE id = ?"
	var fetchedDesc sql.NullString
	err = database.DB.QueryRow(fetchQuery, id).Scan(&createdMinat.ID, &createdMinat.NamaMinat, &fetchedDesc, &createdMinat.CreatedAt, &createdMinat.UpdatedAt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Minat created, but failed to retrieve full details")
		return
	}
	if fetchedDesc.Valid {
		createdMinat.Deskripsi = fetchedDesc.String
	} else {
		createdMinat.Deskripsi = ""
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdMinat)
}

func UpdateMinat(w http.ResponseWriter, r *http.Request, idStr string) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	var input struct {
		NamaMinat *string `json:"nama_minat"`
		Deskripsi *string `json:"deskripsi"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}
	defer r.Body.Close()

	if input.NamaMinat == nil && input.Deskripsi == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No fields provided to update (expected 'nama_minat' and/or 'deskripsi')")
		return
	}

	var setClauses []string
	var args []interface{}

	if input.NamaMinat != nil {
		namaTrimmed := strings.TrimSpace(*input.NamaMinat)
		if namaTrimmed == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Nama Minat cannot be updated to empty")
			return
		}
		setClauses = append(setClauses, "nama_minat = ?")
		args = append(args, namaTrimmed)
	}

	if input.Deskripsi != nil {
		deskripsiTrimmed := strings.TrimSpace(*input.Deskripsi)
		setClauses = append(setClauses, "deskripsi = ?")
		args = append(args, deskripsiTrimmed)
	}

	if len(setClauses) == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "No valid fields provided for update")
		return
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	args = append(args, id)

	queryUpdate := "UPDATE minat SET " + strings.Join(setClauses, ", ") + " WHERE id = ? AND deleted_at IS NULL"
	res, err := database.DB.Exec(queryUpdate, args...)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "Update failed: Nama Minat already exists")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update minat")
		}
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check affected rows after update")
		return
	}
	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Minat not found or already deleted, update failed")
		return
	}

	var updatedMinat models.Minat
	fetchQuery := "SELECT id, nama_minat, deskripsi, created_at, updated_at FROM minat WHERE id = ?"
	var fetchedDesc sql.NullString
	err = database.DB.QueryRow(fetchQuery, id).Scan(&updatedMinat.ID, &updatedMinat.NamaMinat, &fetchedDesc, &updatedMinat.CreatedAt, &updatedMinat.UpdatedAt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Minat updated, but failed to retrieve updated details")
		return
	}
	if fetchedDesc.Valid {
		updatedMinat.Deskripsi = fetchedDesc.String
	} else {
		updatedMinat.Deskripsi = ""
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedMinat)
}

func DeleteMinat(w http.ResponseWriter, r *http.Request, idStr string) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	query := "UPDATE minat SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL"
	res, err := database.DB.Exec(query, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to soft-delete minat")
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check affected rows after delete")
		return
	}

	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Minat not found or already deleted")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Minat marked as deleted successfully"})
}