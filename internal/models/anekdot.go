package models

type Anekdot struct {
	SenderID int
	Text     string
	Rating   int
}

type Sender struct {
	ID int
}
