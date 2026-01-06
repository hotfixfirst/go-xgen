# Signature Example

This example demonstrates the `xgen` HMAC-SHA256 signature functionality.

## Run

```bash
cd _examples/signature
go run main.go
```

## Features Demonstrated

| # | Feature | Function |
| - | ------- | -------- |
| 1 | Build Canonical String | `BuildSignatureCanonicalString()` |
| 2 | Generate Signature | `GenerateSignature()` |
| 3 | Verify Valid Signature | `VerifySignature()` |
| 4 | Detect Tampered Body | `VerifySignature()` |
| 5 | Detect Wrong Secret | `VerifySignature()` |
| 6 | Timestamp Validation (Current) | `IsValidSignatureTimestampDefault()` |
| 7 | Timestamp Validation (Expired) | `IsValidSignatureTimestampDefault()` |
| 8 | Custom Time Drift | `IsValidSignatureTimestamp()` |
| 9 | Full Request Flow | Complete signing workflow |

## How It Works

### Canonical String Format

```text
METHOD\nPATH\nTIMESTAMP\nBODY
```

Example:

```text
POST
/api/v1/users
1704067200
{"name":"John Doe"}
```

### Signing Process

1. Build canonical string from request components
2. Compute HMAC-SHA256 with secret key
3. Return hex-encoded signature

### Verification Process

1. Check timestamp is within allowed drift (default ±5 minutes)
2. Rebuild canonical string from request
3. Compute expected signature
4. Compare with received signature (constant-time)

## Sample Output

```text
=== Signature Examples ===

1. Build Canonical String
-------------------------
   Method:    POST
   Path:      /api/v1/users
   Timestamp: 1704067200
   Body:      {"name":"John Doe","email":"john@example.com"}
   Canonical String:
   ---
POST
/api/v1/users
1704067200
{"name":"John Doe","email":"john@example.com"}
   ---

2. Generate Signature
---------------------
   Secret:    my-api-secret-key
   Signature: a1b2c3d4e5f6...

3. Verify Valid Signature
-------------------------
   Valid: true ✓

4. Verify Tampered Body
-----------------------
   Original Body: {"name":"John Doe","email":"john@example.com"}
   Tampered Body: {"name":"Jane Doe","email":"jane@example.com"}
   Valid: false ✗

...

9. Full Request Signing Flow
-----------------------------
   Client Side:
   - Timestamp: 1704067200
   - Signature: abc123...
   Server Side:
   - Timestamp valid ✓
   - Signature valid ✓
   - Request authenticated! ✓

=== End of Examples ===
```
