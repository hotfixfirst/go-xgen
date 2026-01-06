package xgen

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBuildSignatureCanonicalString(t *testing.T) {
	method := "POST"
	path := "/api/resource"
	timestamp := "1627849200"
	rawBody := "{\"key\":\"value\"}"

	canonical := BuildSignatureCanonicalString(method, path, timestamp, rawBody)
	expected := "POST\n/api/resource\n1627849200\n{\"key\":\"value\"}"
	assert.Equal(t, expected, canonical)
}

func TestGenerateSignature(t *testing.T) {
	secret := "test-secret"
	method := "POST"
	path := "/api/resource"
	timestamp := "1627849200"
	rawBody := "{\"key\":\"value\"}"

	signature, err := GenerateSignature(secret, method, path, timestamp, rawBody)
	assert.NoError(t, err)
	assert.NotEmpty(t, signature)
}

func TestGenerateSignature_MissingFields(t *testing.T) {
	_, err := GenerateSignature("", "POST", "/api/resource", "1627849200", "body")
	assert.Error(t, err)

	_, err = GenerateSignature("secret", "", "/api/resource", "1627849200", "body")
	assert.Error(t, err)

	_, err = GenerateSignature("secret", "POST", "", "1627849200", "body")
	assert.Error(t, err)

	_, err = GenerateSignature("secret", "POST", "/api/resource", "", "body")
	assert.Error(t, err)
}

func TestVerifySignature(t *testing.T) {
	secret := "test-secret"
	method := "POST"
	path := "/api/resource"
	timestamp := "1627849200"
	rawBody := "{\"key\":\"value\"}"

	signature, err := GenerateSignature(secret, method, path, timestamp, rawBody)
	assert.NoError(t, err)

	// Valid case
	isValid := VerifySignature(secret, method, path, timestamp, rawBody, signature)
	assert.True(t, isValid)

	// Invalid signature
	isValid = VerifySignature(secret, method, path, timestamp, rawBody, "invalid-signature")
	assert.False(t, isValid)
}

func TestIsValidSignatureTimestamp(t *testing.T) {
	now := time.Now()
	timestamp := now.Add(-2 * time.Minute).Unix()
	isValid := IsValidSignatureTimestamp(fmt.Sprintf("%d", timestamp), 5*time.Minute)
	assert.True(t, isValid)

	// Expired timestamp
	expiredTimestamp := now.Add(-10 * time.Minute).Unix()
	isValid = IsValidSignatureTimestamp(fmt.Sprintf("%d", expiredTimestamp), 5*time.Minute)
	assert.False(t, isValid)

	// Future timestamp
	futureTimestamp := now.Add(10 * time.Minute).Unix()
	isValid = IsValidSignatureTimestamp(fmt.Sprintf("%d", futureTimestamp), 5*time.Minute)
	assert.False(t, isValid)
}

func TestIsValidSignatureTimestampDefault(t *testing.T) {
	now := time.Now()
	timestamp := now.Add(-2 * time.Minute).Unix()
	isValid := IsValidSignatureTimestampDefault(fmt.Sprintf("%d", timestamp))
	assert.True(t, isValid)

	// Expired timestamp
	expiredTimestamp := now.Add(-10 * time.Minute).Unix()
	isValid = IsValidSignatureTimestampDefault(fmt.Sprintf("%d", expiredTimestamp))
	assert.False(t, isValid)
}
