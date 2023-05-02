package usecase

import (
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
)

type CheckSystemUseCase struct {
	repo repository.CheckSystemRepository
}

func NewCheckSystemUseCase(repo repository.CheckSystemRepository) *CheckSystemUseCase {
	return &CheckSystemUseCase{repo: repo}
}

func (u CheckSystemUseCase) CheckTaskIdAndLang(taskId int, lang string) (int, error) {

	if lang != "py" || lang != "go" {
		return 0, errors.New("the language is not supported")
	}
	taskId, err := u.repo.GetTask(taskId)
	if err != nil {
		return 0, err
	}

	return taskId, nil
}

func (u CheckSystemUseCase) TestUserCode(sourceCode, codeLang string, taskID, userId int) (verdict models.Verdict, err error) {
	return models.Verdict{"OK", "All test passed"}, nil
}
