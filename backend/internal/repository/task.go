package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

const defaultCapacity = 10

type TaskRepository interface {
	GetTasks(id int, userId int) (*[]models.Task, error)
}

type taskRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewTaskRepositoryPostgres(conn *pgx.Conn) *taskRepositoryPostgres {
	return &taskRepositoryPostgres{conn: conn}
}

func (repo *taskRepositoryPostgres) GetTasks(id int, userId int) (*[]models.Task, error) {
	query := `SELECT id, CASE
    WHEN task.id in (select task_id from attempt join verdict v2 ON attempt.verdict_id = v2.id where user_id = $2 and v2.abbr = 'OK') THEN 'OK'
    WHEN task.id in (select task_id from attempt join verdict v2 ON attempt.verdict_id = v2.id where user_id = $2 and v2.abbr = 'WA') THEN 'WA'
    ELSE 'NotSolved'
  	END is_solved ,title,content
	from task WHERE algorithm_id=$1;`

	rows, err := repo.conn.Query(context.Background(), query, id, userId)

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
				return nil, helpers.ErrNoRecord
			} else {
				return nil, err
			}
		} else {
			algoTasks = append(algoTasks, algoTask)
		}
	}

	return &algoTasks, nil
}
