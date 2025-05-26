package service

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(LOGIN_EXPIRED_DURATION).Unix(),
		"iss":      APPLICATION_NAME,
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	tokenString, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		log.Println("error generate token,", err.Error())
		return "", nil
	}

	return tokenString, nil
}
