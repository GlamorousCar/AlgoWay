package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type HomeModel struct {
	Conn *pgx.Conn
}

func (db *DBImpl) GetVersion() (string, error) {
	var version string
	err := db.pool.QueryRow(context.Background(), "select version()").Scan(&version)

	return version, err
}
