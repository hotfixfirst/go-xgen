package xgen

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash hashes the given password using HMAC-SHA256 and bcrypt.
// This is safe for storing in a password database.
func GeneratePasswordHash(secret, password string) (string, error) {
	preHashed, err := hashWithHMACSHA256(secret, password)
	if err != nil {
		return "", err
	}
	hashed, err := bcrypt.GenerateFromPassword(preHashed, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePasswordHash verifies whether the given password matches the bcrypt hash,
// using the same HMAC-SHA256 preprocessing.
func ComparePasswordHash(secret, password, hashed string) bool {
	preHashed, err := hashWithHMACSHA256(secret, password)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashed), preHashed)
	return err == nil
}

// hashWithHMACSHA256 computes HMAC-SHA256(password, secret).
// This is used as a pre-hash step before bcrypt.
func hashWithHMACSHA256(secret, password string) ([]byte, error) {
	if secret == "" || password == "" {
		return nil, errors.New("secret and password must not be empty")
	}
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(password))
	return mac.Sum(nil), nil
}
