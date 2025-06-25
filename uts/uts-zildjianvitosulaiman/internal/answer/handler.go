package answer

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/pkg/utils"
)

type CreateAnswerRequest struct {
	Body string `json:"body"`
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

func (h *Handler) CreateAnswer(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	questionID, _ := strconv.Atoi(r.PathValue("questionId"))

	var req CreateAnswerRequest
	json.NewDecoder(r.Body).Decode(&req)

	a := &domain.Answer{
		Body:       req.Body,
		UserID:     claims.UserID,
		QuestionID: questionID,
	}

	if err := h.service.CreateAnswer(a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

func (h *Handler) GetAnswersForQuestion(w http.ResponseWriter, r *http.Request) {
	questionID, _ := strconv.Atoi(r.PathValue("questionId"))
	answers, err := h.service.GetAnswersForQuestion(questionID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(answers)
}

func (h *Handler) UpdateAnswer(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answerID, _ := strconv.Atoi(r.PathValue("id"))

	var req CreateAnswerRequest
	json.NewDecoder(r.Body).Decode(&req)

	a := &domain.Answer{Body: req.Body}

	if err := h.service.UpdateAnswer(claims.UserID, answerID, a); err != nil {
		if err.Error() == "you are not authorized to update this answer" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteAnswer(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answerID, _ := strconv.Atoi(r.PathValue("id"))

	if err := h.service.DeleteAnswer(claims.UserID, answerID); err != nil {
		if err.Error() == "you are not authorized to delete this answer" {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
