package services

import (
	"errors"
	"os"
	"tugas-5/models"
	"tugas-5/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return errors.New("all fields must be filled")
	}
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id int) (*models.User, error) {
	if id <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.FindByID(id)
}

func (s *UserService) GetAllUsers() ([]*models.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) UpdateUser(user *models.User) error {
	if user.ID <= 0 || user.Name == "" || user.Email == "" {
		return errors.New("invalid user or missing fields")
	}
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	if id <= 0 {
		return errors.New("invalid user ID")
	}
	return s.repo.Delete(id)
}

func (s *UserService) FakeAuth(email, password string) error {
	expectedEmail := os.Getenv("EMAIL")
	expectedPassword := os.Getenv("PASSWORD")

	if expectedEmail == "" || expectedPassword == "" {
		return errors.New("EMAIL or PASSWORD is not set in environment variables")
	}

	if email != expectedEmail || password != expectedPassword {
		return errors.New("invalid email or password")
	}

	return nil
}
