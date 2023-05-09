package usecase

import (
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
	"github.com/GlamorousCar/AlgoWay/internal/services/checkSystem"
)

type CheckSystemUseCase struct {
	checkSystemRepo repository.CheckSystemRepo
}

func NewCheckSystemUseCase(checkSystemRepo repository.CheckSystemRepo) *CheckSystemUseCase {
	return &CheckSystemUseCase{checkSystemRepo: checkSystemRepo}
}

// CheckTask TODO userId сейчас не используется, он нужен будет в дальнейшем при сохранении вердикта
func (u *CheckSystemUseCase) CheckTask(taskID uint64, lang string, code string, userId int) (*models.Verdict, error) {
	helpers.InfoLogger.Println("CheckSystemUseCase: CheckTask")
	checkModule, err := checkSystem.NewCheckSystem(lang) // Переименовал, потому что переменная называлась так же как и модуль
	if err != nil {
		return nil, err
	}

	helpers.InfoLogger.Println("CheckSystemUseCase: Getting test data")
	testData, err := u.checkSystemRepo.GetTestData(taskID)
	if err != nil {
		return nil, err
	}

	helpers.InfoLogger.Println("CheckSystemUseCase: Writing code to file")
	codeFile, err := checkModule.WriteCodeToFile(code)
	if err != nil {
		return nil, err
	}

	helpers.InfoLogger.Println("CheckTask: Running Tests")

	// TODO Сделать обработку ошибок WA (Wrong Answer), CE (Compilation Error), TL (Time Limit) и мб PE (Presentatiom Error)
	verdict, err := checkModule.RunTests(*testData)
	if err != nil {
		return nil, err
	}

	err = codeFile.Close()

	if err != nil {
		panic(err)
	}

	// TODO сделать сохранение результата в БД
	err = u.checkSystemRepo.SaveAttempt(userId, taskID, verdict)

	if err != nil {
		helpers.ErrorLogger.Fatal(err)
	}
	return verdict, nil
}
