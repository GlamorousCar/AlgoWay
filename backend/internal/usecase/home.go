package usecase

import "github.com/GlamorousCar/AlgoWay/internal/repository"

type HomeUseCase struct {
	repo repository.HomeRepository
}

func NewHomeUseCase(repo repository.HomeRepository) *HomeUseCase {
	return &HomeUseCase{repo: repo}
}

func (u HomeUseCase) GetVersion() (string, error) {
	return u.repo.GetVersion()
}
