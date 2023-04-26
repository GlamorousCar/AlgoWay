package helpers

import (
	"errors"
	"regexp"
)

type Validator struct {
}

const (
	emailPattern string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
)

var emailRegexp = regexp.MustCompile(emailPattern)
var userRegexp = regexp.MustCompile("^[a-zA-Z0-9_]+$")
var passRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")

const minPasswordLen = 7
const maxPasswordLen = 71
const minLoginLen = 7

func isEmpty(str string) bool {
	return len(str) == 0
}

func ValidateLogin(login string) error {
	if isEmpty(login) {
		return errors.New("поле должно быть заполнено")
		//"Пустое поле"
	} else if len(login) < minLoginLen {
		//	слишком короткое имя
		return errors.New("логин должен состоять из не менее 5 символов")
	} else if !userRegexp.MatchString(login) {
		return errors.New("Логин может состоять из символов латиницы, цифр и _ ")
		// логин должен состоять из букв цифр и нижнего подчеркивания
	}
	return nil
}
func ValidateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("неверный формат email")
	}
	return nil
}

func ValidatePass(pass string) error {
	if len(pass) < minPasswordLen {
		return errors.New("пароль должен быть больше 7 символов")
	} else if len(pass) > maxPasswordLen {
		return errors.New("пароль должен быть меньше 71 символа")
	} else if !passRegexp.MatchString(pass) {
		return errors.New("неверный формат пароля")
	}
	return nil
}
