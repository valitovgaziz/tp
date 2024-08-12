package models

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("SECRET_KEY"))

type Crenetials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}
