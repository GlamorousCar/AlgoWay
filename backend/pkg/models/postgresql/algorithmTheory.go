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
	query := `SELECT t.id, a.title, t.content FROM algorithm AS a
	JOIN theory AS t
	ON a.theory_id=t.id
	WHERE a.id=$1`

	row := m.Conn.QueryRow(context.Background(), query, id)

	var theory = &models.AlgorithmTheory{}
	err := row.Scan(&theory.ID, &theory.Title, &theory.Content)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return theory, nil
}
