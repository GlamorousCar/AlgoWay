package controllers

import (
	"encoding/json"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type CheckSystemHandler struct {
	userUseCase        *usecase.UserUseCase
	checkSystemUseCase *usecase.CheckSystemUseCase
}

func NewCheckSystemHandler(checkSystemUseCase *usecase.CheckSystemUseCase) *CheckSystemHandler {
	return &CheckSystemHandler{checkSystemUseCase: checkSystemUseCase}
}

func (h *CheckSystemHandler) CheckTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/check_task" {
		helpers.NotFound(w)
		return
	}

	helpers.InfoLogger.Println("CheckSystemHandler: CheckTask")
	request := models.CheckTaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	helpers.InfoLogger.Println("CheckSystemHandler: decode Success")
	err = h.checkSystemUseCase.CheckTask(request.TaskID, request.Lang, request.Code)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
}
