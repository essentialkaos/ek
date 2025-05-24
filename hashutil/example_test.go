package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/sha1"
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleFile() {
	file := "/home/bob/data.txt"
	fmt.Printf("File %s hash is %q", file, File(file, sha1.New()))
}

func ExampleBytes() {
	data := []byte("Test Data\n")

	fmt.Printf("Data hash is %q\n", Bytes(data, sha1.New()))

	// Output:
	// Data hash is "58d3f35ebda104f0cc7ff419cc946832154581fc"
}

func ExampleString() {
	data := "Test Data\n"

	fmt.Printf("String hash is %q\n", String(data, sha1.New()))

	// Output:
	// String hash is "58d3f35ebda104f0cc7ff419cc946832154581fc"
}
