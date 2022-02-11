package models

type StatusAnekdot int

const (
	StatusAnekdotVerification StatusAnekdot = iota + 1
	StatusAnekdotOK
	StatusAnekdotBan
)

func (s StatusAnekdot) String() string {
	return [...]string{"На проверке", "Разрешен", "Отклонен"}[s]
}
