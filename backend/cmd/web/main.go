package main

import "github.com/GlamorousCar/AlgoWay/internal/app"
import _ "github.com/GlamorousCar/AlgoWay/docs"

// @title Swagger Algoway API
// @version 1.0
// @description Документация для нашего студенческого проекта - сервиса для изучения алгоритмов

// @host 127.0.0.1:4000
// @BasePath /
func main() {
	app.RunServer()
}
