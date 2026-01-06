# go-xgen

Go generator & security toolkit: unique IDs, random data, API keys, and HMAC signatures.

## Installation

```bash
go get github.com/hotfixfirst/go-xgen
```

## Quick Start

```go
import "github.com/hotfixfirst/go-xgen"

// Generate a sortable microsecond-based ID
id := xgen.GenerateMicrosID(10)
// id = "0G3KQVH8J5TABCDEFGHIJ"

// Hash a password securely
hash, err := xgen.GeneratePasswordHash("secret", "password")

// Generate and verify API signatures
sig, err := xgen.GenerateSignature("secret", "POST", "/api/users", "1704067200", `{"name":"John"}`)
valid := xgen.VerifySignature("secret", "POST", "/api/users", "1704067200", `{"name":"John"}`, sig)
```

## Features

| Feature                 | Description                              | Documentation                      |
| ----------------------- | ---------------------------------------- | ---------------------------------- |
| [Generator](#generator) | UUID, sortable IDs, API keys             | [Examples](./_examples/generator/) |
| [Hash](#hash)           | Password hashing (HMAC-SHA256 + bcrypt)  | [Examples](./_examples/hash/)      |
| [Signature](#signature) | Request signing & verification           | [Examples](./_examples/signature/) |

## Generator

Generate unique identifiers and random strings.

### Generator Functions

| Function                          | Description                                        |
| --------------------------------- | -------------------------------------------------- |
| `GenerateUUID()`                  | Generate UUID (v7 if possible, otherwise v4)       |
| `GenerateUUIDWithoutDashes()`     | Generate UUID without dashes (32 chars)            |
| `GenerateMicrosID(suffixLength)`  | Generate sortable microsecond-based ID (11+ chars) |
| `GenerateNanosID(suffixLength)`   | Generate sortable nanosecond-based ID (13+ chars)  |
| `RandomBase32String(length)`      | Generate random Base32 Crockford string            |
| `GenerateAPIKey()`                | Generate random API key (32 hex chars)             |
| `GenerateSecretKey()`             | Generate random secret key (64 hex chars)          |

### Generator Usage

```go
// UUID generation
uuid := xgen.GenerateUUID()           // 0194d5a0-1234-7abc-8def-0123456789ab
uuidNoDash := xgen.GenerateUUIDWithoutDashes() // 0194d5a012347abc8def0123456789ab

// Sortable IDs (great for database primary keys)
microsID := xgen.GenerateMicrosID(10) // 0G3KQVH8J5TABCDEFGHIJ (21 chars)
nanosID := xgen.GenerateNanosID(10)   // 0G3KQVH8J5TXYABCDEFGHIJ (23 chars)

// Random strings
random := xgen.RandomBase32String(20) // A1B2C3D4E5F6G7H8J9K0

// API credentials
apiKey, _ := xgen.GenerateAPIKey()    // a1b2c3d4e5f6789012345678abcdef01
secret, _ := xgen.GenerateSecretKey() // a1b2c3d4...abcdef01 (64 chars)
```

## Hash

Secure password hashing using HMAC-SHA256 + bcrypt.

### Hash Functions

| Function                                       | Description                                       |
| ---------------------------------------------- | ------------------------------------------------- |
| `GeneratePasswordHash(secret, password)`       | Hash password with HMAC-SHA256 pre-hash + bcrypt  |
| `ComparePasswordHash(secret, password, hash)`  | Verify password against hash                      |

### Hash Usage

```go
secret := "my-app-secret"
password := "user-password"

// Hash password (for storage)
hash, err := xgen.GeneratePasswordHash(secret, password)
// hash = "$2a$10$..."

// Verify password (on login)
valid := xgen.ComparePasswordHash(secret, password, hash)
// valid = true
```

## Signature

HMAC-SHA256 request signing for API authentication.

### Signature Functions

| Function                                                        | Description                                    |
| --------------------------------------------------------------- | ---------------------------------------------- |
| `BuildSignatureCanonicalString(method, path, timestamp, body)`  | Build canonical string for signing             |
| `GenerateSignature(secret, method, path, timestamp, body)`      | Generate HMAC-SHA256 signature                 |
| `VerifySignature(secret, method, path, timestamp, body, sig)`   | Verify signature (constant-time)               |
| `IsValidSignatureTimestamp(timestamp, drift)`                   | Check if timestamp is within allowed drift     |
| `IsValidSignatureTimestampDefault(timestamp)`                   | Check timestamp with ±5 minute drift           |

### Signature Usage

```go
secret := "my-api-secret"
method := "POST"
path := "/api/v1/users"
timestamp := strconv.FormatInt(time.Now().Unix(), 10)
body := `{"name":"John"}`

// Client: Generate signature
signature, err := xgen.GenerateSignature(secret, method, path, timestamp, body)

// Server: Verify timestamp and signature
if !xgen.IsValidSignatureTimestampDefault(timestamp) {
    // Reject: timestamp too old
}
if !xgen.VerifySignature(secret, method, path, timestamp, body, signature) {
    // Reject: invalid signature
}
// Request authenticated!
```

## Runnable Examples

See the [_examples](./_examples/) directory for runnable examples:

```bash
# Clone and run examples
git clone https://github.com/hotfixfirst/go-xgen.git
cd go-xgen/_examples

# Run generator examples
cd generator && go run main.go

# Run hash examples
cd ../hash && go run main.go

# Run signature examples
cd ../signature && go run main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

MIT License - see [LICENSE](LICENSE) file for details.
