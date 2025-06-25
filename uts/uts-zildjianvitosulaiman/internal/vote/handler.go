package vote

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/internal/auth"
	"uts-zildjianvitosulaiman/pkg/utils"
)

// DTO
type VoteRequest struct {
	Type int `json:"type"` // 1 for upvote, -1 for downvote
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

func (h *Handler) Vote(w http.ResponseWriter, r *http.Request) {
	claims, err := getClaims(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	answerID, _ := strconv.Atoi(r.PathValue("id"))

	var req VoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var voteType domain.VoteType
	if req.Type == 1 {
		voteType = domain.VoteTypeUp
	} else if req.Type == -1 {
		voteType = domain.VoteTypeDown
	} else {
		http.Error(w, "Invalid vote type. Use 1 for upvote or -1 for downvote.", http.StatusBadRequest)
		return
	}

	if err := h.service.Vote(claims.UserID, answerID, voteType); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Vote cast successfully."})
}
