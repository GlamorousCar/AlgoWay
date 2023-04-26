package database

import (
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

// DB TODO разбить на два интерфейса main и authDB
type DB interface {
	GetAlgoTheory(id int) (*models.AlgorithmTheory, error)
	GetVersion() (string, error)
	GetTasks(id int) (*[]models.Task, error)
	GetMenu() (*[]models.ThemeMenu, error)
	Register(user models.RawUser) error
	Login(user models.LoginUser) (string, error)
}
type DBImpl struct {
	conn *pgx.Conn
}

func NewDBImpl(conn *pgx.Conn) *DBImpl {
	return &DBImpl{conn: conn}
}
