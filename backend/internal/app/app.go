package app

import (
	"context"
	"github.com/GlamorousCar/AlgoWay/internal/database"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/transport"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
)

func NewDB() (database.DB, error) {
	pool, err := pgxpool.Connect(context.Background(), "...")
	if err != nil {
		return nil, err
	}

	return database.NewDBImpl(pool), nil
}

func RunServer() {
	infoLogger := helpers.InfoLogger
	errorLogger := helpers.ErrorLogger

	port := os.Getenv("PORT")
	// Проверка, что переменная окружения была найдена
	if port == "" {
		port = ":4000"
		infoLogger.Println("Переменная окружения PORT не установлена")
	} else {
		infoLogger.Println("PORT:", port)
	}

	infoLogger.Printf("Запуск веб-сервера на %s\n", port)

	db, err := NewDB()
	if err != nil {
		log.Fatal(err)
	}

	transport.MakeMainHandler(db)

	srv := http.Server{
		Addr:     port,
		ErrorLog: errorLogger,
		Handler:  routes(),
	}

	err = srv.ListenAndServe()
	errorLogger.Fatal(err)
}
