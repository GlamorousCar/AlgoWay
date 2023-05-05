package models

type RawUser struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `json:"id"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	HashPass string `json:"pass"`
	IsActive bool   `json:"is_active"`
}

type LoginUser struct {
	Id       int    `json:"-"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
