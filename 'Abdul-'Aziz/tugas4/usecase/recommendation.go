package usecase

import (
	"errors"

	"tugas/tugas4/models"
	"tugas/tugas4/repository"
)

type RecommendationUseCase struct {
    repo repository.InMemoryUserRepo
}

func NewRecommendationUseCase(repo repository.InMemoryUserRepo) *RecommendationUseCase {

    return &RecommendationUseCase{repo: repo}
}

func (uc *RecommendationUseCase) AddRecommendation(category, content string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    user.Recommendations = append(user.Recommendations, models.Recommendation{
        Category: category,
        Content:  content,
    })
    return uc.repo.SaveUser(user)
}

func (uc *RecommendationUseCase) UpdateRecommendation(index int, category, content string) error {
    user, err := uc.repo.GetUser()
    if err != nil {
        return err
    }
    
    if index < 0 || index >= len(user.Recommendations) {
        return errors.New("index rekomendasi tidak valid")
    }
    
    user.Recommendations[index] = models.Recommendation{
        Category: category,
        Content:  content,
    }
    return uc.repo.SaveUser(user)
}