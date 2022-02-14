package models

import "time"

type Anekdot struct {
	ID     int
	Sender Sender
	Text   string
	//Rating       int
	Likes        int
	Dislikes     int
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

type AnekdotVote struct {
	AnekdotID int
	UserID    int64
	Value     int
}
