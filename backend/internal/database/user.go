package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

type AuthService struct {
	Conn *pgx.Conn
}

func hashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (db *DBImpl) Register(user models.RawUser) error {
	checkIfUserExistQuery := "SELECT login,email from public.algo_user where login=$1 or email=$2"
	val := db.conn.QueryRow(context.Background(), checkIfUserExistQuery, user.Login, user.Email)

	var login, email string // Если при запросе вернулись данные, значит пользователь существует
	err := val.Scan(&login, &email)

	if err == nil {
		return errors.New("пользователь уже существует") //Дальше регистрировать нельзя
	}

	hash, err := hashAndSalt(user.Password) // Полученный запрос хэшируем
	log.Println()
	if err != nil {
		return errors.New("проблема с паролем")
	}

	query := "INSERT INTO public.algo_user (login, email, hash_pass, is_active) VALUES ($1,$2,$3,TRUE)"

	_, err = db.conn.Exec(context.Background(), query, user.Login, user.Email, hash)
	if err != nil { // Непредвиденные обстоятельства
		return errors.New(fmt.Sprintf("Unable to INSERT: %v\n", err))
	}
	return nil
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

const tokenTTL = 12 * time.Hour

// ParseToken по токену получает id пользователя
func (db *DBImpl) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func (db *DBImpl) GenerateToken(user models.LoginUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}
func (db *DBImpl) Login(user models.LoginUser) (string, error) {
	query := `SELECT id, email, hash_pass from public.algo_user where email=$1`

	val := db.conn.QueryRow(context.Background(), query, user.Email)
	var id int
	var email, hashPass string
	err := val.Scan(&id, &email, &hashPass)
	user.Id = id
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(user.Password)) // проверка паролей
	if err != nil {
		return "", err
	}
	token, err := db.GenerateToken(user)
	return token, err

}
