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

func ExampleParse() {
	uuid, err := Parse(`32186397b16440ad876bec223568b2d1`)
	fmt.Println(uuid, err)

	// Output:
	// 32186397-b164-40ad-876b-ec223568b2d1 <nil>
}

func ExampleIsValid() {
	uuids := []string{
		"32186397-b164-40ad-876b-ec223568b2d1",
		"32186397b16440ad876bec223568b2d1",
		"abcd1234ABCD1234TEST",
	}

	for _, uuid := range uuids {
		fmt.Printf("%s → %t\n", uuid, IsValid(uuid))
	}

	// Output:
	// 32186397-b164-40ad-876b-ec223568b2d1 → true
	// 32186397b16440ad876bec223568b2d1 → true
	// abcd1234ABCD1234TEST → false
}

func ExampleUUID4() {
	fmt.Printf("UUID v4: %s\n", UUID4().String())
}

func ExampleUUID5() {
	fmt.Printf("UUID v5: %s\n", UUID5(NsURL, "http://www.domain.com").String())
}

func ExampleUUID7() {
	fmt.Printf("UUID v7: %s\n", UUID7().String())
}
