package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/jackc/pgx/v4"
)

type UserRepository interface {
	Register(user models.RawUser) error
	Login(user models.LoginUser) (string, error)
	CheckIfUserExists(user models.RawUser) error
}

type userRepositoryPostgres struct {
	conn *pgx.Conn
}

func NewUserRepositoryPostgres(conn *pgx.Conn) *userRepositoryPostgres {
	return &userRepositoryPostgres{conn: conn}
}

func (repo *userRepositoryPostgres) Register(user models.RawUser) error {
	query := "INSERT INTO public.algo_user (login, email, hash_pass, is_active) VALUES ($1,$2,$3,TRUE)"

	_, err := repo.conn.Exec(context.Background(), query, user.Login, user.Email, user.Password)
	if err != nil { // Непредвиденные обстоятельства
		return errors.New(fmt.Sprintf("Unable to INSERT: %v\n", err))
	}
	return nil
}

func (repo *userRepositoryPostgres) CheckIfUserExists(user models.RawUser) error {
	checkIfUserExistQuery := "SELECT login,email from public.algo_user where login=$1 or email=$2"
	row := repo.conn.QueryRow(context.Background(), checkIfUserExistQuery, user.Login, user.Email)

	var login, email string // Если при запросе вернулись данные, значит пользователь существует
	err := row.Scan(&login, &email)

	if err == nil {
		return errors.New("пользователь уже существует") //Дальше регистрировать нельзя
	} else {
		return nil
	}
}

// Login Returns hashPass and error
func (repo *userRepositoryPostgres) Login(user models.LoginUser) (string, error) {
	query := `SELECT id, email, hash_pass from public.algo_user where email=$1`
	row := repo.conn.QueryRow(context.Background(), query, user.Email)

	var id int
	var email, hashPass string
	err := row.Scan(&id, &email, &hashPass)

	user.Id = id
	if err != nil {
		return "", err
	}

	return hashPass, nil
}
