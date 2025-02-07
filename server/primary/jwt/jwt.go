package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")
var ErrInvalidToken = errors.New("invalid token")

func Create(email string) (string, error) {
	expire_at := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expire_at,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(mySigningKey)
}

func Parse(data string) (map[string]interface{}, error) {
	token, err := jwt.Parse(data, func(token *jwt.Token) (interface{}, error) {
		if !token.Valid {
			return nil, ErrInvalidToken
		}

		return token, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
