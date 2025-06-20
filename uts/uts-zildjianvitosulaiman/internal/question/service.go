package question

import (
	"errors"
	"time"
	"uts-zildjianvitosulaiman/domain" // Sesuaikan nama modul
)

type Service interface {
	CreateQuestion(input *domain.Question, userTier domain.UserTier) error
	GetAllQuestions() ([]*domain.Question, error)
	GetQuestionByID(id int) (*domain.Question, error)
	UpdateQuestion(userID int, questionID int, userTier domain.UserTier, input *domain.Question) error
	DeleteQuestion(userID, questionID int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateQuestion(q *domain.Question, userTier domain.UserTier) error {
	if q.Title == "" || q.Body == "" {
		return errors.New("title and body are required")
	}

	if userTier == domain.TierFree {
		count, err := s.repo.CountByUserIDAndDate(q.UserID)
		if err != nil {
			return err
		}

		if count >= 5 {
			return errors.New("free users can only create 5 questions per day. Upgrade to premium for unlimited access!")
		}
	}

	return s.repo.Create(q)
}

func (s *service) GetAllQuestions() ([]*domain.Question, error) {
	return s.repo.FindAll()
}

func (s *service) GetQuestionByID(id int) (*domain.Question, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateQuestion(userID, questionID int, userTier domain.UserTier, input *domain.Question) error {
	q, err := s.repo.FindByID(questionID)
	if err != nil {
		return errors.New("question not found")
	}

	if q.UserID != userID {
		return errors.New("you are not authorized to update this question")
	}

	if userTier == domain.TierFree {
		// Cek apakah sudah lewat 5 menit sejak dibuat
		if time.Since(q.CreatedAt) > 5*time.Minute {
			return errors.New("free users can only edit questions within 5 minutes of posting")
		}
	}

	q.Title = input.Title
	q.Body = input.Body
	return s.repo.Update(q)
}

func (s *service) DeleteQuestion(userID, questionID int) error {
	q, err := s.repo.FindByID(questionID)
	if err != nil {
		return errors.New("question not found")
	}

	if q.UserID != userID {
		return errors.New("you are not authorized to delete this question")
	}

	return s.repo.Delete(questionID)
}
