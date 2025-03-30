package prefixed

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"

	"github.com/essentialkaos/ek/v13/uuid"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleEncode() {
	pu := Encode("user", uuid.UUID7())

	fmt.Printf("UUID: %s\n", pu)
}

func ExampleDecode() {
	prefix, u, err := Decode("user.AZXm2EAle0im2Nloa94Wjg")

	fmt.Printf("Prefix: %s | UUID: %s | Error: %v\n", prefix, u, err)
}
