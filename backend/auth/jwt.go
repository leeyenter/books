package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/leeyenter/books/backend/utils"
	"time"
)

var key = []byte(utils.GetEnv("JWT_SECRET_KEY", "pQH9aGxgJpVhyYDNZc2moJQEm!wQHDC!AwqUU9@Eyz2CE3QFBGDARfc9_7Raon@F"))

func CreateJWT(remoteAddr string, expDuration time.Duration) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"remoteAddr": remoteAddr,
		"exp":        time.Now().Add(expDuration).Unix(),
	})

	return t.SignedString(key)
}

func parseJWT(tokenString string) (string, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	return claims["remoteAddr"].(string), nil
}
