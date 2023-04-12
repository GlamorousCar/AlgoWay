package postgresql

import (
	"github.com/GlamorousCar/AlgoWay/pkg/models"
	"github.com/jackc/pgx/v4"
)

type UserModel struct {
	Conn *pgx.Conn
}

type JWTToken struct {
	id int
}

func (UserModel) Register(user models.User) (*models.User, error) {

	println("assa")
	// check id exist
	//check fileds
	// hash pass
	newUser := models.User{
		Id:       0,
		Login:    "",
		Email:    "",
		HashPass: "",
		IsActive: false,
	}
	return &newUser, nil
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
