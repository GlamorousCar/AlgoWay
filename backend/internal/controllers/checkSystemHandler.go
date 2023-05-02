package controllers

import (
	"encoding/json"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type CheckSystemHandler struct {
	userUseCase *usecase.UserUseCase
	checkSystem *usecase.CheckSystemUseCase
}

func NewCheckSystemHandler(checkSystem *usecase.CheckSystemUseCase) *CheckSystemHandler {
	return &CheckSystemHandler{checkSystem: checkSystem}
}

func (h *CheckSystemHandler) CheckTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем токен из запроса и обрабатываем его
	var userToken string
	userToken = r.Header.Get("user_token")

	if len(userToken) == 0 {
		helpers.Unauthorized(w)
		return
	}
	// валидируем его, получаем id пользователя
	userId, err := h.userUseCase.ValidateToken(userToken) //вернется user_id или ошибка
	if err != nil {
		helpers.Unauthorized(w)
	}

	// получаем из запроса номер задачки и код исх.кода и проверяем их
	rawData := models.CheckSystemIncomingData{}
	err = json.NewDecoder(r.Body).Decode(&rawData)

	rawData.TaskId, err = h.checkSystem.CheckTaskIdAndLang(rawData.TaskId, rawData.CodeLang) //вернется номер задачи либо ошибка_not_found
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Получаем код задачки и отправляем все на тестирование
	if len(rawData.SourceCode) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "the code is missing")
	}

	// тестирование возвращает вердикт проверки
	verdict, err := h.checkSystem.TestUserCode(rawData.SourceCode, rawData.CodeLang, rawData.TaskId, userId)

	if err != nil {
		helpers.ServerError(w, errors.New("Проблема с тестирующей системой"))
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
