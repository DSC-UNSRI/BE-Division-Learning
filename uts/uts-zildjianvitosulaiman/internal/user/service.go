package user

import (
	"errors"
	"uts-zildjianvitosulaiman/domain"
	"uts-zildjianvitosulaiman/pkg/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(user *domain.User) error
	LoginUser(email, password string) (string, error)
	GetUserProfile(userID int) (*domain.User, error)
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

	user.Tier = domain.TierFree

	return s.repo.Create(user)
}

func (s *userService) LoginUser(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetUserProfile(userID int) (*domain.User, error) {
	if userID <= 0 {
		return nil, errors.New("invalid user ID")
	}
	return s.repo.FindByID(userID)
}
