package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

type TaskModel struct {
	Conn *pgx.Conn
}

const defaultCapacity = 10

func (m TaskModel) GetTasks(id int) (*[]models.Task, error) {
	query := `SELECT id, is_solved, title, content 
			FROM task 
			WHERE algorithm_id=$1`

	rows, err := m.Conn.Query(context.Background(), query, id)

	algoTasks := make([]models.Task, 0, defaultCapacity)
	for rows.Next() {
		algoTask := models.Task{}
		err = rows.Scan(
			&algoTask.Id,
			&algoTask.IsSolved,
			&algoTask.Title,
			&algoTask.Content,
		)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, models.ErrNoRecord
			} else {
				return nil, err
			}
		} else {
			algoTasks = append(algoTasks, algoTask)
		}
	}

	return &algoTasks, nil
}
