package question

import (
	"errors"
	"uts-zildjianvitosulaiman/domain" // Sesuaikan nama modul
)

type Service interface {
	CreateQuestion(input *domain.Question) error
	GetAllQuestions() ([]*domain.Question, error)
	GetQuestionByID(id int) (*domain.Question, error)
	UpdateQuestion(userID, questionID int, input *domain.Question) error
	DeleteQuestion(userID, questionID int) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateQuestion(q *domain.Question) error {
	if q.Title == "" || q.Body == "" {
		return errors.New("title and body are required")
	}
	return s.repo.Create(q)
}

func (s *service) GetAllQuestions() ([]*domain.Question, error) {
	return s.repo.FindAll()
}

func (s *service) GetQuestionByID(id int) (*domain.Question, error) {
	return s.repo.FindByID(id)
}

func (s *service) UpdateQuestion(userID, questionID int, input *domain.Question) error {
	q, err := s.repo.FindByID(questionID)
	if err != nil {
		return errors.New("question not found")
	}

	if q.UserID != userID {
		return errors.New("you are not authorized to update this question")
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
