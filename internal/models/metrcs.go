package models

type Metrics struct {
	NumberUsers     int `json:"number_users"`
	NumberAnekdots  int `json:"number_anekdots"`
	NumberUserVotes int `json:"number_user_votes"`
}
