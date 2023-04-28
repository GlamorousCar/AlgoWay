package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
)

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (u *TaskUseCase) GetTasks(id int) (*[]models.Task, error) {
	return u.repo.GetTasks(id)
}
