package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(payload interface{}, secretJWTkey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = now.Add(time.Hour).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := token.SignedString([]byte(secretJWTkey))

	if err != nil {
		return "", fmt.Errorf("generating JWT Token failed: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(token string, signedJWTkey string) (interface{}, error) {
	t, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTkey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims["sub"], nil
}
