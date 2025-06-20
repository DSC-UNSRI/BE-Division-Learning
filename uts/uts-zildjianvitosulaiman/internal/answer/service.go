package answer

import (
	"errors"
	"uts-zildjianvitosulaiman/domain"
)

type Service interface {
	CreateAnswer(input *domain.Answer) error
	GetAnswersForQuestion(questionID int) ([]*domain.Answer, error)
	UpdateAnswer(userID, answerID int, input *domain.Answer) error
	DeleteAnswer(userID, answerID int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateAnswer(a *domain.Answer) error {
	if a.Body == "" {
		return errors.New("answer body cannot be empty")
	}
	return s.repo.Create(a)
}

func (s *service) GetAnswersForQuestion(questionID int) ([]*domain.Answer, error) {
	return s.repo.FindByQuestionID(questionID)
}

func (s *service) UpdateAnswer(userID, answerID int, input *domain.Answer) error {
	a, err := s.repo.FindByID(answerID)
	if err != nil {
		return errors.New("answer not found")
	}

	if a.UserID != userID {
		return errors.New("you are not authorized to update this answer")
	}

	a.Body = input.Body
	return s.repo.Update(a)
}

func (s *service) DeleteAnswer(userID, answerID int) error {
	a, err := s.repo.FindByID(answerID)
	if err != nil {
		return errors.New("answer not found")
	}

	if a.UserID != userID {
		return errors.New("you are not authorized to delete this answer")
	}

	return s.repo.Delete(answerID)
}
