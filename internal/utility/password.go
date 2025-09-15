package utility

import (
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

func HashPassword(plain string) (string, error) {
	if len([]byte(plain)) > 72 {
		return "", fmt.Errorf("password too long")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func VerifyPassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}