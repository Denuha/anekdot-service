package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Realm    string `json:"realm"`
	ID       int    `json:"id"` // userID
}

type Login struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type AccessToken struct {
	AccessToken string `json:"access_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type Session struct {
	UserID                 int
	AccessToken            string
	AccessTokenCreateTime  time.Time
	RefreshToken           string
	RefreshTokenCreateTime time.Time
}

type SessionInsert struct {
	UserID       int
	AccessToken  string
	RefreshToken string
}
