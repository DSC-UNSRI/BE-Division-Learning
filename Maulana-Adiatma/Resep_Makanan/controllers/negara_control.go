package controllers

import (
	"encoding/json"
	"fmt"
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

	var reseps []models.Country
	for rows.Next() {
		var r models.Country
		if err := rows.Scan(&r.ID, &r.NamaNegara); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		reseps = append(reseps, r)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reseps)
}

func GetNegaraByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParam)

	var rcp models.Country
	err := database.DB.QueryRow("SELECT * FROM data_negara WHERE id = ?", id).
		Scan(&rcp.ID, &rcp.NamaNegara)

	if err != nil {
		http.Error(w, "Country not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rcp)
}

func CreateNegara(w http.ResponseWriter, r *http.Request) {
	var rcp models.Country
	err := json.NewDecoder(r.Body).Decode(&rcp)
	fmt.Println("err:", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("INSERT INTO data_negara(negara_asal) VALUES(?)",
		rcp.NamaNegara)

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

	var rcp models.Country
	err := json.NewDecoder(r.Body).Decode(&rcp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = database.DB.Exec("UPDATE data_negara SET negara_asal=? WHERE id=?",
		rcp.NamaNegara, id)

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
