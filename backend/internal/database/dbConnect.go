package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"log"
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

func dbConnect() (*pgx.Conn, error) {

	// Штука для подключения сертификата SSL
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

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
