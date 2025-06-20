package helper

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword mengenkripsi password plain text menjadi hashed password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // DefaultCost = 10
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// CheckPassword membandingkan password input dengan password yang sudah di-hash
func CheckPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("password is incorrect")
	}
	return nil
}
