package path

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

func ExampleIsSafe() {
	path1 := "/home/user/project"
	path2 := "/usr/sbin/project"

	fmt.Printf("%s is safe → %t\n", path1, IsSafe(path1))
	fmt.Printf("%s is safe → %t\n", path2, IsSafe(path2))

	// Output:
	// /home/user/project is safe → true
	// /usr/sbin/project is safe → false
}

func ExampleIsDotfile() {
	file1 := "/home/user/project/file"
	file2 := "/home/user/project/.file"

	fmt.Printf("%s is dotfile → %t\n", file1, IsDotfile(file1))
	fmt.Printf("%s is dotfile → %t\n", file2, IsDotfile(file2))

	// Output:
	// /home/user/project/file is dotfile → false
	// /home/user/project/.file is dotfile → true
}
