package database

import (
	"context"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

func (db *DBImpl) GetAlgoTheory(id int) (*models.AlgorithmTheory, error) {
	query := `SELECT t.id, a.title, t.content FROM algorithm AS a
	JOIN theory AS t
	ON a.theory_id=t.id
	WHERE a.id=$1`

	row := db.conn.QueryRow(context.Background(), query, id)

	var theory = &models.AlgorithmTheory{}
	err := row.Scan(&theory.ID, &theory.Title, &theory.Content)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return theory, nil
}
