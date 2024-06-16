// Package utils provides common utility functions for the application.
package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Encrypt takes a plain text string and returns an encrypted hash string
// and any error that occurred during the encryption process.
func Encrypt(text string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	return string(hash), err
}

// CompareHashAndPassword compares a hash string and a password string,
// and returns true if the password matches the hash, false otherwise.
func CompareHashAndPassword(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
