package auth

import (
	"time"

	"github.com/arianaw15/birdie-talk/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWTToken(secret []byte, userId int) (string, error) {
	expirationTime := time.Second * time.Duration(config.Envs.JWTExpiration) // Token valid for 24 hours
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"expireAt": time.Now().Add(expirationTime).Unix(), // Token valid for 24 hours
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
