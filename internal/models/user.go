package models

import "time"

type User struct {
	ID         int64
	UserName   string
	ExternalID string
	Realm      string
	CreateTime time.Time
}
