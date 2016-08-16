package tmp

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

func ExampleNewTemp() {
	tmp, err := tmp.NewTemp()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(tmp, err)

	// Output: &{/tmp []} <nil>
}

func Example_mkDir() {
	tmp, err := tmp.NewTemp()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(tmp.MkDir())
	fmt.Println(tmp.MkDir("test123"))

	tmp.Clean()

	// Output:
	// /tmp/_tmp_4xrgpNxaH6Gl <nil>
	// /tmp/_oDUNbUndLe2w_test123 <nil>
}

func Example_mkFile() {
	tmp, err := tmp.NewTemp()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(tmp.MkDir())
	fmt.Println(tmp.MkDir("test123"))

	tmp.Clean()

	// Output:
	// &{0xc8200a0d40} /tmp/_tmp_pfR9Qf6I5TZk <nil>
	// &{0xc8200a0e00} /tmp/_l9yKFblzvv4e_test123 <nil>
}

func Example_mkName() {
	tmp, err := tmp.NewTemp()

	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(tmp.MkDir())
	fmt.Println(tmp.MkDir("test123"))

	// Output:
	// /tmp/_tmp_4xrgpNxaH6Gl <nil>
	// /tmp/_oDUNbUndLe2w_test123 <nil>
}
