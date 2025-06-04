package controllers

import (
	"Task_5/database"
	"Task_5/models"
	"Task_5/middleware"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAllTabungan(w http.ResponseWriter, r *http.Request) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, ok := userIDCtx.(int)
	if !ok {
		http.Error(w, "Invalid user ID in context", http.StatusUnauthorized)
		return
	}

	rows, err := database.DB.Query("SELECT id, nama_mata_uang, singkatan, saldo FROM tabungan WHERE user_id = ?", userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tabunganList := []models.Tabungan{}
	for rows.Next() {
		var t models.Tabungan
		err := rows.Scan(&t.ID, &t.NamaMataUang, &t.Singkatan, &t.Saldo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tabunganList = append(tabunganList, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tabungan": tabunganList,
	})
}

func CreateTabungan(w http.ResponseWriter, r *http.Request) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	userID, ok := userIDCtx.(int)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "failed to parse form data", http.StatusBadRequest)
		return
	}

	namaMataUang := r.FormValue("nama_mata_uang")
	singkatan := r.FormValue("singkatan")
	saldoStr := r.FormValue("saldo")

	saldoInt, err := strconv.ParseInt(saldoStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid saldo: must be a number", http.StatusBadRequest)
		return
	}

	t := models.Tabungan{
		NamaMataUang: namaMataUang,
		Singkatan:    singkatan,
		Saldo:        saldoInt,
	}

	res, err := database.DB.Exec(
		"INSERT INTO tabungan(nama_mata_uang, singkatan, saldo, user_id) VALUES (?, ?, ?, ?)",
		t.NamaMataUang, t.Singkatan, t.Saldo, userID,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := res.LastInsertId()
	t.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "Tabungan successfully created",
		"tabungan": t,
	})
}

func UpdateTabungan(w http.ResponseWriter, r *http.Request, idStr string) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, ok := userIDCtx.(int)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	roleCtx := r.Context().Value(middleware.RoleKey)
	role, ok := roleCtx.(string)
	if !ok {
		http.Error(w, "Invalid role", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	if role != "admin" {
		var ownerID int
		err = database.DB.QueryRow("SELECT user_id FROM tabungan WHERE id = ?", id).Scan(&ownerID)
		if err != nil || ownerID != userID {
			http.Error(w, "Unauthorized: You can only update your own tabungan", http.StatusUnauthorized)
			return
		}
	}

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	namaMataUang := r.FormValue("nama_mata_uang")
	singkatan := r.FormValue("singkatan")
	saldoStr := r.FormValue("saldo")

	saldoInt, err := strconv.ParseInt(saldoStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid saldo value", http.StatusBadRequest)
		return
	}

	t := models.Tabungan{
		NamaMataUang: namaMataUang,
		Singkatan:    singkatan,
		Saldo:        saldoInt,
	}

	_, err = database.DB.Exec(
		"UPDATE tabungan SET nama_mata_uang = ?, singkatan = ?, saldo = ? WHERE id = ?",
		t.NamaMataUang, t.Singkatan, t.Saldo, id,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Tabungan updated successfully",
	})
}

func DeleteTabungan(w http.ResponseWriter, r *http.Request, idStr string) {
	userIDCtx := r.Context().Value(middleware.UserIDKey)
	if userIDCtx == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID, ok := userIDCtx.(int)
	if !ok {
		http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		return
	}

	roleCtx := r.Context().Value(middleware.RoleKey)
	role, ok := roleCtx.(string)
	if !ok {
		http.Error(w, "Invalid role", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	var exists bool
	err = database.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tabungan WHERE id = ?)", id).Scan(&exists)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if !exists {
		http.Error(w, "Tabungan not found", http.StatusNotFound)
		return
	}

	if role != "admin" {
		var ownerID int
		err = database.DB.QueryRow("SELECT user_id FROM tabungan WHERE id = ?", id).Scan(&ownerID)
		if err != nil || ownerID != userID {
			http.Error(w, "Unauthorized: You can only delete your own tabungan", http.StatusUnauthorized)
			return
		}
	}

	_, err = database.DB.Exec("DELETE FROM tabungan WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete tabungan", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Tabungan deleted successfully",
		"id":      idStr,
	})
}