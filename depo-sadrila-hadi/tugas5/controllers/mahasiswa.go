package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"pmm/database"
	"pmm/models"
	"pmm/utils"
)

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request) {
	_, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	rows, err := database.DB.Query("SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE deleted_at IS NULL ORDER BY id ASC")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database query error fetching mahasiswa list")
		return
	}
	defer rows.Close()

	mahasiswas := []models.Mahasiswa{}
	for rows.Next() {
		var mhs models.Mahasiswa
		if err := rows.Scan(&mhs.ID, &mhs.Nama, &mhs.CreatedAt, &mhs.UpdatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan mahasiswa row")
			return
		}
		mahasiswas = append(mahasiswas, mhs)
	}
	if err = rows.Err(); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Row iteration error after fetching mahasiswa list")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"mahasiswa": mahasiswas})
}

func GetMahasiswaByID(w http.ResponseWriter, r *http.Request, idStr string) {
	authenticatedUser, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	if authenticatedUser.ID != id {
		utils.RespondWithError(w, http.StatusForbidden, "Forbidden: You can only view your own profile.")
		return
	}

	var mhs models.Mahasiswa
	query := "SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE id = ? AND deleted_at IS NULL"
	err = database.DB.QueryRow(query, id).Scan(&mhs.ID, &mhs.Nama, &mhs.CreatedAt, &mhs.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Database query error fetching mahasiswa by ID")
		}
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, mhs)
}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama     string `json:"nama"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}
	defer r.Body.Close()

	namaTrimmed := strings.TrimSpace(input.Nama)
	passwordTrimmed := strings.TrimSpace(input.Password)

	if namaTrimmed == "" || passwordTrimmed == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Nama and Password are required and cannot be empty")
		return
	}
	if len(passwordTrimmed) < 6 {
		utils.RespondWithError(w, http.StatusBadRequest, "Password must be at least 6 characters long")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordTrimmed), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	query := "INSERT INTO mahasiswa (nama, password) VALUES (?, ?)"
	res, err := database.DB.Exec(query, namaTrimmed, string(hashedPassword))
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			utils.RespondWithError(w, http.StatusConflict, "Nama already exists")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create mahasiswa")
		}
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to get last insert ID")
		return
	}

	var createdMhs models.Mahasiswa
	fetchQuery := "SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE id = ?"
	err = database.DB.QueryRow(fetchQuery, id).Scan(&createdMhs.ID, &createdMhs.Nama, &createdMhs.CreatedAt, &createdMhs.UpdatedAt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Mahasiswa created, but failed to retrieve full details")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, createdMhs)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, idStr string) {
	authenticatedUser, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	if authenticatedUser.ID != id {
		utils.RespondWithError(w, http.StatusForbidden, "Forbidden: You can only update your own profile.")
		return
	}

	var input struct {
		Nama     *string `json:"nama"`
		Password *string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}
	defer r.Body.Close()

	if input.Nama == nil && input.Password == nil {
		utils.RespondWithError(w, http.StatusBadRequest, "No fields provided to update (expected 'nama' and/or 'password')")
		return
	}

	var setClauses []string
	var args []interface{}

	if input.Nama != nil {
		namaTrimmed := strings.TrimSpace(*input.Nama)
		if namaTrimmed == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Nama cannot be updated to empty")
			return
		}
		setClauses = append(setClauses, "nama = ?")
		args = append(args, namaTrimmed)
	}

	if input.Password != nil {
		passwordTrimmed := strings.TrimSpace(*input.Password)
		if passwordTrimmed == "" {
			utils.RespondWithError(w, http.StatusBadRequest, "Password cannot be updated to empty")
			return
		}
		if len(passwordTrimmed) < 6 {
			utils.RespondWithError(w, http.StatusBadRequest, "New password must be at least 6 characters long")
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordTrimmed), bcrypt.DefaultCost)
		if err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash new password")
			return
		}
		setClauses = append(setClauses, "password = ?")
		args = append(args, string(hashedPassword))
	}

	if len(setClauses) == 0 {
		utils.RespondWithError(w, http.StatusBadRequest, "No valid fields provided for update")
		return
	}

	setClauses = append(setClauses, "updated_at = NOW()")
	args = append(args, id)

	queryUpdate := "UPDATE mahasiswa SET " + strings.Join(setClauses, ", ") + " WHERE id = ? AND deleted_at IS NULL"
	res, err := database.DB.Exec(queryUpdate, args...)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && input.Nama != nil {
			utils.RespondWithError(w, http.StatusConflict, "Update failed: Nama already exists")
		} else {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update mahasiswa")
		}
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check affected rows after update")
		return
	}
	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found or already deleted, update failed")
		return
	}

	var updatedMhs models.Mahasiswa
	fetchQuery := "SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE id = ?"
	err = database.DB.QueryRow(fetchQuery, id).Scan(&updatedMhs.ID, &updatedMhs.Nama, &updatedMhs.CreatedAt, &updatedMhs.UpdatedAt)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Mahasiswa updated, but failed to retrieve updated details")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, updatedMhs)
}

func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, idStr string) {
	authenticatedUser, ok := utils.CheckAuthAndRespond(w, r)
	if !ok {
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID format provided")
		return
	}

	if authenticatedUser.ID != id {
		utils.RespondWithError(w, http.StatusForbidden, "Forbidden: You can only delete your own profile.")
		return
	}

	query := "UPDATE mahasiswa SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL"
	res, err := database.DB.Exec(query, id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to soft-delete mahasiswa")
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check affected rows after delete")
		return
	}
	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Mahasiswa not found or already deleted")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Mahasiswa marked as deleted successfully"})
}

func LoginMahasiswa(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Nama     string `json:"nama"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request body format")
		return
	}
	defer r.Body.Close()

	namaTrimmed := strings.TrimSpace(input.Nama)
	passwordTrimmed := strings.TrimSpace(input.Password)

	if namaTrimmed == "" || passwordTrimmed == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Nama and Password are required")
		return
	}

	var user models.Mahasiswa
	var dbHashedPassword string
	query := "SELECT id, nama, password FROM mahasiswa WHERE nama = ? AND deleted_at IS NULL"
	err := database.DB.QueryRow(query, namaTrimmed).Scan(&user.ID, &user.Nama, &dbHashedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials (user not found)")
		} else {
			fmt.Printf("Database error during login for user %s: %v\n", namaTrimmed, err)
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error during login")
		}
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbHashedPassword), []byte(passwordTrimmed))
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials (password mismatch)")
		return
	}

	responseUser := models.Mahasiswa{
		ID:   user.ID,
		Nama: user.Nama,
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Login successful",
		"user":    responseUser,
	})
}