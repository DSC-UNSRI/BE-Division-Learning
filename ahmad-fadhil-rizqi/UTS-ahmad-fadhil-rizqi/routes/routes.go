package routes

import (
	"UTS-Ahmad-Fadhil-Rizqi/controllers"
	"UTS-Ahmad-Fadhil-Rizqi/middleware"
	"net/http"
	"strings"
)


func Routes() {
	http.HandleFunc("/register", controllers.RegisterUser)
	http.HandleFunc("/login", controllers.LoginUser)

	http.HandleFunc("/forgot-password/get-question", controllers.GetSecurityQuestionForUser)
	http.HandleFunc("/forgot-password/reset", controllers.ResetPassword)

	
	http.Handle("/profile", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetMyProfile)))
	http.Handle("/security-question", middleware.AuthMiddleware(http.HandlerFunc(controllers.SetSecurityQuestion)))
	
	http.HandleFunc("/answers/", func(w http.ResponseWriter, r *http.Request) {
		if len(strings.Split(strings.Trim(r.URL.Path, "/"), "/")) != 2 {
			http.NotFound(w, r)
			return
		}
		
		middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case http.MethodPut, http.MethodPatch:
				controllers.UpdateAnswer(w, r)
			case http.MethodDelete:
				controllers.DeleteAnswer(w, r)
			default:
				http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
			}
		})).ServeHTTP(w, r)
	})

	
	http.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		pathParts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
		
		if len(pathParts) == 3 && pathParts[2] == "answers" {
			if r.Method == http.MethodPost {
				middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateAnswer)).ServeHTTP(w,r)
			} else {
				http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
			}
			return
		}
		
		if len(pathParts) == 2 {
			switch r.Method {
			case http.MethodGet: 
				controllers.GetQuestionByID(w, r)
			case http.MethodPut, http.MethodPatch: 
				middleware.AuthMiddleware(http.HandlerFunc(controllers.UpdateQuestion)).ServeHTTP(w, r)
			case http.MethodDelete: 
				middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteQuestion)).ServeHTTP(w, r)
			default:
				http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
			}
			return
		}
		
		if len(pathParts) == 1 {
			switch r.Method {
			case http.MethodGet: 
				controllers.GetAllQuestions(w, r)
			case http.MethodPost: 
				middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateQuestion)).ServeHTTP(w, r)
			default:
				http.Error(w, "Metode tidak diizinkan", http.StatusMethodNotAllowed)
			}
			return
		}
		
		http.NotFound(w, r)
	})
	
	http.HandleFunc("/questions", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/questions/", http.StatusPermanentRedirect)
	})
}