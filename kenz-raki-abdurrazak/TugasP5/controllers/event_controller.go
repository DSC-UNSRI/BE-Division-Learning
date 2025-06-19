package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/config"
	"github.com/artichys/BE-Division-Learning/kenz-raki-abdurrazak/TugasP5/middleware"
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
		if errScan := rows.Scan(&e.ID, &e.Title, &e.Description, &e.SpeakerID); errScan != nil {
			http.Error(w, "Failed to scan event data: "+errScan.Error(), http.StatusInternalServerError)
			return
		}
		events = append(events, e)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating event rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

func GetEventByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	row := config.DB.QueryRow("SELECT id, title, description, speaker_id FROM events WHERE id = ?", id)

	var e models.Event
	err := row.Scan(&e.ID, &e.Title, &e.Description, &e.SpeakerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Event not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to scan event: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e)
}

func GetMyEvents(w http.ResponseWriter, r *http.Request) {
	speakerIDValue := r.Context().Value(middleware.SpeakerIDKey)
	if speakerIDValue == nil {
		http.Error(w, "Speaker ID not found in context", http.StatusInternalServerError)
		return
	}
	speakerID, ok := speakerIDValue.(int)
	if !ok {
		http.Error(w, "Speaker ID in context is not of expected type", http.StatusInternalServerError)
		return
	}
	var events []models.Event
	rows, err := config.DB.Query("SELECT id, title, description, speaker_id FROM events WHERE speaker_id = ?", speakerID)
	if err != nil {
		http.Error(w, "Failed to retrieve events: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var event models.Event
		if errScan := rows.Scan(&event.ID, &event.Title, &event.Description, &event.SpeakerID); errScan != nil {
			http.Error(w, "Failed to scan event: "+errScan.Error(), http.StatusInternalServerError)
			return
		}
		events = append(events, event)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Error during rows iteration: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var e models.Event
	if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	speakerIDValue := r.Context().Value(middleware.SpeakerIDKey)
	if speakerIDValue == nil {
		http.Error(w, "Authentication error: Speaker ID not found in context", http.StatusUnauthorized)
		return
	}
	authenticatedSpeakerID, ok := speakerIDValue.(int)
	if !ok {
		http.Error(w, "Internal error: Speaker ID in context is not of expected type", http.StatusInternalServerError)
		return
	}

	e.SpeakerID = authenticatedSpeakerID

	if e.Title == "" {
		http.Error(w, "Event title is required", http.StatusBadRequest)
		return
	}

	stmt, err := config.DB.Prepare("INSERT INTO events(title, description, speaker_id) VALUES(?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(e.Title, e.Description, e.SpeakerID)
	if err != nil {
		http.Error(w, "Failed to create event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	e.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventIDStr := vars["id"]
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, "Invalid event ID format", http.StatusBadRequest)
		return
	}

	var eventDataFromRequest models.Event
	if err := json.NewDecoder(r.Body).Decode(&eventDataFromRequest); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	speakerIDValue := r.Context().Value(middleware.SpeakerIDKey)
	if speakerIDValue == nil {
		http.Error(w, "Authentication error: Speaker ID not found in context", http.StatusUnauthorized)
		return
	}
	authenticatedSpeakerID, ok := speakerIDValue.(int)
	if !ok {
		http.Error(w, "Internal error: Speaker ID in context is not of expected type", http.StatusInternalServerError)
		return
	}

	var currentEventOwnerID int
	err = config.DB.QueryRow("SELECT speaker_id FROM events WHERE id = ?", eventID).Scan(&currentEventOwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Event not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to verify event ownership: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if currentEventOwnerID != authenticatedSpeakerID {
		http.Error(w, "Forbidden: You do not own this event", http.StatusForbidden)
		return
	}

	stmt, err := config.DB.Prepare("UPDATE events SET title=?, description=? WHERE id=? AND speaker_id=?")
	if err != nil {
		http.Error(w, "Failed to prepare statement: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(eventDataFromRequest.Title, eventDataFromRequest.Description, eventID, authenticatedSpeakerID)
	if err != nil {
		http.Error(w, "Failed to update event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Event not found with the specified owner or no changes made", http.StatusNotFound)
		return
	}

	eventDataFromRequest.ID = eventID
	eventDataFromRequest.SpeakerID = authenticatedSpeakerID

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(eventDataFromRequest)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventIDStr := vars["id"]
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		http.Error(w, "Invalid event ID format", http.StatusBadRequest)
		return
	}

	speakerIDValue := r.Context().Value(middleware.SpeakerIDKey)
	if speakerIDValue == nil {
		http.Error(w, "Authentication error: Speaker ID not found in context", http.StatusUnauthorized)
		return
	}
	authenticatedSpeakerID, ok := speakerIDValue.(int)
	if !ok {
		http.Error(w, "Internal error: Speaker ID in context is not of expected type", http.StatusInternalServerError)
		return
	}

	var currentEventOwnerID int
	err = config.DB.QueryRow("SELECT speaker_id FROM events WHERE id = ?", eventID).Scan(&currentEventOwnerID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Event not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to verify event ownership: "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if currentEventOwnerID != authenticatedSpeakerID {
		http.Error(w, "Forbidden: You do not own this event", http.StatusForbidden)
		return
	}

	stmt, err := config.DB.Prepare("DELETE FROM events WHERE id=? AND speaker_id=?")
	if err != nil {
		http.Error(w, "Failed to prepare statement: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(eventID, authenticatedSpeakerID)
	if err != nil {
		http.Error(w, "Failed to delete event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Failed to check rows affected: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Event not found with the specified owner or already deleted", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event deleted successfully"})
}

func GetEventsBySpeakerID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	speakerID := vars["id"]

	rows, err := config.DB.Query("SELECT id, title, description, speaker_id FROM events WHERE speaker_id = ?", speakerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var e models.Event
		if errScan := rows.Scan(&e.ID, &e.Title, &e.Description, &e.SpeakerID); errScan != nil {
			http.Error(w, "Failed to scan event data: "+errScan.Error(), http.StatusInternalServerError)
			return
		}
		events = append(events, e)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, "Error iterating event rows: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(events)
}

type FullEventRequest struct {
	Speaker models.Speaker `json:"speaker"`
	Event   models.Event   `json:"event"`
}

func CreateFullEvent(w http.ResponseWriter, r *http.Request) {
	speakerIDValue := r.Context().Value(middleware.SpeakerIDKey)
	if speakerIDValue == nil {
		http.Error(w, "Authentication error: Speaker ID not found in context", http.StatusUnauthorized)
		return
	}
	authenticatedSpeakerID, ok := speakerIDValue.(int)
	if !ok {
		http.Error(w, "Internal error: Speaker ID in context is not of expected type", http.StatusInternalServerError)
		return
	}

	var req FullEventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	if req.Event.Title == "" {
		http.Error(w, "Event title is required", http.StatusBadRequest)
		return
	}
	
	stmt, err := config.DB.Prepare("INSERT INTO events(title, description, speaker_id) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	res, err := stmt.Exec(req.Event.Title, req.Event.Description, authenticatedSpeakerID)
	if err != nil {
		http.Error(w, "Failed to create event: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newEventID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to get last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Event created successfully by authenticated speaker",
		"event_id": newEventID,
		"speaker_id": authenticatedSpeakerID,
		"title": req.Event.Title,
		"description": req.Event.Description,
	})
}