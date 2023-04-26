package database

import (
	"context"
)

func (db *DBImpl) GetVersion() (string, error) {
	var version string
	err := db.conn.QueryRow(context.Background(), "select version()").Scan(&version)

	return version, err
}
