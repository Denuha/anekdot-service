package models

import "time"

type Anekdot struct {
	ID           int       `json:"id"`
	Sender       Sender    `json:"sender"`
	Text         string    `json:"text"`
	Likes        int       `json:"likes"`
	Dislikes     int       `json:"dislikes"`
	ExternalID   string    `json:"external_id"`
	CreateTime   time.Time `json:"create_time"`
	Status       int       `json:"status"`
	StatusString string    `json:"-"`
}

type Sender struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AnekdotVote struct {
	AnekdotID int   `json:"anekdot_id"`
	UserID    int64 `json:"user_id"`
	Value     int   `json:"value"`
}
