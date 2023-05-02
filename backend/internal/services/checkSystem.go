package services

import (
	"github.com/jackc/pgx/v4"
)

type Verdict struct {
	Abbr  string
	Title string
}

type CheckSystem interface {
	validateData(language string) error
	getTestData(task_id int) error
	runTests(code string, input_data, answers string) (verdict Verdict, err error)
	insertDataInDB(user_id int, task_id int, verdict Verdict)
}
type PythonCheckSystem struct {
}

func NewPythonCheckSystem(conn *pgx.Conn) *PythonCheckSystem {
	return &PythonCheckSystem{}
}
