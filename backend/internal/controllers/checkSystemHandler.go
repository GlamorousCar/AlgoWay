package controllers

import (
	"encoding/json"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
	"strconv"
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
	userToken := r.URL.Query().Get("token")

	if len(userToken) == 0 {
		helpers.Unauthorized(w)
		return
	}
	// валидируем его, получаем id пользователя
	userId, err := h.userUseCase.ValidateToken(userToken) //вернется user_id или ошибка
	if err != nil {
		helpers.ClientError(w, http.StatusUnauthorized, err.Error())
	}

	// получаем из запроса номер задачки и код исх.кода и проверяем их
	taskID, err := strconv.Atoi(r.URL.Query().Get("task_id"))
	codeLang := r.URL.Query().Get("code_language")

	taskID, err = h.checkSystem.CheckTaskIdAndLang(taskID, codeLang) //вернется номер задачи либо ошибка_not_found
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	// Получаем код задачки и отправляем все на тестирование
	sourceCode := r.URL.Query().Get("source_code")
	if len(sourceCode) == 0 {
		helpers.ClientError(w, http.StatusBadRequest, "the code is missing")
	}

	// тестирование возвращает вердикт проверки
	verdict, err := h.checkSystem.TestUserCode(sourceCode, codeLang, taskID, userId)

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
