package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"backend/database"
	"backend/models"
	"log"
)

var jwtKey = []byte("supersecretkey")

type Claims struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
	User    struct {
		ID   uint   `json:"id"`
		Role string `json:"role"`
	} `json:"user"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	var user models.User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Role = r.FormValue("role")
	user.ProfilePicture = "/static/default_profile.jpg" 

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Could not hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if result := database.DB.Create(&user); result.Error != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	email := r.FormValue("email")
	password := r.FormValue("password")

	var user models.User
	if err := database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Id:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	var loginResponse LoginResponse
	loginResponse.Message = "Login successful"
	loginResponse.Token = tokenString
	loginResponse.User.ID = user.ID
	loginResponse.User.Role = user.Role

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResponse)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Logged out successfully"})
}

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		log.Println("Middleware JWT: Menerima request ke:", r.URL.Path)

		if authHeader == "" {
			log.Println("Middleware JWT: Header Authorization kosong")
			http.Error(w, "Unauthorized: Token not provided", http.StatusUnauthorized)
			return
		}
		
		tokenString := strings.Split(authHeader, " ")[1]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		
		if err != nil || !token.Valid {
			log.Println("Middleware JWT: Token tidak valid")
			http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
			return
		}

		log.Println("Middleware JWT: Token valid, User ID:", claims.Id)
		ctx := context.WithValue(r.Context(), "userID", claims.Id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}