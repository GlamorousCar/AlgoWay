package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"log"
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

var emailRegexp = regexp.MustCompile(EmailPattern)
var userRegexp = regexp.MustCompile("^[a-zA-Z0-9_]+$")
var passRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")

func isEmpty(str string) bool {
	return len(str) == 0
}

// TODO перенести в отдельный файл?
func validateLogin(login string) error {
	if isEmpty(login) {
		return errors.New("Поле должно быть заполнено")
		//"Пустое поле"
	} else if len(login) < 7 {
		//	слишком короткое имя
		return errors.New("Логин должен состоять из не менее 7 символов")
	} else if !userRegexp.MatchString(login) {
		return errors.New("Логин может состоять из символов латиницы, цифр и _ ")
		// логин должен состоять из букв цифр и нижнего подчеркивания
	}
	return nil
}
func validateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("Неверный формат email")
	}
	return nil
}

func validatePass(pass string) error {
	if len(pass) < 7 {
		return errors.New("Пароль должен быть больше 7 символов")
	} else if len(pass) > 71 {
		return errors.New("Пароль должен быть меньше 71 символа")
	} else if !passRegexp.MatchString(pass) {
		return errors.New("Неверный формат пароля")
	}
	return nil
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/auth/register" {
		app.notFound(w)
		return
	}
	rawUser := models.RawUser{}
	err := json.NewDecoder(r.Body).Decode(&rawUser)

	if err != nil {
		log.Println(err)
	}
	err = validateLogin(rawUser.Login)
	if err != nil {
		app.customError(w, err.Error())
		return
	}
	err = validateEmail(rawUser.Email)
	if err != nil {
		app.customError(w, err.Error())
		return
	}
	err = validatePass(rawUser.Password)
	if err != nil {
		app.customError(w, err.Error())
		return
	}
	err = app.PostgresqlConfig.UserModel.Register(rawUser)
	if err != nil {
		log.Println()
	}
	w.WriteHeader(http.StatusOK)

	//log.Println(!isEmpty(rawUser.Login))

	//log.Printf("%v,%T", rawUser.HashPass, rawUser.HashPass)
	userJson, err := json.Marshal(rawUser)
	w.Write(userJson)
	//rawUser = models.User{}
	//log.Println(r.Body)
	//log.Println(r.GetBody)
	//log.Println(r.Body.Read(r))
	//rawid := r.
	//fmt.Println(rawid())
}
