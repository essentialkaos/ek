package uuid

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
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
