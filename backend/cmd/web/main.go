package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/GlamorousCar/AlgoWay/pkg/models/postgresql"
)

func main() {
	infoLogger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLogger := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	port := os.Getenv("PORT")
	// Проверка, что переменная окружения была найдена
	if port == "" {
		port = ":4000"
		infoLogger.Println("Переменная окружения PORT не установлена")
	} else {
		infoLogger.Println("PORT:", port)
	}

	conn, err := dbConnect()
	if err != nil {
		errorLogger.Fatal(err)
	}
	defer conn.Close(context.Background())

	postgresqlConfig := postgresql.NewConfig(conn)

	app := &application{
		errorLogger:      errorLogger,
		infoLogger:       infoLogger,
		PostgresqlConfig: postgresqlConfig,
	}

	infoLogger.Printf("Запуск веб-сервера на %s\n", port)
	srv := http.Server{
		Addr:     port,
		ErrorLog: errorLogger,
		Handler:  app.routes(),
	}

	err = srv.ListenAndServe()
	errorLogger.Fatal(err)
}

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
