package transport

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/app"
	"github.com/GlamorousCar/AlgoWay/internal/database"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"log"
	"net/http"
	"strconv"
)

const algorithmId = "algo_id"

type MainHandler struct {
	db database.DB
}

func MakeMainHandler(db database.DB) *MainHandler {
	return &MainHandler{db: db}
}

func (h *MainHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := 3 // parse id from request

	user, err := h.db.GetUserById(id)
	if err != nil {
		// handle error
	}
	log.Println(user)
	return
	// return user as JSON
}

func (h *MainHandler) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.NotFound(w)
		return
	}

	version, err := h.db.GetVersion()

	if err != nil {
		helpers.ServerError(w, err)
	}
	_, err = w.Write([]byte(fmt.Sprintf("Успешное подключение\n %s", version)))
	if err != nil {
		app.errorLogger.Println(err)
	}
}

func (h *MainHandler) GetThemesMenu(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/themes/menu" {
		app.notFound(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	menus, err := database.ThemeMenuModel.Get()
	h.db.
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

func GetAlgorithmTheory(w http.ResponseWriter, r *http.Request) {
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

func GetAlgorithmTasks(w http.ResponseWriter, r *http.Request) {
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

	tasks, err := database.TaskModel.GetTasks(algoId)
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
