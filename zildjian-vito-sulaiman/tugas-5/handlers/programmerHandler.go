package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"tugas-5/models"
	"tugas-5/services"
)

type ProgrammerHandler struct {
	service *services.ProgrammerService
}

func NewProgrammerHandler(service *services.ProgrammerService) *ProgrammerHandler {
	return &ProgrammerHandler{service: service}
}

func (h *ProgrammerHandler) GetAllProgrammers(w http.ResponseWriter, r *http.Request) {
	programmers, err := h.service.GetAllProgrammers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(programmers)
}

func (h *ProgrammerHandler) GetProgrammerByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Missing ID in path", http.StatusBadRequest)
		return
	}

	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	programmer, err := h.service.GetProgrammer(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(programmer)
}

func (h *ProgrammerHandler) CreateProgrammer(w http.ResponseWriter, r *http.Request) {
	var programmer models.Programmer
	if err := json.NewDecoder(r.Body).Decode(&programmer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if programmer.Name == "" || programmer.Email == "" || programmer.Language == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateProgrammer(&programmer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(programmer)
}

func (h *ProgrammerHandler) UpdateProgrammer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Missing ID in path", http.StatusBadRequest)
		return
	}

	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var programmer models.Programmer
	if err := json.NewDecoder(r.Body).Decode(&programmer); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	programmer.ID = id
	if err := h.service.UpdateProgrammer(&programmer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(programmer)
}

func (h *ProgrammerHandler) DeleteProgrammer(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Missing ID in path", http.StatusBadRequest)
		return
	}

	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteProgrammer(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
