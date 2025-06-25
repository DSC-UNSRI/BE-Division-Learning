package routes

import (
	"net/http"

	"utsquora/controllers"
	"utsquora/middlewares"
)

func SetupRoutes() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to PT Backend Abadi!"))
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		controllers.Register(w, r)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		controllers.Login(w, r)
	})

	http.HandleFunc("/question", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			controllers.GetAllQuestions(w, r)
		case "POST":
			middleware.AuthMiddleware(middleware.PremiumOnly(controllers.CreateQuestion))(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/question/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			middleware.AuthMiddleware(middleware.PremiumOnly(controllers.UpdateQuestion))(w, r)
		case "DELETE":
			middleware.AuthMiddleware(middleware.PremiumOnly(controllers.DeleteQuestion))(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/answer", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			middleware.AuthMiddleware(controllers.CreateAnswer)(w, r)
		case "GET":
			middleware.AuthMiddleware(controllers.GetAnswersByQuestionID)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/requestreset", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			controllers.RequestResetPassword(w, r) // TANPA middleware
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/resetpassword", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			controllers.ResetPassword(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/upgrade", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			middleware.AuthMiddleware(controllers.UpgradeToPremium)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
