package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"nobar-backend/auth"
	"nobar-backend/database"
	"nobar-backend/middleware"
	"nobar-backend/models"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetMe(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.UserIDKey).(uint)
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		http.Error(w, `{"message":"User not found"}`, http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, `{"message":"ID is missing in parameters"}`, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"message":"Invalid ID"}`, http.StatusBadRequest)
		return
	}

	userIDFromToken := r.Context().Value(middleware.UserIDKey).(uint)
	userRoleFromToken := r.Context().Value(middleware.UserRoleKey).(string)

	if userIDFromToken != uint(id) && userRoleFromToken != "admin" {
		http.Error(w, `{"message":"Forbidden"}`, http.StatusForbidden)
		return
	}

	var user models.User
	if result := database.DB.First(&user, id); result.Error != nil {
		http.Error(w, `{"message":"User not found"}`, http.StatusNotFound)
		return
	}

	r.ParseMultipartForm(10 << 20)
	
	name := r.FormValue("name")
	if name != "" {
		user.Name = name
	}

	password := r.FormValue("password")
	if password != "" {
		hashedPassword, err := auth.HashPassword(password)
		if err != nil {
			http.Error(w, `{"message":"Failed to hash password"}`, http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}

	file, handler, err := r.FormFile("profile_picture")
	if err == nil {
		defer file.Close()
		
		uploadDir := "./assets"
		os.MkdirAll(uploadDir, os.ModePerm)
		
		ext := filepath.Ext(handler.Filename)
		newFileName := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		filePath := filepath.Join(uploadDir, newFileName)
		
		dst, err := os.Create(filePath)
		if err != nil {
			http.Error(w, `{"message":"Unable to create the file for writing"}`, http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, `{"message":"Unable to write the file"}`, http.StatusInternalServerError)
			return
		}
		user.ProfilePicture = "/api/assets/" + newFileName
	}


	if result := database.DB.Save(&user); result.Error != nil {
		http.Error(w, `{"message":"Failed to update profile"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}