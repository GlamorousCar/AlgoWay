package main

import (
	"encoding/json"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"net/http"
)

const (
	EmailPattern string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/register" {
		app.notFound(w)
		return
	}
	rawUser := models.RawUser{}
	err := json.NewDecoder(r.Body).Decode(&rawUser)

	if err != nil {
		app.errorLogger.Println(err)
		app.clientError(w, http.StatusBadRequest, "Проблема с введенными данными, проверьте их корректность")
		return
	}
	err = app.Validator.validateLogin(rawUser.Login)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = app.Validator.validateEmail(rawUser.Email)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = app.Validator.validatePass(rawUser.Password)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = app.PostgresqlConfig.AuthService.Register(rawUser)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
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

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/login" {
		app.notFound(w)
		return
	}
	loginUser := models.LoginUser{}
	err := json.NewDecoder(r.Body).Decode(&loginUser)

	if err != nil {
		app.errorLogger.Println(err)
		app.clientError(w, http.StatusBadRequest, "Проблема с введенными данными, проверьте их корректность")
		return
	}

	tokenValue, err := app.PostgresqlConfig.AuthService.Login(loginUser)
	if err != nil {
		app.clientError(w, http.StatusBadRequest, err.Error())
		return
	}
	token := Token{}
	token.Token = tokenValue

	if err != nil {
		app.serverError(w, err)
		return
	}

	jsonResp, err := json.Marshal(token)
	if err != nil {
		app.serverError(w, err)
		return
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
