package kv

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

func ExampleSort() {
	data := []KV{
		{"test1", "1"},
		{"test5", "2"},
		{"test3", "3"},
		{"test2", "4"},
	}

	Sort(data)

	for _, kv := range data {
		fmt.Println(kv.Key, kv.Value)
	}

	// Output:
	// test1 1
	// test2 4
	// test3 3
	// test5 2
}
