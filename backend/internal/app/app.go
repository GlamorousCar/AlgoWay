package app

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	host      = getEnvVar("host")
	dbPort    = getEnvVar("dbport")
	dbUser    = getEnvVar("dbuser")
	password  = getEnvVar("dbpass")
	dbname    = getEnvVar("dbname")
	ca        = getEnvVar("ca")
	secretKey = getEnvVar("secret_key")
)

func RunServer() {
	infoLogger := helpers.InfoLogger
	errorLogger := helpers.ErrorLogger

	os.Mkdir("temp", 0755)

	port := os.Getenv("PORT")

	// Проверка, что переменная окружения была найдена
	if port == "" {
		port = ":4000"
		infoLogger.Println("Переменная окружения PORT не установлена")
	} else {
		port = ":" + port
		infoLogger.Println("PORT", port)
	}

	infoLogger.Printf("Запуск веб-сервера на %s\n", port)

	conn, err := getConnection()
	if err != nil {
		log.Fatal(err)
	}

	mux := initServeMux(conn)
	srv := http.Server{
		Addr:     port,
		ErrorLog: errorLogger,
		Handler:  mux,
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

func getConnection() (*pgx.Conn, error) {
	rootCertPool := x509.NewCertPool()
	pem, err := os.ReadFile(ca)
	if err != nil {
		return nil, err
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		return nil, err
	}

	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=verify-full target_session_attrs=read-write",
		host, dbPort, dbname, dbUser, password)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		return nil, err
	}

	connConfig.TLSConfig = &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
	}

	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
