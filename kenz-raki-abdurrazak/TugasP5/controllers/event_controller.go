package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/models"
	"github.com/gorilla/mux"
)

func GetAllEvents(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, title, description, speaker_id FROM events")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		rows.Scan(&e.ID, &e.Title, &e.Description, &e.SpeakerID)
		events = append(events, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func GetEventByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	row := config.DB.QueryRow("SELECT id, title, description, speaker_id FROM events WHERE id = ?", id)

	var e models.Event
	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.SpeakerID)
	if err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var e models.Event
	json.NewDecoder(r.Body).Decode(&e)

	stmt, err := config.DB.Prepare("INSERT INTO events(title, description, speaker_id) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(e.Title, e.Description, e.SpeakerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, _ := res.LastInsertId()
	e.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var e models.Event
	json.NewDecoder(r.Body).Decode(&e)

	stmt, err := config.DB.Prepare("UPDATE events SET title=?, description=?, speaker_id=? WHERE id=?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = stmt.Exec(e.Title, e.Description, e.SpeakerID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event updated"})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	stmt, err := config.DB.Prepare("DELETE FROM events WHERE id=?")
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
	json.NewEncoder(w).Encode(map[string]string{"message": "Event deleted"})
}
