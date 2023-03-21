package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	log.Println("Запуск сервера на http://127.0.0.1:4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
