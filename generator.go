package xgen

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Base32Crockford is the alphabet used for Base32 Crockford encoding.
var Base32Crockford = []rune("0123456789ABCDEFGHJKMNPQRSTVWXYZ")

// GenerateUUID returns a UUID (v7 if possible, otherwise v4).
func GenerateUUID() uuid.UUID {
	if id, err := uuid.NewV7(); err == nil {
		return id
	}
	return uuid.New()
}

// GenerateUUIDWithoutDashes generates a UUID without dashes.
func GenerateUUIDWithoutDashes() string {
	return strings.ReplaceAll(GenerateUUID().String(), "-", "")
}

// GenerateMicrosID generates a timestamp-based ID with random suffix.
// Output: minimum 11 characters (prefix only), format: [11-char timestamp][suffix]
func GenerateMicrosID(suffixLength int) string {
	// Input validation
	if suffixLength < 0 {
		suffixLength = 0
	}
	prefix := encodeTimestampMicrosBase32(11)
	suffix := RandomBase32String(suffixLength)
	return prefix + suffix
}

// GenerateNanosID generates a timestamp-based ID with nanosecond precision and random suffix.
// Output: minimum 13 characters (prefix only), format: [13-char timestamp][suffix]
func GenerateNanosID(suffixLength int) string {
	// Input validation
	if suffixLength < 0 {
		suffixLength = 0
	}
	prefix := encodeTimestampNanosBase32(13)
	suffix := RandomBase32String(suffixLength)
	return prefix + suffix
}

// RandomBase32String generates a random Base32 string of specified length.
func RandomBase32String(n int) string {
	if n <= 0 {
		return ""
	}
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		// Fallback to deterministic random generation if crypto/rand fails
		fillBufferWithFallbackRandom(buf)
	}
	// Pre-allocate result slice with exact capacity
	result := make([]byte, n)
	for i := range n {
		result[i] = byte(Base32Crockford[int(buf[i])%32])
	}
	return string(result)
}

// encodeTimestampMicrosBase32 encodes current timestamp to Base32 with specified length.
func encodeTimestampMicrosBase32(prefixLength int) string {
	ts := timestampMicros()
	code := encodeBase32(ts)
	// Handle length efficiently
	if len(code) < prefixLength {
		// Pad left with zeros using strings.Repeat
		padding := strings.Repeat("0", prefixLength-len(code))
		return padding + code
	}
	// Keep only the last prefixLength chars
	if len(code) > prefixLength {
		return code[len(code)-prefixLength:]
	}
	return code
}

// encodeTimestampNanosBase32 encodes current timestamp in nanoseconds to Base32 with specified length.
func encodeTimestampNanosBase32(prefixLength int) string {
	ts := timestampNanos()
	code := encodeBase32(ts)
	// Handle length efficiently
	if len(code) < prefixLength {
		// Pad left with zeros using strings.Repeat
		padding := strings.Repeat("0", prefixLength-len(code))
		return padding + code
	}
	// Keep only the last prefixLength chars
	if len(code) > prefixLength {
		return code[len(code)-prefixLength:]
	}
	return code
}

// fillBufferWithFallbackRandom fills buffer with pseudo-random data using LCG algorithm.
// This is used as fallback when crypto/rand is not available.
func fillBufferWithFallbackRandom(buf []byte) {
	if len(buf) == 0 {
		return
	}
	// Use current time as seed for deterministic but unpredictable sequence
	seed := uint64(time.Now().UnixNano())
	// Same constants used by glibc: a=1103515245, c=12345
	for i := range buf {
		buf[i] = byte(seed % 256)
		seed = seed*1103515245 + 12345
	}
}

// encodeBase32 encodes a number into a Base32 string using the Crockford alphabet.
func encodeBase32(num uint64) string {
	if num == 0 {
		return "0"
	}
	// Calculate approximate length to reduce allocations
	// log32(num) ≈ log(num) / log(32) ≈ log(num) / 5
	length := 1
	temp := num
	for temp >= 32 {
		temp /= 32
		length++
	}
	// Use strings.Builder for efficient string building
	var builder strings.Builder
	builder.Grow(length)
	// Build string from right to left, then reverse
	digits := make([]byte, 0, length)
	for num > 0 {
		digits = append(digits, byte(Base32Crockford[num%32]))
		num /= 32
	}
	// Reverse the digits and write to builder
	for i := len(digits) - 1; i >= 0; i-- {
		builder.WriteByte(digits[i])
	}
	return builder.String()
}

// timestampMicros returns current timestamp in microseconds.
func timestampMicros() uint64 {
	return uint64(time.Now().UnixNano() / 1000)
}

// timestampNanos returns current timestamp in nanoseconds.
func timestampNanos() uint64 {
	return uint64(time.Now().UnixNano())
}

// Generate a random API key with prefix
func GenerateAPIKey() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// Generate a random secret key (32 bytes = 64 hex chars)
func GenerateSecretKey() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
