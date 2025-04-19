package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"percobaan3/database"

	"github.com/gorilla/mux"
)

// Struktur untuk hasil join
type ResepJoin struct {
	ID             int    `json:"id"`
	NamaResep      string `json:"nama_resep"`
	DeskripsiResep string `json:"deskripsi_resep"`
	BahanUtama     string `json:"bahan_utama"`
	WaktuMasak     string `json:"waktu_masak"`
	NamaNegara     string `json:"nama_negara"`
}

// GET /resep-negara
func GetResepJoinNegara(w http.ResponseWriter, r *http.Request) {
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

	var results []ResepJoin
	for rows.Next() {
		var r ResepJoin
		if err := rows.Scan(&r.ID, &r.NamaResep, &r.DeskripsiResep, &r.BahanUtama, &r.WaktuMasak, &r.NamaNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

// GET /resep-negara/{id}
func GetResepByNegaraID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	negaraID, _ := strconv.Atoi(idParam)

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

	var results []ResepJoin
	for rows.Next() {
		var r ResepJoin
		if err := rows.Scan(&r.ID, &r.NamaResep, &r.DeskripsiResep, &r.BahanUtama, &r.WaktuMasak, &r.NamaNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		results = append(results, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
