package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type HomeRepository interface {
	GetVersion() (string, error)
}

type homeRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewHomeRepositoryPostgres(conn *pgx.Conn) *homeRepositoryPostgres {
	return &homeRepositoryPostgres{conn: conn}
}

func (repo *homeRepositoryPostgres) GetVersion() (string, error) {
	var version string
	err := repo.conn.QueryRow(context.Background(), "select version()").Scan(&version)

	return version, err
}
