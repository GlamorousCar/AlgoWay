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

func NewCheckSystemHandler(
	userUseCase *usecase.UserUseCase,
	checkSystemUseCase *usecase.CheckSystemUseCase,
) *CheckSystemHandler {
	return &CheckSystemHandler{
		userUseCase:        userUseCase,
		checkSystemUseCase: checkSystemUseCase,
	}
}

func (h *CheckSystemHandler) CheckTask(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/check_task" {
		helpers.NotFound(w)
		return
	}
	helpers.InfoLogger.Println("CheckSystemHandler: CheckTask")

	w.Header().Set("Content-Type", "application/json")

	userToken := r.Header.Get("user_token")
	if len(userToken) == 0 {
		helpers.Unauthorized(w)
		return
	}

	// валидируем его, получаем id пользователя
	userId, err := h.userUseCase.ValidateToken(userToken) //вернется user_id или ошибка
	if err != nil {
		helpers.Unauthorized(w)
	}

	request := models.CheckTaskRequest{}
	err = json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	if len(request.SourceCode) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "The code is missing")
	}

	helpers.InfoLogger.Println("CheckSystemHandler: decode Success")
	verdict, err := h.checkSystemUseCase.CheckTask(request.TaskID, request.Lang, request.SourceCode, userId)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	jsonResp, err := json.Marshal(verdict)
	if err != nil {
		helpers.ServerError(w, err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		helpers.ErrorLogger.Println(err)
	}
}
