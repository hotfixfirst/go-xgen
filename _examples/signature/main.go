// Package main demonstrates the usage of the xgen signature functionality.
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hotfixfirst/go-xgen"
)

func main() {
	fmt.Println("=== Signature Examples ===")
	fmt.Println()

	// Configuration
	secret := "my-api-secret-key"
	method := "POST"
	path := "/api/v1/users"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	rawBody := `{"name":"John Doe","email":"john@example.com"}`

	// Example 1: Build Canonical String
	fmt.Println("1. Build Canonical String")
	fmt.Println("-------------------------")
	canonical := xgen.BuildSignatureCanonicalString(method, path, timestamp, rawBody)
	fmt.Printf("   Method:    %s\n", method)
	fmt.Printf("   Path:      %s\n", path)
	fmt.Printf("   Timestamp: %s\n", timestamp)
	fmt.Printf("   Body:      %s\n", rawBody)
	fmt.Println("   Canonical String:")
	fmt.Printf("   ---\n%s\n   ---\n", canonical)
	fmt.Println()

	// Example 2: Generate Signature
	fmt.Println("2. Generate Signature")
	fmt.Println("---------------------")
	signature, err := xgen.GenerateSignature(secret, method, path, timestamp, rawBody)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Secret:    %s\n", secret)
	fmt.Printf("   Signature: %s\n", signature)
	fmt.Println()

	// Example 3: Verify Valid Signature
	fmt.Println("3. Verify Valid Signature")
	fmt.Println("-------------------------")
	isValid := xgen.VerifySignature(secret, method, path, timestamp, rawBody, signature)
	fmt.Printf("   Valid: %t ✓\n", isValid)
	fmt.Println()

	// Example 4: Verify Tampered Body
	fmt.Println("4. Verify Tampered Body")
	fmt.Println("-----------------------")
	tamperedBody := `{"name":"Jane Doe","email":"jane@example.com"}`
	isValid = xgen.VerifySignature(secret, method, path, timestamp, tamperedBody, signature)
	fmt.Printf("   Original Body: %s\n", rawBody)
	fmt.Printf("   Tampered Body: %s\n", tamperedBody)
	fmt.Printf("   Valid: %t ✗\n", isValid)
	fmt.Println()

	// Example 5: Verify Wrong Secret
	fmt.Println("5. Verify Wrong Secret")
	fmt.Println("----------------------")
	wrongSecret := "wrong-secret"
	isValid = xgen.VerifySignature(wrongSecret, method, path, timestamp, rawBody, signature)
	fmt.Printf("   Wrong Secret: %s\n", wrongSecret)
	fmt.Printf("   Valid: %t ✗\n", isValid)
	fmt.Println()

	// Example 6: Timestamp Validation (Valid)
	fmt.Println("6. Timestamp Validation (Current Time)")
	fmt.Println("---------------------------------------")
	currentTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	isValidTime := xgen.IsValidSignatureTimestampDefault(currentTimestamp)
	fmt.Printf("   Timestamp: %s (now)\n", currentTimestamp)
	fmt.Printf("   Valid (±5min): %t ✓\n", isValidTime)
	fmt.Println()

	// Example 7: Timestamp Validation (Expired)
	fmt.Println("7. Timestamp Validation (Expired)")
	fmt.Println("----------------------------------")
	oldTimestamp := strconv.FormatInt(time.Now().Add(-10*time.Minute).Unix(), 10)
	isValidTime = xgen.IsValidSignatureTimestampDefault(oldTimestamp)
	fmt.Printf("   Timestamp: %s (10 min ago)\n", oldTimestamp)
	fmt.Printf("   Valid (±5min): %t ✗\n", isValidTime)
	fmt.Println()

	// Example 8: Custom Drift
	fmt.Println("8. Timestamp Validation (Custom Drift)")
	fmt.Println("---------------------------------------")
	isValidTime = xgen.IsValidSignatureTimestamp(oldTimestamp, 15*time.Minute)
	fmt.Printf("   Timestamp: %s (10 min ago)\n", oldTimestamp)
	fmt.Printf("   Valid (±15min): %t ✓\n", isValidTime)
	fmt.Println()

	// Example 9: Full Request Signing Flow
	fmt.Println("9. Full Request Signing Flow")
	fmt.Println("-----------------------------")
	fmt.Println("   Client Side:")
	clientTimestamp := strconv.FormatInt(time.Now().Unix(), 10)
	clientSignature, _ := xgen.GenerateSignature(secret, "GET", "/api/v1/orders", clientTimestamp, "")
	fmt.Printf("   - Timestamp: %s\n", clientTimestamp)
	fmt.Printf("   - Signature: %s\n", clientSignature)

	fmt.Println("   Server Side:")
	// Verify timestamp first
	if !xgen.IsValidSignatureTimestampDefault(clientTimestamp) {
		fmt.Println("   - Timestamp expired! ✗")
	} else {
		fmt.Println("   - Timestamp valid ✓")
		// Then verify signature
		if xgen.VerifySignature(secret, "GET", "/api/v1/orders", clientTimestamp, "", clientSignature) {
			fmt.Println("   - Signature valid ✓")
			fmt.Println("   - Request authenticated! ✓")
		} else {
			fmt.Println("   - Signature invalid! ✗")
		}
	}
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
