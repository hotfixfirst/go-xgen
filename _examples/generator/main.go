// Package main demonstrates the usage of the xgen generator functionality.
package main

import (
	"fmt"

	"github.com/hotfixfirst/go-xgen"
)

func main() {
	fmt.Println("=== Generator Examples ===")
	fmt.Println()

	// Example 1: Generate UUID
	fmt.Println("1. Generate UUID")
	fmt.Println("----------------")
	uuid := xgen.GenerateUUID()
	fmt.Printf("   UUID (v7/v4): %s\n", uuid.String())

	uuidNoDash := xgen.GenerateUUIDWithoutDashes()
	fmt.Printf("   UUID without dashes: %s\n", uuidNoDash)
	fmt.Println()

	// Example 2: Generate Microsecond-based ID
	fmt.Println("2. Generate Microsecond ID (Sortable)")
	fmt.Println("--------------------------------------")
	microsID0 := xgen.GenerateMicrosID(0)
	fmt.Printf("   MicrosID (no suffix):  %s (len=%d)\n", microsID0, len(microsID0))

	microsID5 := xgen.GenerateMicrosID(5)
	fmt.Printf("   MicrosID (5 suffix):   %s (len=%d)\n", microsID5, len(microsID5))

	microsID10 := xgen.GenerateMicrosID(10)
	fmt.Printf("   MicrosID (10 suffix):  %s (len=%d)\n", microsID10, len(microsID10))
	fmt.Println()

	// Example 3: Generate Nanosecond-based ID
	fmt.Println("3. Generate Nanosecond ID (Higher Precision)")
	fmt.Println("---------------------------------------------")
	nanosID0 := xgen.GenerateNanosID(0)
	fmt.Printf("   NanosID (no suffix):  %s (len=%d)\n", nanosID0, len(nanosID0))

	nanosID5 := xgen.GenerateNanosID(5)
	fmt.Printf("   NanosID (5 suffix):   %s (len=%d)\n", nanosID5, len(nanosID5))

	nanosID10 := xgen.GenerateNanosID(10)
	fmt.Printf("   NanosID (10 suffix):  %s (len=%d)\n", nanosID10, len(nanosID10))
	fmt.Println()

	// Example 4: Random Base32 String
	fmt.Println("4. Random Base32 String (Crockford)")
	fmt.Println("------------------------------------")
	rand10 := xgen.RandomBase32String(10)
	fmt.Printf("   Random (10 chars): %s\n", rand10)

	rand20 := xgen.RandomBase32String(20)
	fmt.Printf("   Random (20 chars): %s\n", rand20)
	fmt.Println()

	// Example 5: Generate API Key
	fmt.Println("5. Generate API Key")
	fmt.Println("-------------------")
	apiKey, err := xgen.GenerateAPIKey()
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   API Key (32 hex): %s\n", apiKey)
	}
	fmt.Println()

	// Example 6: Generate Secret Key
	fmt.Println("6. Generate Secret Key")
	fmt.Println("----------------------")
	secretKey, err := xgen.GenerateSecretKey()
	if err != nil {
		fmt.Printf("   Error: %v\n", err)
	} else {
		fmt.Printf("   Secret Key (64 hex): %s\n", secretKey)
	}
	fmt.Println()

	// Example 7: Sortability Demonstration
	fmt.Println("7. Sortability Demonstration")
	fmt.Println("----------------------------")
	fmt.Println("   Generating 5 MicrosIDs in sequence:")
	for i := 1; i <= 5; i++ {
		id := xgen.GenerateMicrosID(5)
		fmt.Printf("   [%d] %s\n", i, id)
	}
	fmt.Println()

	fmt.Println("=== End of Examples ===")
}
