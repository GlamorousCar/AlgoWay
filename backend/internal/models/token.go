package models

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type Token struct {
	Token string `json:"token"`
}
