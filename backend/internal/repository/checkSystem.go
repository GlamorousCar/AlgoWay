package repository

import (
	"context"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

type CheckSystemRepository interface {
	Pass(id int) error
}

type checkSystemRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewCheckSystemRepositoryPostgres(conn *pgx.Conn) *checkSystemRepositoryPostgres {
	return &checkSystemRepositoryPostgres{conn: conn}
}

func (repo checkSystemRepositoryPostgres) Pass(id int) error {
	query := `SELECT t.id, a.title, t.content FROM algorithm AS a
	JOIN theory AS t
	ON a.theory_id=t.id
	WHERE a.id=$1`

	row := repo.conn.QueryRow(context.Background(), query, id)

	var theory = &models.AlgorithmTheory{}
	err := row.Scan(&theory.ID, &theory.Title, &theory.Content)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return helpers.ErrNoRecord
		} else {
			return err
		}
	}
	return nil
}
