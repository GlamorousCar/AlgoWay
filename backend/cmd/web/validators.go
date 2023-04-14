package main

import (
	"errors"
	"regexp"
)

type Validator struct {
}

var emailRegexp = regexp.MustCompile(EmailPattern)
var userRegexp = regexp.MustCompile("^[a-zA-Z0-9_]+$")
var passRegexp = regexp.MustCompile("^[a-zA-Z0-9!#$%&'*+/=?^_`{|}~.-]+$")
var minPasswordLen = 7
var maxPasswordLen = 71
var minLoginLen = 7

func isEmpty(str string) bool {
	return len(str) == 0
}

func (m Validator) validateLogin(login string) error {
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
func (m Validator) validateEmail(email string) error {
	if !emailRegexp.MatchString(email) {
		return errors.New("неверный формат email")
	}
	return nil
}

func (m Validator) validatePass(pass string) error {
	if len(pass) < minPasswordLen {
		return errors.New("пароль должен быть больше 7 символов")
	} else if len(pass) > maxPasswordLen {
		return errors.New("пароль должен быть меньше 71 символа")
	} else if !passRegexp.MatchString(pass) {
		return errors.New("неверный формат пароля")
	}
	return nil
}
