package hash

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

func ExampleJCHash() {
	var (
		key     uint64 = 0xDEADBEEF
		buckets int    = 128 * 1024
	)

	hash := JCHash(key, buckets)

	fmt.Println(hash)

	// Output:
	// 110191
}

func ExampleFileHash() {
	hash := FileHash("/path/to/some/file")

	fmt.Println(hash)
}
