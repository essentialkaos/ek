package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGenPassword() {
	weakPassword := GenPassword(16, STRENGTH_WEAK)
	medPassword := GenPassword(16, STRENGTH_MEDIUM)
	strongPassword := GenPassword(16, STRENGTH_STRONG)

	fmt.Printf("Weak password: %s\n", weakPassword)
	fmt.Printf("Medium password: %s\n", medPassword)
	fmt.Printf("Strong password: %s\n", strongPassword)
}

func ExampleGetPasswordStrength() {
	strength := GetPasswordStrength("secret1234%$")

	switch strength {
	case STRENGTH_STRONG:
		fmt.Println("Password is strong")
	case STRENGTH_MEDIUM:
		fmt.Println("Password is ok")
	case STRENGTH_WEAK:
		fmt.Println("Password is weak")
	}

	// Output:
	// Password is ok
}

func ExampleGenPasswordVariations() {
	password := "myPassword12345"
	variants := GenPasswordVariations(password)

	fmt.Printf("Variants: %v\n", variants)

	// Output:
	// Variants: [MYpASSWORD12345 MyPassword12345 myPassword1234]
}

func ExampleHash() {
	password, pepper := "MyPassword", "ABCD1234abcd1234"
	hash, err := Hash(password, pepper)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Hash: %s\n", hash)
}

func ExampleHashBytes() {
	password, pepper := []byte("MyPassword"), []byte("ABCD1234abcd1234")
	hash, err := HashBytes(password, pepper)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Hash: %s\n", hash)
}

func ExampleCheck() {
	password, pepper := "MyPassword", "ABCD1234abcd1234"
	hash, err := Hash(password, pepper)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Valid: %t\n", Check(password, pepper, hash))

	// Output:
	// Valid: true
}

func ExampleCheckBytes() {
	password, pepper := []byte("MyPassword"), []byte("ABCD1234abcd1234")
	hash, err := HashBytes(password, pepper)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Valid: %t\n", CheckBytes(password, pepper, hash))

	// Output:
	// Valid: true
}
