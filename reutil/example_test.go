package reutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"regexp"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleReplace() {
	re := regexp.MustCompile(`([a-zA-Z0-9_.]+)@([a-zA-Z0-9-_.]+\.[a-z]{2,3})`)
	mails := "jdoe@yahoo.com bob.entus@gmail.com dina@mail.org"
	mappings := map[string]string{
		"yahoo.com": "YH",
		"gmail.com": "GML",
		"mail.org":  "ML",
	}

	repl, err := Replace(re, mails, func(found string, submatch []string) string {
		return submatch[0] + ":" + mappings[submatch[1]]
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%s → %s\n", mails, repl)

	// Output:
	// jdoe@yahoo.com bob.entus@gmail.com dina@mail.org → jdoe:YH bob.entus:GML dina:ML
}
