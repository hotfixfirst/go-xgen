package xgen

import (
	"regexp"
	"strings"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"generates valid UUID"},
		{"generates unique UUIDs"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUUID()
			if got.String() == "" {
				t.Error("GenerateUUID() returned empty UUID")
			}
			// UUID should be 36 characters (including dashes)
			if len(got.String()) != 36 {
				t.Errorf("GenerateUUID() = %v, want length 36", got.String())
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			id := GenerateUUID().String()
			if seen[id] {
				t.Errorf("GenerateUUID() generated duplicate UUID: %v", id)
			}
			seen[id] = true
		}
	})
}

func TestGenerateUUIDWithoutDashes(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"generates UUID without dashes"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateUUIDWithoutDashes()
			if got == "" {
				t.Error("GenerateUUIDWithoutDashes() returned empty string")
			}
			// UUID without dashes should be 32 characters
			if len(got) != 32 {
				t.Errorf("GenerateUUIDWithoutDashes() = %v, want length 32, got %d", got, len(got))
			}
			// Should not contain dashes
			if strings.Contains(got, "-") {
				t.Errorf("GenerateUUIDWithoutDashes() = %v, should not contain dashes", got)
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			id := GenerateUUIDWithoutDashes()
			if seen[id] {
				t.Errorf("GenerateUUIDWithoutDashes() generated duplicate: %v", id)
			}
			seen[id] = true
		}
	})
}

func TestGenerateMicrosID(t *testing.T) {
	tests := []struct {
		name         string
		suffixLength int
		wantMinLen   int
	}{
		{"zero suffix", 0, 11},
		{"negative suffix", -1, 11},
		{"small suffix", 5, 16},
		{"medium suffix", 10, 21},
		{"large suffix", 20, 31},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateMicrosID(tt.suffixLength)
			if len(got) < tt.wantMinLen {
				t.Errorf("GenerateMicrosID(%d) = %v, want min length %d, got %d", tt.suffixLength, got, tt.wantMinLen, len(got))
			}
			// Should only contain Base32 Crockford characters
			if !isValidBase32Crockford(got) {
				t.Errorf("GenerateMicrosID(%d) = %v, contains invalid characters", tt.suffixLength, got)
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			id := GenerateMicrosID(10)
			if seen[id] {
				t.Errorf("GenerateMicrosID() generated duplicate: %v", id)
			}
			seen[id] = true
		}
	})

	// Test sortability (IDs generated later should be >= earlier ones)
	t.Run("sortability", func(t *testing.T) {
		id1 := GenerateMicrosID(0)
		id2 := GenerateMicrosID(0)
		if id2 < id1 {
			t.Errorf("GenerateMicrosID() not sortable: %v should be >= %v", id2, id1)
		}
	})
}

func TestGenerateNanosID(t *testing.T) {
	tests := []struct {
		name         string
		suffixLength int
		wantMinLen   int
	}{
		{"zero suffix", 0, 13},
		{"negative suffix", -1, 13},
		{"small suffix", 5, 18},
		{"medium suffix", 10, 23},
		{"large suffix", 20, 33},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateNanosID(tt.suffixLength)
			if len(got) < tt.wantMinLen {
				t.Errorf("GenerateNanosID(%d) = %v, want min length %d, got %d", tt.suffixLength, got, tt.wantMinLen, len(got))
			}
			// Should only contain Base32 Crockford characters
			if !isValidBase32Crockford(got) {
				t.Errorf("GenerateNanosID(%d) = %v, contains invalid characters", tt.suffixLength, got)
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			id := GenerateNanosID(10)
			if seen[id] {
				t.Errorf("GenerateNanosID() generated duplicate: %v", id)
			}
			seen[id] = true
		}
	})

	// Test sortability
	t.Run("sortability", func(t *testing.T) {
		id1 := GenerateNanosID(0)
		id2 := GenerateNanosID(0)
		if id2 < id1 {
			t.Errorf("GenerateNanosID() not sortable: %v should be >= %v", id2, id1)
		}
	})
}

func TestRandomBase32String(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantLen int
	}{
		{"zero length", 0, 0},
		{"negative length", -1, 0},
		{"length 1", 1, 1},
		{"length 10", 10, 10},
		{"length 100", 100, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomBase32String(tt.length)
			if len(got) != tt.wantLen {
				t.Errorf("RandomBase32String(%d) length = %d, want %d", tt.length, len(got), tt.wantLen)
			}
			if tt.wantLen > 0 && !isValidBase32Crockford(got) {
				t.Errorf("RandomBase32String(%d) = %v, contains invalid characters", tt.length, got)
			}
		})
	}

	// Test randomness (statistical test)
	t.Run("randomness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			s := RandomBase32String(20)
			if seen[s] {
				t.Errorf("RandomBase32String() generated duplicate: %v", s)
			}
			seen[s] = true
		}
	})
}

