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
	// 1. Parsing multipart form data
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	var user models.User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Role = r.FormValue("role")
	
	user.ProfilePicture = "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAOEAAADhCAMAAAAJbSJIAAAAb1BMVEX///8CAgJCQkL5+fl/f3/09PT8/Pzw8PCoqKjIyMjExMRMTEwrKyvd3d1mZmaGhoZRUVGenp7k5OTV1dWsrKwlJSXb29s7OzszMzO5ubmXl5diYmJ7e3tycnJXV1cwMDAeHh6Pj4+0tLQREREXFxdRdW5XAAAD6UlEQVR4nO3ca3eiMBAGYCMCFbQIahUv1Nr+/9+4UNf1siCZSJhM+z7f9mzP6byHAMkkdDAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAPi1smizL+JhXOzTKOMupnvzZKKuLZM5d0ld8ha38U4mC4+7sK4clzX5vi/kkbu0TkxnDfkqxZS7vOctHuSrvHAX+KxRS0ClRtwlPiV4bQ2o1Ct3lU/wdAKWEQPuQo21D1HhA/VFM6DYx82bdkClZL40VoSEK+5iTeiPUaHj1P8kJfz0uQsmS0kBlUq5C6by1sSEa2kLjTExoFJj7pKJ9uSEe+6Safx3csJ3Wc8aytv+TNZbv21VWGfBXTRJYpAw4S6aRG/ZdEvUOjF41JtpEktaJgaUWffZSlJC8oymImpW4xldQ0kJg9ggoaj70OhZuuUumkS3B3VNVj+KtsA/kbXMnxsklLXdFjZtNzVbhtxF09AfNaImbaWcnDDnLpnIpw5TaYN0MPggJvzgLpgsIyYUeDqDdhHlXcJybkpKKGpOehYRAkprlv6l36yR1aK5CArNgIXIMVoJh1oBh+JehRehzlp/JThgObVpv4qxrG7+/9rWwrLWvbXyRzPUpbT5dq2w+a2RiL4Fr2T1Q3UkcC7aKFzc9/lnC/HXL4yim3/70WZ7Oio82W6i2+5vFEmLG+a7QxllU/NfddOXTfmzh10uJqUf/WvS7HQa9d7u/OOvkYR3Y5he78rE7VvX2fUOwDp1/UIG6eHuidLW573vHR9Sp6fh45qvDoaP1n7zmlndxN3Fot9wiGbWVPK4YYtq7+jtOK/7bOR8e2X3Yy/I0uYzNxMnO/zHxnq/DUf5NDw9Wr1wmo9aFh0OfmuicxbxaxUXRRGvvjR+1rmzih8aRdM41l/cdB6wfkbExmRDtJ1DW6aUzihF1P6r+0HdotDnyhTO5GiJnpg72gn1zDqFE+8Me2O04kKjw+RwkD4Hdvfph/Jp+GeoW8sJ2S+iyZF1mjfmhCYH2GiY2/6+9YBK8S6HWxaFneBdKtp+zlR4z532EFApzoC2X4YnnL03m1PSC87JaR+3Ie+NSD8na2LJF9DrJaBSfB9i2F04XfAtoUzOq5vgW1/087LgfF3Y6rHd4+u5/fyEP3+UTntKyPeVN+2gsznGfW+Tz33pZnwBLe3I3OPcoemjicG8e7Frr+9pvH8SJOwhIXNj3/6dyL5PanfbwoGe98C3t3tYceGcu6/73YiJwoGA5UrfXmd/5MopvjH9bybpeHfpCF/e/d0Yu/ahQpYnRTzsRlwkuQvb2wAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAVvwBKrAwIMPWeSgAAAAASUVORK5CYII=" 

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