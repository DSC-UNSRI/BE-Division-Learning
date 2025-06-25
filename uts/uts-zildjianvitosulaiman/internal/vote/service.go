package vote

import (
	"database/sql"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/internal/answer"
)

type Service interface {
	Vote(userID, answerID int, voteType domain.VoteType) error
}

type service struct {
	voteRepo   Repository
	answerRepo answer.Repository
}

func NewService(voteRepo Repository, answerRepo answer.Repository) Service {
	return &service{voteRepo: voteRepo, answerRepo: answerRepo}
}

func (s *service) Vote(userID, answerID int, voteType domain.VoteType) error {
	existingVote, err := s.voteRepo.FindByUserAndAnswer(userID, answerID)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existingVote != nil {
		if existingVote.Type == voteType {
			if err := s.voteRepo.Delete(userID, answerID); err != nil {
				return err
			}
		} else {
			if err := s.voteRepo.Upsert(&domain.Vote{UserID: userID, AnswerID: answerID, Type: voteType}); err != nil {
				return err
			}
		}
	} else {
		if err := s.voteRepo.Upsert(&domain.Vote{UserID: userID, AnswerID: answerID, Type: voteType}); err != nil {
			return err
		}
	}

	upvotes, downvotes, err := s.voteRepo.GetVoteCounts(answerID)
	if err != nil {
		return err
	}

	return s.answerRepo.UpdateVoteCounts(answerID, upvotes, downvotes)
}
