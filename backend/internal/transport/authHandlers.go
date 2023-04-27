package transport

import (
	"encoding/json"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"net/http"
)

func (h *MainHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
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
	err = helpers.ValidateLogin(rawUser.Login)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = helpers.ValidateEmail(rawUser.Email)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = helpers.ValidatePass(rawUser.Password)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = h.db.Register(rawUser)
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

type Token struct {
	Token string `json:"token"`
}

func (h *MainHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
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

	tokenValue, err := h.db.Login(loginUser)
	if err != nil {
		helpers.ClientError(w, http.StatusBadRequest, err.Error())
		return
	}
	token := Token{}
	token.Token = tokenValue

	if err != nil {
		helpers.ServerError(w, err)
		return
	}

	jsonResp, err := json.Marshal(token)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		helpers.ServerError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
