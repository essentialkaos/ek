package tmp

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

func ExampleNewTemp() {
	tmp, _ := NewTemp()

	fmt.Println(tmp.Dir)

	// Output: /tmp
}

func ExampleTemp_MkDir() {
	tmp, _ := NewTemp()

	fmt.Println(tmp.MkDir())
	fmt.Println(tmp.MkDir("test123"))

	tmp.Clean()
}

func ExampleTemp_MkFile() {
	tmp, _ := NewTemp()

	fmt.Println(tmp.MkFile())
	fmt.Println(tmp.MkFile("test123"))

	tmp.Clean()
}

func ExampleTemp_MkName() {
	tmp, _ := NewTemp()

	fmt.Println(tmp.MkDir())
	fmt.Println(tmp.MkDir("test123"))
}

func ExampleTemp_Clean() {
	tmp, _ := NewTemp()

	tmp.MkDir()

	// All temporary data will be removed
	tmp.Clean()
}
