package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
)

type ThemeMenuUseCase struct {
	repo repository.ThemeMenuRepository
}

func NewThemeMenuUseCase(repo repository.ThemeMenuRepository) *ThemeMenuUseCase {
	return &ThemeMenuUseCase{
		repo: repo,
	}
}

func (u *ThemeMenuUseCase) GetMenu() (*[]models.ThemeMenu, error) {
	return u.repo.GetMenu()
}
