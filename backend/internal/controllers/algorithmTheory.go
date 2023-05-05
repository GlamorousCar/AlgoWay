package controllers

import (
	"encoding/json"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
	"strconv"
)

type AlgoTheoryHandler struct {
	algoTheoryUseCase *usecase.AlgorithmTheoryUseCase
}

func NewAlgoTheoryHandler(algoTheoryUseCase *usecase.AlgorithmTheoryUseCase) *AlgoTheoryHandler {
	return &AlgoTheoryHandler{algoTheoryUseCase: algoTheoryUseCase}
}

// GetAlgorithmTheory Home godoc
//
//		@Summary		Теория к алгоритму
//		@Description	Получение теории к алгоритму по его id
//		@Tags			main
//		@Accept			json
//		@Success 200 {object} models.AlgorithmTheory
//		@Failure 404
//		@Failure 500
//	    @Param        algo_id   query     string _  "Получение задачи по id алгоритма"
//	    @Router			/theory [get]
func (h *AlgoTheoryHandler) GetAlgorithmTheory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.URL.Query().Get(algorithmId))
	if err != nil || id < 1 {
		helpers.NotFound(w)
		return
	}

	theory, err := h.algoTheoryUseCase.GetAlgoTheory(id)
	if err != nil {
		if errors.Is(err, helpers.ErrNoRecord) {
			helpers.NotFound(w)
		} else {
			helpers.ServerError(w, err)
		}
		return
	}

	jsonResp, err := json.Marshal(theory)
	if err != nil {
		helpers.ServerError(w, err)
	}
	_, err = w.Write(jsonResp)
	if err != nil {
		helpers.ErrorLogger.Println(err)
	}
}
