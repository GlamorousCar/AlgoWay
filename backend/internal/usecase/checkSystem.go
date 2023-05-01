package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/repository"
)

type CheckSystemUseCase struct {
	repo repository.CheckSystemRepository
}

func NewCheckSystemUseCase(repo repository.CheckSystemRepository) *CheckSystemUseCase {
	return &CheckSystemUseCase{repo: repo}
}

func (u CheckSystemUseCase) Pass(id int) error {
	return u.repo.Pass(id)
}
