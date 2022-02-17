package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Realm    string `json:"realm"`
	ID       int    `json:"id"`
}

type Login struct {
	AccessToken string        `json:"access_token"`
	Expires     time.Duration `json:"expires_seconds" swaggertype:"integer"`
}
