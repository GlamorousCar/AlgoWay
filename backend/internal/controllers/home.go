package controllers

import (
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type HomeHandler struct {
	homeUseCase *usecase.HomeUseCase
}

func NewHomeHandler(homeUseCase *usecase.HomeUseCase) *HomeHandler {
	return &HomeHandler{homeUseCase: homeUseCase}
}

// Home godoc
//
//	@Summary		Главная страница
//	@Description	Страница с проверкой подключения к бд (тестовая)
//	@Tags			main
//	@Accept			json
//	@Produce		plain
//	@Success		200	{string}	string	"Успешное подключение -  версия базы данных"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/ [get]
func (h *HomeHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.NotFound(w)
		return
	}

	version, err := h.homeUseCase.GetVersion()

	if err != nil {
		helpers.ServerError(w, err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("Успешное подключение\n %s", version)))
	if err != nil {
		helpers.ErrorLogger.Println(err)
	}
}
