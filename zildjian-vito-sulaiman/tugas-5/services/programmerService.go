package services

import (
	"errors"
	"tugas-5/models"
	"tugas-5/repositories"
)

type ProgrammerService struct {
	repo *repositories.ProgrammerRepository
}

func NewProgrammerService(repo *repositories.ProgrammerRepository) *ProgrammerService {
	return &ProgrammerService{repo: repo}
}

func (s *ProgrammerService) CreateProgrammer(p *models.Programmer) error {
	if p.Name == "" || p.Email == "" || p.Language == "" || p.YearsOfExperience < 0 {
		return errors.New("all fields must be filled and years of experience must be valid")
	}
	return s.repo.Create(p)
}

func (s *ProgrammerService) GetProgrammer(id int) (*models.Programmer, error) {
	if id <= 0 {
		return nil, errors.New("invalid programmer ID")
	}
	return s.repo.FindByID(id)
}

func (s *ProgrammerService) GetAllProgrammers() ([]*models.Programmer, error) {
	return s.repo.FindAll()
}

func (s *ProgrammerService) UpdateProgrammer(p *models.Programmer) error {
	if p.ID <= 0 || p.Name == "" || p.Email == "" || p.Language == "" {
		return errors.New("invalid programmer or missing fields")
	}
	return s.repo.Update(p)
}

func (s *ProgrammerService) DeleteProgrammer(id int) error {
	if id <= 0 {
		return errors.New("invalid programmer ID")
	}
	return s.repo.Delete(id)
}
