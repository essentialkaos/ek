package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleUUID4() {
	fmt.Printf("UUID v4: %s\n", UUID4().String())
}

func ExampleUUID5() {
	fmt.Printf("UUID v5: %s\n", UUID5(NsURL, "http://www.domain.com").String())
}

func ExampleUUID7() {
	fmt.Printf("UUID v7: %s\n", UUID7().String())
}
