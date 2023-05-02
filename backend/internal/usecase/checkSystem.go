package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
	"github.com/GlamorousCar/AlgoWay/internal/services/checkSystem"
)

type CheckSystemUseCase struct {
	checkSystemRepo repository.CheckSystemRepo
}

func NewCheckSystemUseCase(checkSystemRepo repository.CheckSystemRepo) *CheckSystemUseCase {
	return &CheckSystemUseCase{checkSystemRepo: checkSystemRepo}
}

func (u *CheckSystemUseCase) CheckTask(taskID uint64, lang string, code string) error {
	helpers.InfoLogger.Println("CheckSystemUseCase: CheckTask")
	checkSystem, err := checkSystem.NewCheckSystem(lang)

	helpers.InfoLogger.Println("CheckSystemUseCase: Getting test data")
	testData, err := u.checkSystemRepo.GetTestData(taskID)
	if err != nil {
		return err
	}

	helpers.InfoLogger.Println("CheckSystemUseCase: Writing code to file")
	err = checkSystem.WriteCodeToFile(code)
	if err != nil {
		return err
	}

	helpers.InfoLogger.Println("Running Tests")
	err = checkSystem.RunTests(*testData)
	return err
}
