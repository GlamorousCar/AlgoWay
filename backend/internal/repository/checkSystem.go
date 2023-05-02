package repository

import (
	"context"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/jackc/pgx/v4"
)

type CheckSystemRepository interface {
	GetTask(id int) (int, error)
}

type checkSystemRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewCheckSystemRepositoryPostgres(conn *pgx.Conn) *checkSystemRepositoryPostgres {
	return &checkSystemRepositoryPostgres{conn: conn}
}

func (repo checkSystemRepositoryPostgres) GetTask(id int) (int, error) {
	query := `SELECT * FROM task WHERE id = $1`

	row := repo.conn.QueryRow(context.Background(), query, id)

	var taskId int
	err := row.Scan(&taskId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, helpers.ErrNoRecord
		} else {
			return 0, err
		}
	}
	return taskId, nil
}
