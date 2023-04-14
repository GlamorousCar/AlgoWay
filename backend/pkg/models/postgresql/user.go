package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	Conn *pgx.Conn
}

type JWTToken struct {
	id int
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (m UserModel) Register(user models.RawUser) error {
	checkIfUserExistQuery := `SELECT login,email from public.algo_user where login=$1 or email=$2`
	val := m.Conn.QueryRow(context.Background(), checkIfUserExistQuery, user.Login, user.Email)

	var login, email string // Если при запросе вернулись данные, значит пользователь существует
	err := val.Scan(&login, &email)

	if err == nil {
		return errors.New("пользователь уже существует") //Дальше регистрировать нельзя
	}

	hash, err := hashAndSalt(user.Password) // Полученный запрос хэшируем
	if err != nil {
		return errors.New("Проблема с паролем")
	}

	query := "INSERT INTO public.algo_user (login, email, hash_pass, is_active) VALUES ($1,$2,$3,TRUE)"

	_, err = m.Conn.Exec(context.Background(), query, user.Login, user.Email, hash)
	if err != nil { // Непредвиденные обстоятельства
		return errors.New(fmt.Sprintf("Unable to INSERT: %v\n", err))
	}
	return nil
}
