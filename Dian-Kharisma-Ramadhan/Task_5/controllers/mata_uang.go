package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"Task_5/database"
	"Task_5/models"
)

func GetAllNilaiMataUang(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT id, nama_mata_uang, singkatan, exchange_rate FROM nilai_mata_uang")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var mataUang []models.NilaiMataUang
	for rows.Next() {
		var m models.NilaiMataUang
		err := rows.Scan(&m.ID, &m.NamaMataUang, &m.Singkatan, &m.ExchangeRate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		mataUang = append(mataUang, m)
	}
	json.NewEncoder(w).Encode(mataUang)
}

func GetNilaiMataUangByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var m models.NilaiMataUang
	err := database.DB.QueryRow("SELECT id, nama_mata_uang, singkatan, exchange_rate FROM nilai_mata_uang WHERE id = ?", id).
		Scan(&m.ID, &m.NamaMataUang, &m.Singkatan, &m.ExchangeRate)
	if err == sql.ErrNoRows {
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(m)
}

func CreateNilaiMataUang(w http.ResponseWriter, r *http.Request) {
	var m models.NilaiMataUang
	json.NewDecoder(r.Body).Decode(&m)

	_, err := database.DB.Exec("INSERT INTO nilai_mata_uang(nama_mata_uang, singkatan, exchange_rate) VALUES (?, ?, ?)", m.NamaMataUang, m.Singkatan, m.ExchangeRate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Nilai mata uang created"})
}

func UpdateNilaiMataUang(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var m models.NilaiMataUang
	json.NewDecoder(r.Body).Decode(&m)

	_, err := database.DB.Exec("UPDATE nilai_mata_uang SET nama_mata_uang = ?, singkatan = ?, exchange_rate = ? WHERE id = ?", m.NamaMataUang, m.Singkatan, m.ExchangeRate, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Nilai mata uang updated"})
}

func DeleteNilaiMataUang(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_, err := database.DB.Exec("DELETE FROM nilai_mata_uang WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Nilai mata uang deleted"})
}
