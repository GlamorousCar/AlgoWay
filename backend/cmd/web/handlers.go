package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"net/http"
	"regexp"
	"strconv"
)

const algorithmId = "algo_id"

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	version, err := app.PostgresqlConfig.HomeModel.GetVersion()

	if err != nil {
		app.serverError(w, err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("Успешное подключение\n %s", version)))
	if err != nil {
		app.errorLogger.Println(err)
	}
}

func (app *application) getThemesMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/themes/menu" {
		app.notFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	menus, err := app.PostgresqlConfig.ThemeMenuModel.Get()
	if err != nil {
		app.serverError(w, err)
		return
	}

	jsonStr, err := json.Marshal(menus)
	if err != nil {
		app.serverError(w, err)
		return
	}
	_, err = w.Write(jsonStr)
	if err != nil {
		app.errorLogger.Println(err)
	}
}

func (app *application) getAlgorithmTheory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get(algorithmId))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	theory, err := app.PostgresqlConfig.AlgorithmTheoryModel.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	jsonResp, err := json.Marshal(theory)
	if err != nil {
		app.serverError(w, err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		app.errorLogger.Println(err)
	}
}

func (app *application) getAlgorithmTasks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/task" {
		app.notFound(w)
		return
	}
	rawId := r.URL.Query().Get(algorithmId)
	algoId, err := strconv.Atoi(rawId)
	if err != nil || algoId < 1 {
		app.notFound(w)
		return
	}

	tasks, err := app.PostgresqlConfig.TaskModel.GetTasks(algoId)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
	}

	jsonStr, err := json.Marshal(tasks)
	if err != nil {
		app.serverError(w, err)
	}

	_, err = w.Write(jsonStr)
	if err != nil {
		app.errorLogger.Println(err)
	}
}

const (
	EmailPattern string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var rxEmail = regexp.MustCompile(EmailPattern)

func IsEmail(str string) bool {
	return rxEmail.MatchString(str)
}

func isEmpty(str string) bool {
	return len(str) > 0
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/register" {
		app.notFound(w)
		return
	}
	rawUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&rawUser)
	if err != nil {
		panic(err)
	}
	userJson, err := json.Marshal(rawUser)
	w.Write(userJson)
	//rawUser = models.User{}
	//log.Println(r.Body)
	//log.Println(r.GetBody)
	//log.Println(r.Body.Read(r))
	//rawid := r.
	//fmt.Println(rawid())
}
