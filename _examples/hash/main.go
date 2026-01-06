// Package main demonstrates the usage of the xgen hash functionality.
package main

import (
	"fmt"

	"github.com/hotfixfirst/go-xgen"
)

func main() {
	fmt.Println("=== Hash Examples ===")
	fmt.Println()

	// Configuration
	secret := "my-super-secret-key"
	password := "user-password-123"

	// Example 1: Generate Password Hash
	fmt.Println("1. Generate Password Hash")
	fmt.Println("-------------------------")
	fmt.Printf("   Secret:   %s\n", secret)
	fmt.Printf("   Password: %s\n", password)

	hash, err := xgen.GeneratePasswordHash(secret, password)
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
		return
	}
	fmt.Printf("   Hash:     %s\n", hash)
	fmt.Println()

	// Example 2: Verify Correct Password
	fmt.Println("2. Verify Correct Password")
	fmt.Println("--------------------------")
	isValid := xgen.ComparePasswordHash(secret, password, hash)
	fmt.Printf("   Password: %s\n", password)
	fmt.Printf("   Valid:    %t ✓\n", isValid)
	fmt.Println()

	// Example 3: Verify Wrong Password
	fmt.Println("3. Verify Wrong Password")
	fmt.Println("------------------------")
	wrongPassword := "wrong-password"
	isValid = xgen.ComparePasswordHash(secret, wrongPassword, hash)
	fmt.Printf("   Password: %s\n", wrongPassword)
	fmt.Printf("   Valid:    %t ✗\n", isValid)
	fmt.Println()

	// Example 4: Verify Wrong Secret
	fmt.Println("4. Verify Wrong Secret")
	fmt.Println("----------------------")
	wrongSecret := "wrong-secret"
	isValid = xgen.ComparePasswordHash(wrongSecret, password, hash)
	fmt.Printf("   Secret:   %s\n", wrongSecret)
	fmt.Printf("   Password: %s\n", password)
	fmt.Printf("   Valid:    %t ✗\n", isValid)
	fmt.Println()

	// Example 5: Different Hashes for Same Password
	fmt.Println("5. Different Hashes for Same Password (bcrypt salt)")
	fmt.Println("----------------------------------------------------")
	hash1, _ := xgen.GeneratePasswordHash(secret, password)
	hash2, _ := xgen.GeneratePasswordHash(secret, password)
	fmt.Printf("   Hash 1: %s\n", hash1)
	fmt.Printf("   Hash 2: %s\n", hash2)
	fmt.Printf("   Same?:  %t (each hash is unique due to bcrypt salt)\n", hash1 == hash2)

	// Both should still validate
	valid1 := xgen.ComparePasswordHash(secret, password, hash1)
	valid2 := xgen.ComparePasswordHash(secret, password, hash2)
	fmt.Printf("   Both valid: %t && %t = %t ✓\n", valid1, valid2, valid1 && valid2)
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
