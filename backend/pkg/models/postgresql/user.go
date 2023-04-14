package postgresql

import (
	"context"
	"errors"
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserModel struct {
	Conn *pgx.Conn
}

type JWTToken struct {
	id int
}

func hashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (m UserModel) Register(user models.RawUser) error {
	checkIfUserExistQuery := `SELECT login,email from public.algo_user where login=$1 or email=$2`
	val := m.Conn.QueryRow(context.Background(), checkIfUserExistQuery, user.Login, user.Email)

	var login, email string // Если при запросе вернулись данные, значит пользователь существует
	err := val.Scan(&login, &email)

	if err == nil {
		//log.Println("Пользователь уже существует")
		return errors.New("пользователь уже существует") //Дальше регистрировать нельзя
	}

	hash := hashAndSalt(user.Password) // Полученный запрос хэшируем

	query := "INSERT INTO public.algo_user (login, email, hash_pass, is_active) VALUES ($1,$2,$3,TRUE)"

	_, err = m.Conn.Exec(context.Background(), query, user.Login, user.Email, hash)
	if err != nil {
		log.Printf("Unable to INSERT: %v\n", err) // Непредвиденные обстоятельства
		return err
	}
	return nil
}

//func checkJWT() {
//
//}
//
//func createJWT() {
//
//}
//
//func loginUser() {
//
//}
