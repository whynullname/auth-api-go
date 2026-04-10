package auth

import "github.com/golang-jwt/jwt/v5"

type UserDataInput struct {
	Login    string
	Password string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Login string
}
