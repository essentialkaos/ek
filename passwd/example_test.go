package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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

func ExampleGenAuth() {
	auth := GenAuth(16, STRENGTH_MEDIUM)

	fmt.Printf("Password: %s\n", auth.Password)
	fmt.Printf("Salt: %s\n", auth.Salt)
	fmt.Printf("Hash: %s\n", auth.Hash)
}

func ExampleCreateAuth() {
	auth := CreateAuth("secret")

	fmt.Printf("Password: %s\n", auth.Password)
	fmt.Printf("Salt: %s\n", auth.Salt)
	fmt.Printf("Hash: %s\n", auth.Hash)
}

func ExampleGenHash() {
	hash := GenHash("secret", "1234abcd")

	fmt.Printf("Hash: %s\n", hash)

	// Output:
	// Hash: 679aca5f6295d9d5ff79fae69ad8150808b60728bf374de8857df5bc6d184f6d
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
