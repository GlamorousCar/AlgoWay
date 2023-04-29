package usecase

import (
	"errors"
	"github.com/GlamorousCar/AlgoWay/internal/helpers"
	"github.com/GlamorousCar/AlgoWay/internal/models"
	"github.com/GlamorousCar/AlgoWay/internal/repository"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const tokenTTL = 12 * time.Hour

type UserUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) Register(user models.RawUser) error {
	err := u.repo.CheckIfUserExists(user)
	if err != nil {
		return err
	}

	hash, err := helpers.HashAndSalt(user.Password) // Полученный запрос хэшируем
	if err != nil {
		return errors.New("проблема с паролем")
	}

	user.Password = hash

	err = u.repo.Register(user)
	return err
}

// Login Returns Token and error
func (u *UserUseCase) Login(user models.LoginUser) (string, error) {
	hashPass, err := u.repo.Login(user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(user.Password)) // проверка паролей
	if err != nil {
		return "", err
	}

	token, err := generateToken(user)
	return token, err
}

func (u *UserUseCase) ValidateUser(rawUser models.RawUser) error {
	err := helpers.ValidateLogin(rawUser.Login)
	if err != nil {
		return err
	}

	err = helpers.ValidateEmail(rawUser.Email)
	if err != nil {
		return err
	}

	err = helpers.ValidatePass(rawUser.Password)
	if err != nil {
		return err
	}

	return nil
}

func generateToken(user models.LoginUser) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.Id,
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

// parseToken По токену получает id пользователя
func parseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &models.TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*models.TokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}
