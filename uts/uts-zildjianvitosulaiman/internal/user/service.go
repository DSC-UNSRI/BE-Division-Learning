package user

import (
	"errors"
	"uts-zildjianvitosulaiman/domain" // Sesuaikan nama modul

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *domain.User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterUser(user *domain.User) error {
	if user.Email == "" || user.Password == "" || user.Name == "" || user.SecurityQuestion == "" || user.SecurityAnswer == "" {
		return errors.New("all fields are required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	hashedAnswer, err := bcrypt.GenerateFromPassword([]byte(user.SecurityAnswer), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.SecurityAnswer = string(hashedAnswer)

	// 4. Set tier default
	user.Tier = domain.TierFree

	return s.repo.Create(user)
}
