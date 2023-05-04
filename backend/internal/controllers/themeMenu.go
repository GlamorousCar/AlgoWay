package controllers

import (
	"encoding/json"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type ThemeMenuHandler struct {
	themeMenuUseCase *usecase.ThemeMenuUseCase
}

func NewThemeMenuHandler(themeMenuUseCase *usecase.ThemeMenuUseCase) *ThemeMenuHandler {
	return &ThemeMenuHandler{themeMenuUseCase: themeMenuUseCase}
}

// GetThemeMenu  godoc
//
//	@Summary		Меню сайта
//	@Description	Получение всех тем и подтем меню сайта
//	@Tags			main
//	@Accept			json
//	@Success		200	{array} models.ThemeMenu "Возвращаются темы и подтемы"
//	@Failure		404
//	@Failure		500
//	@Router			/themes/menu [get]
func (h *ThemeMenuHandler) GetThemeMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/themes/menu" {
		helpers.NotFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	menus, err := h.themeMenuUseCase.GetMenu()

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	jsonStr, err := json.Marshal(menus)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = w.Write(jsonStr)
	if err != nil {
		helpers.ErrorLogger.Println(err)
	}
}
