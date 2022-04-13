package path

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleBase() {
	fmt.Println(Base("/home/user/project"))

	// Output:
	// project
}

func ExampleClean() {
	fmt.Println(Clean("/project//abc"))

	// Output:
	// /project/abc
}

func ExampleDir() {
	fmt.Println(Dir("/home/user/project"))

	// Output:
	// /home/user
}

func ExampleExt() {
	fmt.Println(Ext("/home/user/file.zip"))

	// Output:
	// .zip
}

func ExampleIsAbs() {
	fmt.Println(IsAbs("/dev/null"))

	// Output:
	// true
}

func ExampleJoin() {
	fmt.Println(Join("home", "user", "project"))

	// Output:
	// home/user/project
}

func ExampleMatch() {
	fmt.Println(Match("/home/*", "/home/user"))

	// Output:
	// true <nil>
}

func ExampleSplit() {
	fmt.Println(Split("/home/user/john/file.zip"))

	// Output:
	// /home/user/john/ file.zip
}

func ExampleCompact() {
	fmt.Println(Compact("/home/user/john/file.zip"))

	// Output:
	// /h/u/j/file.zip
}

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

func ExampleIsGlob() {
	file1 := "file"
	file2 := "*.file"

	fmt.Printf("%s is glob → %t\n", file1, IsGlob(file1))
	fmt.Printf("%s is glob → %t\n", file2, IsGlob(file2))

	// Output:
	// file is glob → false
	// *.file is glob → true
}
