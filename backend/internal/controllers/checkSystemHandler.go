package controllers

import (
	"github.com/GlamorousCar/AlgoWay/internal/usecase"
	"net/http"
)

type CheckSystemHandler struct {
	userUseCase *usecase.UserUseCase
	checkSystem *usecase.CheckSystemUseCase
}

func NewCheckSystemHandler(checkSystem *usecase.CheckSystemUseCase) *CheckSystemHandler {
	return &CheckSystemHandler{checkSystem: checkSystem}
}

func (h *CheckSystemHandler) CheckTask(w http.ResponseWriter, r *http.Request) {

}
