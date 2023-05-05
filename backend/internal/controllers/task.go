package controllers

import (
	"encoding/json"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
	"strconv"
)

const algorithmId = "algo_id"

type TaskHandler struct {
	taskUseCase *usecase.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecase.TaskUseCase) *TaskHandler {
	return &TaskHandler{taskUseCase: taskUseCase}
}

// GetAlgorithmTasks Home godoc
//
//		@Summary		Список задач по алгоритму
//		@Description	Получение заданий к определенному алгоритму. Пока есть ошибка связанная с возвращаемым полем is_solved, будет исправлено скоро
//		@Tags			main
//		@Accept			json
//	    @Param        algo_id   query     string  3  "Получение задач по id алгоритма"
//	    @Success		200	{array} models.Task "Возвращается список задач"
//	    @Failure		500
//		@Router			/task [get]
func (h *TaskHandler) GetAlgorithmTasks(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/task" {
		helpers.NotFound(w)
		return
	}
	rawId := r.URL.Query().Get(algorithmId)
	algoId, err := strconv.Atoi(rawId)
	if err != nil || algoId < 1 {
		helpers.NotFound(w)
		return
	}

	tasks, err := h.taskUseCase.GetTasks(algoId)
	if err != nil {
		if errors.Is(err, helpers.ErrNoRecord) {
			helpers.NotFound(w)
		} else {
			helpers.ServerError(w, err)
		}
	}

	jsonStr, err := json.Marshal(tasks)
	if err != nil {
		helpers.ServerError(w, err)
	}

	_, err = w.Write(jsonStr)
	if err != nil {
		helpers.ErrorLogger.Println(err)
	}
}
