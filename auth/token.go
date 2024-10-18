package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTGenerator struct {
	secretKey string
}

func NewJWTGenerator() *JWTGenerator {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	return &JWTGenerator{secretKey: secretKey}
}

func (gen *JWTGenerator) GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(gen.secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func (gen *JWTGenerator) ValidateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(gen.secretKey), nil
	})
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}
	userId := claims["sub"].(string)
	return userId, nil
}


