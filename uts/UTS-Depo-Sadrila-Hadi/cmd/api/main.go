package main

import (
	"encoding/json"
	"fmt"
	"forum-app/internal/controllers"
	"forum-app/internal/database"
	"forum-app/internal/middlewares"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	defer database.DB.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/register", controllers.Register)
	mux.HandleFunc("POST /api/auth/login", controllers.Login)
	mux.HandleFunc("POST /api/auth/forgot-password/request-question", controllers.RequestSecurityQuestion)
	mux.HandleFunc("POST /api/auth/forgot-password/reset", controllers.ResetPassword)
	mux.HandleFunc("GET /api/user/profile", middlewares.AuthMiddleware(getProfile))
	mux.HandleFunc("GET /api/premium/content", middlewares.AuthMiddleware(middlewares.RoleMiddleware("premium")(getPremiumContent)))

	mux.HandleFunc("GET /api/questions", controllers.GetAllQuestions)
	mux.HandleFunc("GET /api/questions/{id}", controllers.GetQuestionByID)
	mux.HandleFunc("POST /api/questions", middlewares.AuthMiddleware(controllers.CreateQuestion))
	mux.HandleFunc("PUT /api/questions/{id}", middlewares.AuthMiddleware(controllers.UpdateQuestion))
	mux.HandleFunc("DELETE /api/questions/{id}", middlewares.AuthMiddleware(controllers.DeleteQuestion))

	mux.HandleFunc("GET /api/questions/{questionID}/answers", controllers.GetAnswersForQuestion)
	mux.HandleFunc("POST /api/questions/{questionID}/answers", middlewares.AuthMiddleware(controllers.CreateAnswer))
	mux.HandleFunc("PUT /api/answers/{answerID}", middlewares.AuthMiddleware(controllers.UpdateAnswer))
	mux.HandleFunc("DELETE /api/answers/{answerID}", middlewares.AuthMiddleware(controllers.DeleteAnswer))

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	claims, ok := middlewares.GetClaimsFromContext(r)
	if !ok {
		http.Error(w, "Could not retrieve user data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":   fmt.Sprintf("Welcome, %s!", claims.Username),
		"user_id":   claims.UserID,
		"user_type": claims.UserType,
	})
}

func getPremiumContent(w http.ResponseWriter, r *http.Request) {
	claims, _ := middlewares.GetClaimsFromContext(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("Hello premium user %s! Here is your exclusive content.", claims.Username),
	})
}