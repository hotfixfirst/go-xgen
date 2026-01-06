# Hash Example

This example demonstrates the `xgen` password hashing functionality.

## Run

```bash
cd _examples/hash
go run main.go
```

## Features Demonstrated

| # | Feature | Function |
| - | ------- | -------- |
| 1 | Password Hash Generation | `GeneratePasswordHash()` |
| 2 | Verify Correct Password | `ComparePasswordHash()` |
| 3 | Verify Wrong Password | `ComparePasswordHash()` |
| 4 | Verify Wrong Secret | `ComparePasswordHash()` |
| 5 | Unique Hashes (bcrypt salt) | Multiple `GeneratePasswordHash()` |

## How It Works

The password hashing uses a two-step process:

1. **HMAC-SHA256**: Pre-hash the password with the secret key
2. **bcrypt**: Hash the result using bcrypt with default cost

This provides:

- **Secret binding**: Passwords are tied to your application secret
- **Slow hashing**: bcrypt is intentionally slow to prevent brute-force
- **Unique salts**: Each hash is unique even for the same password

## Sample Output

```text
=== Hash Examples ===

1. Generate Password Hash
-------------------------
   Secret:   my-super-secret-key
   Password: user-password-123
   Hash:     $2a$10$abc123...xyz789

2. Verify Correct Password
--------------------------
   Password: user-password-123
   Valid:    true ✓

3. Verify Wrong Password
------------------------
   Password: wrong-password
   Valid:    false ✗

4. Verify Wrong Secret
----------------------
   Secret:   wrong-secret
   Password: user-password-123
   Valid:    false ✗

5. Different Hashes for Same Password (bcrypt salt)
----------------------------------------------------
   Hash 1: $2a$10$abc123...
   Hash 2: $2a$10$def456...
   Same?:  false (each hash is unique due to bcrypt salt)
   Both valid: true && true = true ✓

=== End of Examples ===
```
