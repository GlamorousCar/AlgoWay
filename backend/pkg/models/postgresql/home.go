package postgresql

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type HomeModel struct {
	Conn *pgx.Conn
}

func (m HomeModel) GetVersion() (string, error) {
	var version string
	err := m.Conn.QueryRow(context.Background(), "select version()").Scan(&version)

	return version, err
}
