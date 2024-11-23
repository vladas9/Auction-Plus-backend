package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Faled to hash password: %v", err.Error())
	}
	return string(hashedPassword), nil
}

func CompareHashPassword(password string, storedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password)); err != nil {
		return fmt.Errorf("Invalid password: %v", err.Error())
	}
	return nil
}
