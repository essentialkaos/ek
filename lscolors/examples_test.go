package lscolors

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

func ExampleGetColor() {
	file := "myfile.txt"
	dir := "/home/john"

	colorSeq := GetColor(file)

	fmt.Printf(
		"%s/"+colorSeq+"%s"+GetColor(RESET)+"\n",
		dir, file,
	)
}

func ExampleColorize() {
	file := "myfile.txt"
	fmt.Println(Colorize(file))
}

func ExampleColorizePath() {
	path := "/home/john/myfile.txt"
	fmt.Println(ColorizePath(path))
}
