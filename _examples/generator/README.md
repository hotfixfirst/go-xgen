# Generator Example

This example demonstrates the `xgen` generator functionality.

## Run

```bash
cd _examples/generator
go run main.go
```

## Features Demonstrated

| #   | Feature                | Function                                         |
| --- | ---------------------- | ------------------------------------------------ |
| 1   | UUID Generation        | `GenerateUUID()`, `GenerateUUIDWithoutDashes()`  |
| 2   | Microsecond ID         | `GenerateMicrosID()`                             |
| 3   | Nanosecond ID          | `GenerateNanosID()`                              |
| 4   | Random Base32 String   | `RandomBase32String()`                           |
| 5   | API Key                | `GenerateAPIKey()`                               |
| 6   | Secret Key             | `GenerateSecretKey()`                            |
| 7   | Sortability Demo       | Sequential ID generation                         |

## Sample Output

```text
=== Generator Examples ===

1. Generate UUID
----------------
   UUID (v7/v4): 0194d5a0-1234-7abc-8def-0123456789ab
   UUID without dashes: 0194d5a012347abc8def0123456789ab

2. Generate Microsecond ID (Sortable)
--------------------------------------
   MicrosID (no suffix):  0G3KQVH8J5T (len=11)
   MicrosID (5 suffix):   0G3KQVH8J5TABCDE (len=16)
   MicrosID (10 suffix):  0G3KQVH8J5TABCDEFGHIJ (len=21)

3. Generate Nanosecond ID (Higher Precision)
---------------------------------------------
   NanosID (no suffix):  0G3KQVH8J5TXY (len=13)
   NanosID (5 suffix):   0G3KQVH8J5TXYABCDE (len=18)
   NanosID (10 suffix):  0G3KQVH8J5TXYABCDEFGHIJ (len=23)

4. Random Base32 String (Crockford)
------------------------------------
   Random (10 chars): A1B2C3D4E5
   Random (20 chars): A1B2C3D4E5F6G7H8J9K0

5. Generate API Key
-------------------
   API Key (32 hex): a1b2c3d4e5f6789012345678abcdef01

6. Generate Secret Key
----------------------
   Secret Key (64 hex): a1b2c3d4e5f6789012345678abcdef01a1b2c3d4e5f6789012345678abcdef01

7. Sortability Demonstration
----------------------------
   Generating 5 MicrosIDs in sequence:
   [1] 0G3KQVH8J5TABCD
   [2] 0G3KQVH8J5TEFGH
   [3] 0G3KQVH8J5TIJKL
   [4] 0G3KQVH8J5TMNOP
   [5] 0G3KQVH8J5TQRST

=== End of Examples ===
```
