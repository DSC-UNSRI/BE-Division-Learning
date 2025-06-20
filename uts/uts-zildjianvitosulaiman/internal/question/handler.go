package question

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/pkg/utils"
)

type CreateQuestionRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type QuestionResponse struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func getClaims(r *http.Request) (*utils.JWTClaims, error) {
	claims, ok := r.Context().Value(auth.ClaimsContextKey).(*utils.JWTClaims)
	if !ok {
		return nil, errors.New("could not retrieve user claims")
	}
	return claims, nil
}

func (h *Handler) CreateQuestion(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var req CreateQuestionRequest
	json.NewDecoder(r.Body).Decode(&req)

	q := &domain.Question{
		Title:  req.Title,
		Body:   req.Body,
		UserID: claims.UserID,
	}

	if err := h.service.CreateQuestion(q); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *Handler) GetAllQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.service.GetAllQuestions()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(questions)
}

func (h *Handler) GetQuestionByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	question, err := h.service.GetQuestionByID(id)
	if err != nil {
		http.Error(w, "Question not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(question)
}

func (h *Handler) UpdateQuestion(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := strconv.Atoi(r.PathValue("id"))

	var req CreateQuestionRequest
	json.NewDecoder(r.Body).Decode(&req)

	q := &domain.Question{
		Title: req.Title,
		Body:  req.Body,
	}

	if err := h.service.UpdateQuestion(claims.UserID, id, q); err != nil {
		if err.Error() == "you are not authorized to update this question" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteQuestion(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(r.PathValue("id"))

	if err := h.service.DeleteQuestion(claims.UserID, id); err != nil {
		if err.Error() == "you are not authorized to delete this question" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
