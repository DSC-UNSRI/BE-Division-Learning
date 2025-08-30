package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"nobar-backend/database"
	"nobar-backend/models"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)


func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	database.DB.Find(&events)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string][]models.Event{"events": events})
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	location := r.FormValue("location")
	if location == "" {
		http.Error(w, `{"message":"Location is required"}`, http.StatusBadRequest)
		return
	}
	event := models.Event{Location: location}
	database.DB.Create(&event)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Event created successfully"})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var event models.Event
	if result := database.DB.First(&event, id); result.Error != nil {
		http.Error(w, `{"message":"Event not found"}`, http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(10 << 20)

	location := r.FormValue("location")
	if location != "" {
		event.Location = location
	}

	startStr := r.FormValue("start")
	if startStr != "" {
		layout := "2006-01-02"
		parsedTime, err := time.Parse(layout, startStr)
		if err == nil {
			event.Start = &parsedTime
		}
	}
	
	file, handler, err := r.FormFile("cover")
	if err == nil {
		defer file.Close()
		uploadDir := "./assets"
		os.MkdirAll(uploadDir, os.ModePerm)
		ext := filepath.Ext(handler.Filename)
		newFileName := "event-" + strconv.Itoa(id) + "-" + strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		filePath := filepath.Join(uploadDir, newFileName)
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, `{"message":"Unable to create file"}`, http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, `{"message":"Unable to save file"}`, http.StatusInternalServerError)
			return
		}
		event.Cover = "/api/assets/" + newFileName
	}

	database.DB.Save(&event)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Event updated successfully"})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	database.DB.Delete(&models.Event{}, id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Event deleted successfully"})
}