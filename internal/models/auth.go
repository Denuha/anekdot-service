package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserName string
	Realm    string
	ID       int
}
