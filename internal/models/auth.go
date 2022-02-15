package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Realm    string `json:"realm"`
	ID       int    `json:"id"`
}
