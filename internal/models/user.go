package models

import "time"

type User struct {
	ID         int64     `json:"id"`
	UserName   string    `json:"username"`
	ExternalID string    `json:"external_id"`
	Realm      string    `json:"realm"`
	CreateTime time.Time `json:"create_time"`
	IsAdmin    bool      `json:"-"`
}

type UserRegistation struct {
	UserLogin
	Realm string `json:"realm"`
}

type UserLogin struct {
	Username string `json:"username"`
	Realm    string `json:"realm"`
	Password string `json:"password"`
}
