package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleIsSafe() {
	path1 := "/home/user/project"
	path2 := "/usr/sbin/myapp"

	fmt.Printf("%s is safe → %t\n", path1, IsSafe(path1))
	fmt.Printf("%s is safe → %t\n", path2, IsSafe(path2))

	// Output:
	// /home/user/project is safe → true
	// /usr/sbin/myapp is safe → false
}

func ExampleDirN() {
	path1 := "/home/user/project/config/file.cfg"
	path2 := "/usr/sbin/myapp"

	fmt.Printf("Config dir: %s\n", DirN(path1, 4))
	fmt.Printf("Bin dir: %s\n", DirN(path2, 2))

	// Output:
	// Config dir: /home/user/project/config
	// Bin dir: /usr/sbin
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
