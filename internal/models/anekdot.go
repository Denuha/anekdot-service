package models

import "time"

type Anekdot struct {
	ID           int
	Sender       Sender
	Text         string
	Rating       int
	ExternalID   string
	CreateTime   time.Time
	Status       int
	StatusString string
}

type Sender struct {
	ID          int
	Name        string
	Description string
}
