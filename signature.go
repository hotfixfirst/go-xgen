package xgen

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

// defaultSignatureDrift is the default allowed time drift for signature validation.
const defaultSignatureDrift = 5 * time.Minute

// BuildSignatureCanonicalString formats request components into a deterministic string for signing.
// Recommended structure: METHOD\nPATH\nTIMESTAMP\nBODY
func BuildSignatureCanonicalString(method, path, timestamp, rawBody string) string {
	return fmt.Sprintf("%s\n%s\n%s\n%s", method, path, timestamp, rawBody)
}

// GenerateSignature creates a HMAC-SHA256 signature in hex format using the given secret and canonical string.
func GenerateSignature(secret, method, path, timestamp, rawBody string) (string, error) {
	if secret == "" || method == "" || path == "" || timestamp == "" {
		return "", errors.New("missing required fields for signature")
	}
	canonical := BuildSignatureCanonicalString(method, path, timestamp, rawBody)
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(canonical))
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// VerifySignature compares the expected signature generated from request info with the received one.
// Uses constant-time comparison to prevent timing attacks.
func VerifySignature(secret, method, path, timestamp, rawBody, receivedSig string) bool {
	expectedSig, err := GenerateSignature(secret, method, path, timestamp, rawBody)
	if err != nil {
		return false
	}
	expected, err1 := hex.DecodeString(expectedSig)
	received, err2 := hex.DecodeString(receivedSig)
	if err1 != nil || err2 != nil {
		return false
	}
	return hmac.Equal(expected, received)
}

// IsValidSignatureTimestamp checks if the given timestamp (in string, unix seconds)
// is within the allowed TTL window, as specified by allowedDrift.
func IsValidSignatureTimestamp(timestamp string, allowedDrift time.Duration) bool {
	tsInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return false
	}
	ts := time.Unix(tsInt, 0)
	now := time.Now()
	diff := now.Sub(ts)
	if diff < 0 {
		diff = -diff
	}
	return diff <= allowedDrift
}

// IsValidSignatureTimestampDefault uses the default allowed TTL of ±5 minutes.
func IsValidSignatureTimestampDefault(timestamp string) bool {
	return IsValidSignatureTimestamp(timestamp, defaultSignatureDrift)
}
