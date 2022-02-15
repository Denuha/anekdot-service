package auth

import (
	"time"

	"github.com/Denuha/anekdot-service/internal/models"
)

var DebugUser = models.User{
	ID:         2,
	UserName:   "quest",
	ExternalID: "",
	Realm:      "anekdot",
	CreateTime: time.Now(),
}
