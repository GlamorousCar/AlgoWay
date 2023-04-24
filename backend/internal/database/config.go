package database

import (
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB interface {
	GetAlgoTheory(id int) (*models.AlgorithmTheory, error)
	GetVersion() (string, error)
	Get() (*[]models.ThemeMenu, error)
}
type DBImpl struct {
	pool *pgxpool.Pool
}

func NewDBImpl(pool *pgxpool.Pool) *DBImpl {
	return &DBImpl{pool: pool}
}

//func (db *DBImpl) GetUserById(id int) (*models.User, error) {
//	conn, err := db.pool.Acquire(context.Background())
//	if err != nil {
//		return nil, err
//	}
//	defer conn.Release()
//
//	// ...
//}
