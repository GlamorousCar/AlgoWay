package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(conn, context.Background())

	err := dbConnect()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/themes/menu", getThemesMenu)
	mux.HandleFunc("/theory", getAlgorithmTheory)
	mux.HandleFunc("/task", getAlgorithmTasks)

	port := os.Getenv("PORT")

	// Проверка, что переменная окружения была найдена
	if port == "" {
		log.Println("Переменная окружения PORT не установлена")
		log.Println("Пробую запуститься на 4000 порту")
		err = http.ListenAndServe(":4000", mux)
	} else {
		log.Println("PORT:", port)
		err = http.ListenAndServe(":"+port, mux)
	}

	log.Println("Запуск сервера на http://127.0.0.1")

	log.Fatal(err)
}
