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

func GetEventsBySpeakerID(w http.ResponseWriter, r *http.Request) {
	speakerID := mux.Vars(r)["id"]

	rows, err := config.DB.Query("SELECT id, title, description, speaker_id FROM events WHERE speaker_id = ?", speakerID)
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

type FullEventRequest struct {
	Speaker models.Speaker `json:"speaker"`
	Event   models.Event   `json:"event"`
}

func CreateFullEvent(w http.ResponseWriter, r *http.Request) {
	var req FullEventRequest
	json.NewDecoder(r.Body).Decode(&req)

	tx, err := config.DB.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var speakerID int
	err = tx.QueryRow("SELECT id FROM speakers WHERE name=? AND auth_key=?", req.Speaker.Name, req.Speaker.AuthKey).Scan(&speakerID)

	if err != nil {
		res, err := tx.Exec("INSERT INTO speakers(name, expertise, auth_key) VALUES (?, ?, ?)",
			req.Speaker.Name, req.Speaker.Expertise, req.Speaker.AuthKey)

		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastID, _ := res.LastInsertId()
		speakerID = int(lastID)
	}

	_, err = tx.Exec("INSERT INTO events(title, description, speaker_id) VALUES (?, ?, ?)",
		req.Event.Title, req.Event.Description, speakerID)

	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tx.Commit()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Speaker and Event saved successfully"})
}
