package postgresql

import "github.com/jackc/pgx/v4"

type Config struct {
	AlgorithmTheoryModel *AlgorithmTheoryModel
	TaskModel            *TaskModel
	HomeModel            *HomeModel
	ThemeMenuModel       *ThemeMenuModel
}

func NewConfig(conn *pgx.Conn) *Config {
	algorithmTheoryModel := &AlgorithmTheoryModel{Conn: conn}
	taskModel := &TaskModel{Conn: conn}
	homeModel := &HomeModel{Conn: conn}
	themeMenuModel := &ThemeMenuModel{Conn: conn}

	return &Config{
		AlgorithmTheoryModel: algorithmTheoryModel,
		TaskModel:            taskModel,
		HomeModel:            homeModel,
		ThemeMenuModel:       themeMenuModel,
	}
}