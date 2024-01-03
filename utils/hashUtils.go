package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password using bcrypt.
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}

// CheckPasswordHash compares a plain text password with a hashed password.
func CheckPasswordHash(password, hash string) bool {
	fmt.Println(password,hash)
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
