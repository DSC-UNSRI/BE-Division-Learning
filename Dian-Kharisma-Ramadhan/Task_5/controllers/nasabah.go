package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"Task_5/database"
	"Task_5/models"
)

func GetAllNasabah(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nama FROM nasabah")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	nasabahs := []models.Nasabah{}
	for rows.Next() {
		var n models.Nasabah
		err := rows.Scan(&n.ID, &n.Nama)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		nasabahs = append(nasabahs, n)
	}
	json.NewEncoder(w).Encode(nasabahs)
}

func GetNasabahByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var n models.Nasabah
	err := database.DB.QueryRow("SELECT id, nama FROM nasabah WHERE id = ?", id).Scan(&n.ID, &n.Nama)
	if err == sql.ErrNoRows {
		http.Error(w, "Nasabah not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(n)
}

func CreateNasabah(w http.ResponseWriter, r *http.Request) {
	var n models.Nasabah
	json.NewDecoder(r.Body).Decode(&n)

	_, err := database.DB.Exec("INSERT INTO nasabah(nama, password) VALUES (?, ?)", n.Nama, n.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Nasabah created"})
}

func UpdateNasabah(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var n models.Nasabah
	json.NewDecoder(r.Body).Decode(&n)

	_, err := database.DB.Exec("UPDATE nasabah SET nama = ?, password = ? WHERE id = ?", n.Nama, n.Password, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Nasabah updated"})
}

func DeleteNasabah(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_, err := database.DB.Exec("DELETE FROM nasabah WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Nasabah deleted"})
}

func LoginNasabah(w http.ResponseWriter, r *http.Request) {
	var n models.Nasabah
	json.NewDecoder(r.Body).Decode(&n)

	var result models.Nasabah
	err := database.DB.QueryRow("SELECT id, nama FROM nasabah WHERE nama = ? AND password = ?", n.Nama, n.Password).Scan(&result.ID, &result.Nama)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Nasabah login"})
}