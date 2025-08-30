package handlers

import (
	"encoding/json"
	"net/http"
	"nobar-backend/auth"
	"nobar-backend/database"
	"nobar-backend/models"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	password := r.FormValue("password")
	user.Role = r.FormValue("role")

	if user.Role == "" || (user.Role != "user" && user.Role != "admin") {
		user.Role = "user"
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		http.Error(w, `{"message":"Failed to hash password"}`, http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword
	user.ProfilePicture = "/api/assets/default.png"

	if result := database.DB.Create(&user); result.Error != nil {
		http.Error(w, `{"message":"Could not create user"}`, http.StatusBadRequest)
		return
	}
	
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User registered successfully",
		"user":    user,
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	email := r.FormValue("email")
	password := r.FormValue("password")

	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, `{"message":"Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		http.Error(w, `{"message":"Invalid credentials"}`, http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateJWT(user)
	if err != nil {
		http.Error(w, `{"message":"Could not generate token"}`, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    tokenString,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login successful",
		"user": map[string]interface{}{
			"id":   user.ID,
			"role": user.Role,
		},
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
}