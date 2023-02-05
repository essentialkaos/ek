package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleGenUUID() {
	fmt.Printf("UUID: %s\n", GenUUID())
}

func ExampleGenUUID4() {
	fmt.Printf("UUID v4: %s\n", GenUUID4())
}

func ExampleGenUUID5() {
	fmt.Printf("UUID v5: %s\n", GenUUID5(NsURL, "http://www.domain.com"))
}
