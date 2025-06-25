package answer

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/pkg/utils"
)

type CreateAnswerRequest struct {
	Body string `json:"body"`
}

type AnswerResponse struct {
	ID         int    `json:"id"`
	Body       string `json:"body"`
	UserID     int    `json:"user_id"`
	QuestionID int    `json:"question_id"`
	Upvotes    int    `json:"upvotes"`
	Downvotes  int    `json:"downvotes"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func toAnswerResponse(a *domain.Answer) AnswerResponse {
	return AnswerResponse{
		ID:         a.ID,
		Body:       a.Body,
		UserID:     a.UserID,
		QuestionID: a.QuestionID,
		Upvotes:    a.Upvotes,
		Downvotes:  a.Downvotes,
		CreatedAt:  a.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  a.UpdatedAt.Format(time.RFC3339),
	}
}

func toAnswerListResponse(answers []*domain.Answer) []AnswerResponse {
	responses := make([]AnswerResponse, len(answers))
	for i, a := range answers {
		responses[i] = toAnswerResponse(a)
	}
	return responses
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

	response := toAnswerResponse(a)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) GetAnswersForQuestion(w http.ResponseWriter, r *http.Request) {
	questionID, _ := strconv.Atoi(r.PathValue("questionId"))
	answers, err := h.service.GetAnswersForQuestion(questionID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responses := toAnswerListResponse(answers)

	json.NewEncoder(w).Encode(responses)
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
