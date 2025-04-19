vpackage controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"percobaan3/database"
	"percobaan3/models"

	"github.com/gorilla/mux"
)

func GetAllNegara(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM data_negara")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var countries []models.Country
	for rows.Next() {
		var c models.Country
		if err := rows.Scan(&c.ID, &c.NamaNegara, &c.KodeNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		countries = append(countries, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func GetNegaraByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var c models.Country
	err := database.DB.QueryRow("SELECT * FROM data_negara WHERE id = ?", id).
		Scan(&c.ID, &c.NamaNegara, &c.KodeNegara)

	if err != nil {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func CreateNegara(w http.ResponseWriter, r *http.Request) {
	var c models.Country
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("INSERT INTO data_negara (negara_asal, kode_negara) VALUES (?, ?)", c.NamaNegara, c.KodeNegara)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Country created"))
}

func UpdateNegara(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var c models.Country
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE data_negara SET negara_asal=?, kode_negara=? WHERE id=?", c.NamaNegara, c.KodeNegara, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Country updated"))
}

func DeleteNegara(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	_, err := database.DB.Exec("DELETE FROM data_negara WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Country deleted"))
}
