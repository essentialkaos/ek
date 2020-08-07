package passwd

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2020 ESSENTIAL KAOS                          //
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
