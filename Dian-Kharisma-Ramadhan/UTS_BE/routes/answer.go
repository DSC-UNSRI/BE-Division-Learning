package routes

import (
	"UTS_BE/controllers"
	"UTS_BE/middleware"
	"UTS_BE/utils"
	"net/http"
)

func AnswerRoutes() {
	http.HandleFunc("/answers/create", middleware.AuthMiddleware(controllers.CreateAnswer))
	http.HandleFunc("/answers/question/", middleware.WithAuth(getAnswersByQuestionHandler))
	http.HandleFunc("/answers/update/", middleware.WithAuth(updateAnswerHandler))
	http.HandleFunc("/answers/delete/", middleware.AuthMiddleware(deleteAnswerHandler))
}

func getAnswersByQuestionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 || parts[1] != "question" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	qid := parts[2]
	controllers.GetAnswersByQuestion(w, r, qid)
}

func updateAnswerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	id := parts[2]
	controllers.UpdateAnswer(w, r, id)
}

func deleteAnswerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	parts, err := utils.SplitPath(r.URL.Path)
	if err != nil || len(parts) < 3 {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}
	aid := parts[2]
	controllers.DeleteAnswer(w, r, aid)
}
