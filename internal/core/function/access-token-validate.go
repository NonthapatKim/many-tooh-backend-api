package function

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateAccessToken(tokenString *string) (string, error) {
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return "", errors.New("secret key not found in environment variables")
	}

	token, err := jwt.ParseWithClaims(*tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return "", fmt.Errorf("error while parsing token: %v", err)
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return "", errors.New("token has expired")
		}

		userId := claims.Subject
		if userId == "" {
			return "", errors.New("user_id not found in token claims")
		}
		return userId, nil
	}

	return "", errors.New("invalid token")
}
