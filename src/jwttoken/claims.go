package jwttoken

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Data
	jwt.RegisteredClaims
}
