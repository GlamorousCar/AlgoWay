package app

import (
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"log"
)

type Application struct {
	errorLogger *log.Logger
	infoLogger  *log.Logger
	Validator   *helpers.Validator
}
