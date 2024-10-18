package utils

import (
	"reflect"
	"golang.org/x/crypto/bcrypt"
)

func DefaultValue[T interface{}](value T, defaultValue T) T {
    if reflect.ValueOf(value).IsZero() {
        return defaultValue
    }
    return value
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}