package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"github.com/jackc/pgx/v4"
)

type AlgorithmTheoryModel struct {
	Conn *pgx.Conn
}

func (m AlgorithmTheoryModel) Get(id int) (*models.AlgorithmTheory, error) {
	query := "SELECT id,content FROM theory WHERE algorithm_id=$1"
	row := m.Conn.QueryRow(context.Background(), query, id)

	var theory = &models.AlgorithmTheory{}
	err := row.Scan(&theory.ID, &theory.Content)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return theory, nil
}
