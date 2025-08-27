package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"log"
	"strconv"
	"github.com/gorilla/mux"
	"backend/database"
	"backend/models"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	// ID user didapat dari JWT di middleware
	userID, ok := r.Context().Value("userID").(uint) // <-- Perbaiki di sini, gunakan 'uint'
	if !ok {
		log.Println("GetMe: User ID tidak ditemukan atau tidak valid") // Log 5
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}
	log.Println("GetMe: Mencari data user dengan ID:", userID) // Log 6

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		log.Println("GetMe: Gagal menemukan user:", err) // Log 7
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	user.Password = "" // Jangan kirim password ke frontend
	log.Println("GetMe: Berhasil mengambil data user:", user.Email) // Log 8
	json.NewEncoder(w).Encode(user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userIDFromToken, ok := r.Context().Value("userID").(float64)
	if !ok || uint(userIDFromToken) != uint(id) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	
	file, handler, err := r.FormFile("cover")
	if err == nil {
		defer file.Close()
		
		dstPath := filepath.Join("static", handler.Filename)
		dst, err := os.Create(dstPath)
		if err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		io.Copy(dst, file)
		
		if result := database.DB.Model(&models.User{}).Where("id = ?", id).Update("profile_picture", "/static/"+handler.Filename); result.Error != nil {
			http.Error(w, "Failed to update profile picture", http.StatusInternalServerError)
			return
		}
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}