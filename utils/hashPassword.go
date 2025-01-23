package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword untuk melakukan hashing password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
