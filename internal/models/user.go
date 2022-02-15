package models

import "time"

type User struct {
	ID         int64
	UserName   string
	ExternalID string
	Realm      string
	CreateTime time.Time
	IsAdmin    bool `json:"-"`
}

type UserRegistation struct {
	UserLogin
	Realm string
}

type UserLogin struct {
	Username string
	Realm    string
	Password string
}
