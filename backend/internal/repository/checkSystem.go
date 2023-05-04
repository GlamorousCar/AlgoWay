package repository

import (
	"context"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

type CheckSystemRepo interface {
	GetTestData(taskID uint64) (*models.TestData, error)
}

type checkSystemRepoPostgres struct {
	conn *pgx.Conn
}

func NewCheckSystemRepoPostgres(conn *pgx.Conn) *checkSystemRepoPostgres {
	return &checkSystemRepoPostgres{conn: conn}
}

func (repo *checkSystemRepoPostgres) GetTestData(taskID uint64) (*models.TestData, error) {
	query := `SELECT input_data, output_data
		FROM test_data
		WHERE task_id=$1`

	rows, err := repo.conn.Query(context.Background(), query, taskID)
	if err != nil {
		return nil, err
	}

	inputData := make([]string, 0)
	outputData := make([]string, 0)

	for rows.Next() {
		var input, output string
		err := rows.Scan(&input, &output)
		if err != nil {
			return nil, err
		}

		inputData = append(inputData, input)
		outputData = append(outputData, output)
	}

	return &models.TestData{InputData: inputData, OutputData: outputData}, nil
}
