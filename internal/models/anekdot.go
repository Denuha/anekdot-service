package models

import "time"

type Anekdot struct {
	ID         int
	SenderID   int
	Text       string
	Rating     int
	ExternalID string
	CreateTime time.Time
	Status     int
}

type Sender struct {
	ID          int
	Name        string
	Description string
}
