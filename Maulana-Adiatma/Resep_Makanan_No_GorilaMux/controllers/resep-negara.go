package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"resepku/database"
	"resepku/models"
)

func GetResepResultNegara(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query(`
		SELECT data_resep.id, data_resep.nama_resep, data_resep.description,
		       data_resep.bahan_utama, data_resep.waktu_masak,
		       data_negara.negara_asal
		FROM data_resep
		JOIN data_negara ON data_resep.negara_id = data_negara.id
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var r models.Result
		if err := rows.Scan(&r.ID, &r.NamaResep, &r.DeskripsiResep, &r.BahanUtama, &r.WaktuMasak, &r.NamaNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetResepByNegaraID(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	negaraID, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid negara ID", http.StatusBadRequest)
		return
	}

	rows, err := database.DB.Query(`
		SELECT data_resep.id, data_resep.nama_resep, data_resep.description,
		       data_resep.bahan_utama, data_resep.waktu_masak,
		       data_negara.negara_asal
		FROM data_resep
		JOIN data_negara ON data_resep.negara_id = data_negara.id
		WHERE data_negara.id = ?`, negaraID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.Result
	for rows.Next() {
		var r models.Result
		if err := rows.Scan(&r.ID, &r.NamaResep, &r.DeskripsiResep, &r.BahanUtama, &r.WaktuMasak, &r.NamaNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, r)
	}

	if len(results) == 0 {
		http.Error(w, "No resep found for the given negara ID", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)}