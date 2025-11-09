package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/Denuha/anekdot-service/internal/models"
	"github.com/dgrijalva/jwt-go"
)

// ParseToken Вернет ошибку, если токен просрочен
func ParseToken(parseToken string, signedKey []byte) (*models.Claims, error) {
	token, err := jwt.ParseWithClaims(
		parseToken, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return signedKey, nil
		})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*models.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("ErrInvalidAccessToken")
}

func GenerateToken(now time.Time, userID int, user *models.UserLogin, signedKey string, tokenExpires time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &models.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Unix() + int64(tokenExpires.Seconds()),
			IssuedAt:  now.Unix(),
		},
		UserName: user.Username,
		Realm:    user.Realm,
		ID:       userID,
	})

	return token.SignedString([]byte(signedKey))
}
