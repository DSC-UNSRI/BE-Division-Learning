package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/models"
	"github.com/gorilla/mux"
)

func GetAllSpeakers(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, expertise, auth_key FROM speakers")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var speakers []models.Speaker
	for rows.Next() {
		var s models.Speaker
		rows.Scan(&s.ID, &s.Name, &s.Expertise, &s.AuthKey)
		speakers = append(speakers, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(speakers)
}

func GetSpeakerByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := config.DB.QueryRow("SELECT id, name, expertise, auth_key FROM speakers WHERE id = ?", id)

	var s models.Speaker
	err := row.Scan(&s.ID, &s.Name, &s.Expertise, &s.AuthKey)
	if err != nil {
		http.Error(w, "Speaker not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func CreateSpeaker(w http.ResponseWriter, r *http.Request) {
	var s models.Speaker
	json.NewDecoder(r.Body).Decode(&s)

	stmt, err := config.DB.Prepare("INSERT INTO speakers(name, expertise, auth_key) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(s.Name, s.Expertise, s.AuthKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := res.LastInsertId()
	s.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

func UpdateSpeaker(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var s models.Speaker
	json.NewDecoder(r.Body).Decode(&s)

	stmt, err := config.DB.Prepare("UPDATE speakers SET name=?, expertise=?, auth_key=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(s.Name, s.Expertise, s.AuthKey, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Speaker updated"})
}

func DeleteSpeaker(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stmt, err := config.DB.Prepare("DELETE FROM speakers WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Speaker deleted"})
}
