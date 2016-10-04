package hash

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
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
