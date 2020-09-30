package models

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
