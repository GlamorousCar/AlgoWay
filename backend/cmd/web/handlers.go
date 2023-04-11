package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"net/http"
	"strconv"
)

const algorithmId = "algo_id"
const theory_id = "theory_id"

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
	id, err := strconv.Atoi(r.URL.Query().Get(theory_id))
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
