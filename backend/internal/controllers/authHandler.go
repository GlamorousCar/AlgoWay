package controllers

import (
	"encoding/json"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type AuthHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewAuthHandler(userUseCase *usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{userUseCase: userUseCase}
}

// RegisterUser auth godoc
//
//		    @Summary		регистрация пользователя
//	     @Description	после успешной регистрации возвращается сообщение об успешном выполнении
//		    @Tags			auth
//		    @Accept			json
//		    @Produce		plain
//			@Param			message body		models.RawUser	true "User Info"
//		    @Success		200	{string} 3 Регистрация прошла успешно
//		    @Failure		400
//		    @Failure		404
//		    @Failure		500
//		    @Router			/auth/register [post]
func (h *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/register" {
		helpers.NotFound(w)
		return
	}

	rawUser := models.RawUser{}
	err := json.NewDecoder(r.Body).Decode(&rawUser)

	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, "Проблема с введенными данными, проверьте их корректность")
		return
	}

	err = h.userUseCase.ValidateUser(rawUser)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.userUseCase.Register(rawUser)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Регистрация прошла успешно"))
	if err != nil {
		return
	}
}

// LoginUser auth godoc
//
//		@Summary		авторизация пользователя
//		@Description	после успешной авторизации возвращается токен
//		@Tags			auth
//		@Accept			json
//		@Produce		plain
//		@Param			message body		models.LoginUser	true "Account Info"
//	    @Success		200	{object} models.Token
//	    @Failure		400
//	    @Failure		404
//	    @Failure		500
//	    @Router			/auth/login [post]
func (h *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/login" {
		helpers.NotFound(w)
		return
	}

	loginUser := models.LoginUser{}
	err := json.NewDecoder(r.Body).Decode(&loginUser)

	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, "Проблема с введенными данными, проверьте их корректность")
		return
	}

	tokenValue, err := h.userUseCase.Login(loginUser)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}
	token := models.Token{}
	token.Token = tokenValue

	jsonResp, err := json.Marshal(token)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResp)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}

}
