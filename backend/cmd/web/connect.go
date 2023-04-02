package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
)

var (
	host     = getEnvVar("host")
	dbport   = getEnvVar("dbport")
	user     = getEnvVar("dbuser")
	password = getEnvVar("dbpass")
	dbname   = getEnvVar("dbname")
	ca       = getEnvVar("ca")
)

var conn *pgx.Conn

func dbConnect() error {

	// Штука для подключения сертификата SSL
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(ca)
	if err != nil {
		panic(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		panic("Failed to append PEM.")
	}

	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=verify-full target_session_attrs=read-write",
		host, dbport, dbname, user, password)

	connConfig, err := pgx.ParseConfig(connString)
	if err != nil {
		fmt.Printf("Unable to parse config: %v\n", err)
		return err
	}

	connConfig.TLSConfig = &tls.Config{
		RootCAs:            rootCertPool,
		InsecureSkipVerify: true,
	}

	conn, err = pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)

		return err
	}
	return err
}
