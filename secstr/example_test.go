package secstr

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

func ExampleNewSecureString() {
	passwd := "MySuppaPassword12345"
	ss, err := NewSecureString(&passwd)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Password: %s", string(ss.Data))
}
