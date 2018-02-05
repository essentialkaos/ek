package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2018 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
