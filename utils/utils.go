package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// SetPassword sets password
func SetPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password should not be empty")
	}

	return Hash(password)
}

// Hash hashes password
func Hash(word string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(word), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePwd compare password
func ComparePwd(hashedPassword, onCheck string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(onCheck))
}
