package controllers

import (
	"UTS-Ahmad-Fadhil-Rizqi/database"
	"UTS-Ahmad-Fadhil-Rizqi/models"
	"UTS-Ahmad-Fadhil-Rizqi/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

func GetMyProfile(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(utils.UserIDKey).(int64)
	if !ok {
		http.Error(w, "Gagal mengambil ID pengguna dari context", http.StatusInternalServerError)
		return
	}
	
	var user models.User
	
	err := database.DB.QueryRow("SELECT id, username, tier, created_at FROM users WHERE id = ? AND deleted_at IS NULL", userID).
		Scan(&user.ID, &user.Username, &user.Tier, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Pengguna tidak ditemukan", http.StatusNotFound)
		} else {
			http.Error(w, "Terjadi kesalahan internal", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}