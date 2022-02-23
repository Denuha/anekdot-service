package models

import "time"

// User describe table `user`
type User struct {
	ID         int64     `json:"id"`
	UserName   string    `json:"username"`
	ExternalID string    `json:"external_id"`
	ChatID     *int64    `json:"chat_id"`
	Realm      string    `json:"realm"`
	CreateTime time.Time `json:"create_time"`
	IsAdmin    bool      `json:"-"`
}

type UserRegistation struct {
	Username string `json:"username"`
	Realm    string `json:"-"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Realm    string `json:"realm"`
	Password string `json:"password"`
}
