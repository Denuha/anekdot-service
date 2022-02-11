package models

type Anekdot struct {
	SenderID   int
	Text       string
	Rating     int
	ExternalID string
}

type Sender struct {
	ID          int
	Description string
	Name        string
}
