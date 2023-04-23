package main

import (
	"github.com/GlamorousCar/AlgoWay/pkg/models/postgresql"
	"log"
)

type application struct {
	errorLogger      *log.Logger
	infoLogger       *log.Logger
	PostgresqlConfig *postgresql.Config
	Validator        *Validator
}
