package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
)

type AlgorithmTheoryUseCase struct {
	repo repository.AlgoTheoryRepository
}

func NewAlgorithmTheoryUseCase(repo repository.AlgoTheoryRepository) *AlgorithmTheoryUseCase {
	return &AlgorithmTheoryUseCase{repo: repo}
}

func (u AlgorithmTheoryUseCase) GetAlgoTheory(id int) (*models.AlgorithmTheory, error) {
	return u.repo.GetAlgoTheory(id)
}
