package postgresql

import (
	"context"
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
	checkQuery := `SELECT login,email from public.algo_user where login='$1' or email='$2'`
	val, err := m.Conn.Query(context.Background(), checkQuery, user.Login, user.Email)
	test := models.User{}
	val.Scan(&test.Login, &test.Email)
	//log.Println(g)
	//log.Println(val)
	if val != nil {
		log.Printf("Пользователь уже существует: %v\n", err)
		return err
	}
	//
	//query := `INSERT INTO public.algo_user (login, email, hash_pass, is_active) VALUES ($1,$2,$3,TRUE)`
	//hash := hashAndSalt(user.Password)
	//log.Println("Salted Hash", hash)
	//v, err := m.Conn.Query(context.Background(), query, user.Login, user.Email, hash)
	//log.Println(v)
	//if err != nil {
	//	log.Printf("Unable to INSERT: %v\n", err)
	//	return err
	//}
	return nil
}

//	var newUser = new(user)
//	// Проверка полей
//
//
//
//
//	log.Println(newUser)
////}
//

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
