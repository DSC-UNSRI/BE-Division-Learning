package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"log"
	"path/filepath"
	"github.com/gorilla/mux"
	"backend/database"
	"backend/models"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	database.DB.Find(&events)

	// DEBUGGING: Cetak slice events ke terminal
	log.Println("Data Events yang diambil dari database:", events)

	json.NewEncoder(w).Encode(events)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	json.NewDecoder(r.Body).Decode(&event)
	database.DB.Create(&event)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event created successfully"})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	
	var event models.Event
	if err := database.DB.First(&event, id).Error; err != nil {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}
	
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	
	if r.FormValue("location") != "" {
		event.Location = r.FormValue("location")
	}
	if r.FormValue("start") != "" {
		event.Start = r.FormValue("start")
	}

	file, handler, err := r.FormFile("cover")
	if err == nil {
		defer file.Close()
		
		dstPath := filepath.Join("static", "events", handler.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)
		
		event.Cover = "/static/events/" + handler.Filename
	}

	database.DB.Save(&event)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event updated successfully"})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	database.DB.Delete(&models.Event{}, id)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event deleted successfully"})
}