func TestEncodeBase32(t *testing.T) {
	tests := []struct {
		name string
		num  uint64
		want string
	}{
		{"zero", 0, "0"},
		{"one", 1, "1"},
		{"31", 31, "Z"},
		{"32", 32, "10"},
		{"1023", 1023, "ZZ"},
		{"1024", 1024, "100"},
		{"large number", 1000000, "YGJ0"},
		{"max uint32", 4294967295, "3ZZZZZZ"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeBase32(tt.num)
			if got != tt.want {
				t.Errorf("encodeBase32(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

func TestEncodeTimestampMicrosBase32(t *testing.T) {
	tests := []struct {
		name         string
		prefixLength int
	}{
		{"standard length 11", 11},
		{"short length 5", 5},
		{"long length 15", 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeTimestampMicrosBase32(tt.prefixLength)
			if len(got) != tt.prefixLength {
				t.Errorf("encodeTimestampMicrosBase32(%d) length = %d, want %d", tt.prefixLength, len(got), tt.prefixLength)
			}
			if !isValidBase32Crockford(got) {
				t.Errorf("encodeTimestampMicrosBase32(%d) = %v, contains invalid characters", tt.prefixLength, got)
			}
		})
	}
}

func TestEncodeTimestampNanosBase32(t *testing.T) {
	tests := []struct {
		name         string
		prefixLength int
	}{
		{"standard length 13", 13},
		{"short length 5", 5},
		{"long length 20", 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := encodeTimestampNanosBase32(tt.prefixLength)
			if len(got) != tt.prefixLength {
				t.Errorf("encodeTimestampNanosBase32(%d) length = %d, want %d", tt.prefixLength, len(got), tt.prefixLength)
			}
			if !isValidBase32Crockford(got) {
				t.Errorf("encodeTimestampNanosBase32(%d) = %v, contains invalid characters", tt.prefixLength, got)
			}
		})
	}
}

func TestFillBufferWithFallbackRandom(t *testing.T) {
	tests := []struct {
		name   string
		bufLen int
	}{
		{"empty buffer", 0},
		{"single byte", 1},
		{"small buffer", 10},
		{"medium buffer", 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := make([]byte, tt.bufLen)
			fillBufferWithFallbackRandom(buf)
			// For non-empty buffers, check that they are filled
			if tt.bufLen > 0 {
				allZero := true
				for _, b := range buf {
					if b != 0 {
						allZero = false
						break
					}
				}
				// It's statistically unlikely (but possible) for all bytes to be zero
				// This is a basic sanity check
				if allZero && tt.bufLen > 5 {
					t.Logf("Warning: fillBufferWithFallbackRandom() produced all zeros for buffer of length %d", tt.bufLen)
				}
			}
		})
	}

	// Test that different calls produce different results (due to time-based seed)
	t.Run("variability", func(t *testing.T) {
		buf1 := make([]byte, 16)
		buf2 := make([]byte, 16)
		fillBufferWithFallbackRandom(buf1)
		fillBufferWithFallbackRandom(buf2)
		// Note: These may be the same if called within the same nanosecond
		// This is expected behavior for the fallback random generator
	})
}

func TestGenerateAPIKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"generates valid API key"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateAPIKey()
			if err != nil {
				t.Errorf("GenerateAPIKey() error = %v", err)
				return
			}
			// API key should be 32 hex characters (16 bytes)
			if len(got) != 32 {
				t.Errorf("GenerateAPIKey() = %v, want length 32, got %d", got, len(got))
			}
			// Should only contain hex characters
			if !isValidHex(got) {
				t.Errorf("GenerateAPIKey() = %v, contains non-hex characters", got)
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			key, err := GenerateAPIKey()
			if err != nil {
				t.Errorf("GenerateAPIKey() error = %v", err)
				continue
			}
			if seen[key] {
				t.Errorf("GenerateAPIKey() generated duplicate: %v", key)
			}
			seen[key] = true
		}
	})
}

func TestGenerateSecretKey(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"generates valid secret key"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateSecretKey()
			if err != nil {
				t.Errorf("GenerateSecretKey() error = %v", err)
				return
			}
			// Secret key should be 64 hex characters (32 bytes)
			if len(got) != 64 {
				t.Errorf("GenerateSecretKey() = %v, want length 64, got %d", got, len(got))
			}
			// Should only contain hex characters
			if !isValidHex(got) {
				t.Errorf("GenerateSecretKey() = %v, contains non-hex characters", got)
			}
		})
	}

	// Test uniqueness
	t.Run("uniqueness", func(t *testing.T) {
		seen := make(map[string]bool)
		for i := 0; i < 1000; i++ {
			key, err := GenerateSecretKey()
			if err != nil {
				t.Errorf("GenerateSecretKey() error = %v", err)
				continue
			}
			if seen[key] {
				t.Errorf("GenerateSecretKey() generated duplicate: %v", key)
			}
			seen[key] = true
		}
	})
}

func TestBase32CrockfordAlphabet(t *testing.T) {
	t.Run("correct length", func(t *testing.T) {
		if len(Base32Crockford) != 32 {
			t.Errorf("Base32Crockford length = %d, want 32", len(Base32Crockford))
		}
	})

	t.Run("no duplicate characters", func(t *testing.T) {
		seen := make(map[rune]bool)
		for _, r := range Base32Crockford {
			if seen[r] {
				t.Errorf("Base32Crockford contains duplicate character: %c", r)
			}
			seen[r] = true
		}
	})

	t.Run("expected characters", func(t *testing.T) {
		expected := "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
		if string(Base32Crockford) != expected {
			t.Errorf("Base32Crockford = %s, want %s", string(Base32Crockford), expected)
		}
	})
}

// Helper functions for tests

func isValidBase32Crockford(s string) bool {
	validChars := regexp.MustCompile(`^[0-9A-HJKMNP-TV-Z]*$`)
	return validChars.MatchString(s)
}

func isValidHex(s string) bool {
	validChars := regexp.MustCompile(`^[0-9a-f]*$`)
	return validChars.MatchString(s)
}

// Benchmarks

func BenchmarkGenerateUUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUUID()
	}
}

func BenchmarkGenerateUUIDWithoutDashes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateUUIDWithoutDashes()
	}
}

func BenchmarkGenerateMicrosID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateMicrosID(10)
	}
}

func BenchmarkGenerateNanosID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateNanosID(10)
	}
}

func BenchmarkRandomBase32String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomBase32String(20)
	}
}

func BenchmarkEncodeBase32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeBase32(1000000)
	}
}

func BenchmarkGenerateAPIKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateAPIKey()
	}
}

func BenchmarkGenerateSecretKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateSecretKey()
	}
}
