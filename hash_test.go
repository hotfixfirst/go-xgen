package xgen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePasswordHash(t *testing.T) {
	secret := "test-secret"
	password := "test-password"

	hash, err := GeneratePasswordHash(secret, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Verify the hash can be used with ComparePasswordHash
	isValid := ComparePasswordHash(secret, password, hash)
	assert.True(t, isValid)
}

func TestGeneratePasswordHash_EmptyInputs(t *testing.T) {
	_, err := GeneratePasswordHash("", "password")
	assert.Error(t, err)

	_, err = GeneratePasswordHash("secret", "")
	assert.Error(t, err)
}

func TestComparePasswordHash(t *testing.T) {
	secret := "test-secret"
	password := "test-password"

	hash, err := GeneratePasswordHash(secret, password)
	assert.NoError(t, err)

	// Valid case
	isValid := ComparePasswordHash(secret, password, hash)
	assert.True(t, isValid)

	// Invalid password
	isValid = ComparePasswordHash(secret, "wrong-password", hash)
	assert.False(t, isValid)

	// Invalid secret
	isValid = ComparePasswordHash("wrong-secret", password, hash)
	assert.False(t, isValid)

	// Invalid hash
	isValid = ComparePasswordHash(secret, password, "invalid-hash")
	assert.False(t, isValid)
}

func TestHashWithHMACSHA256(t *testing.T) {
	secret := "test-secret"
	password := "test-password"

	hash, err := hashWithHMACSHA256(secret, password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)

	// Empty inputs
	_, err = hashWithHMACSHA256("", password)
	assert.Error(t, err)

	_, err = hashWithHMACSHA256(secret, "")
	assert.Error(t, err)
}
