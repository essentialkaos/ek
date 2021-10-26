package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleCheckPerms() {
	dir := "/home/john"
	file := dir + "/test.txt"

	// Target is file, readable and non-empty
	if CheckPerms("FRS", file) {
		fmt.Println("Everything fine!")
	}

	// Target is readable, writable and executable
	if CheckPerms("RWX", file) {
		fmt.Println("Everything fine!")
	}

	// Target is directory, readable, writable and executable
	if CheckPerms("DRWX", file) {
		fmt.Println("Everything fine!")
	}
}

func ExampleValidatePerms() {
	dir := "/home/john"
	file := dir + "/test.txt"

	// Target is file, readable and non-empty
	err := ValidatePerms("FRS", file)

	if err != nil {
		fmt.Println(err.Error())
	}
}

func ExampleProperPath() {
	paths := []string{
		"/home/john/.config/myapp/config",
		"/home/john/.myappconfig",
		"/etc/myapp.conf",
	}

	config := ProperPath("FRS", paths)

	if config != "" {
		fmt.Printf("Used configuration file: %s\n", config)
	} else {
		fmt.Println("Can't find configuration file")
	}
}

func ExampleIsExist() {
	file := "/home/john/test.txt"

	if IsExist(file) {
		fmt.Printf("File %s is found on system!\n", file)
	} else {
		fmt.Printf("File %s does not exist!\n", file)
	}
}

func ExampleIsRegular() {
	target := "/home/john/test.txt"

	if IsRegular(target) {
		fmt.Printf("%s is a regular file!\n", target)
	} else {
		fmt.Printf("%s is NOT a regular file!\n", target)
	}
}

func ExampleIsSocket() {
	target := "/var/run/myapp.sock"

	if IsSocket(target) {
		fmt.Printf("%s is a socket file!\n", target)
	} else {
		fmt.Printf("%s is NOT a socket file!\n", target)
	}
}

func ExampleIsBlockDevice() {
	target := "/dev/sda"

	if IsBlockDevice(target) {
		fmt.Printf("%s is a block device!\n", target)
	} else {
		fmt.Printf("%s is NOT a block device!\n", target)
	}
}

func ExampleIsCharacterDevice() {
	target := "/dev/tty0"

	if IsCharacterDevice(target) {
		fmt.Printf("%s is a character device!\n", target)
	} else {
		fmt.Printf("%s is NOT a character device!\n", target)
	}
}

func ExampleIsDir() {
	target := "/home/john"

	if IsDir(target) {
		fmt.Printf("%s is a directory!\n", target)
	} else {
		fmt.Printf("%s is NOT a directory!\n", target)
	}
}

func ExampleIsLink() {
	target := "/dev/stdout"

	if IsLink(target) {
		fmt.Printf("%s is a link!\n", target)
	} else {
		fmt.Printf("%s is NOT a link!\n", target)
	}
}

func ExampleIsReadable() {
	target := "/home/john/test.txt"

	if IsReadable(target) {
		fmt.Printf("%s is readable!\n", target)
	} else {
		fmt.Printf("%s is NOT readable!\n", target)
	}
}

func ExampleIsReadableByUser() {
	target := "/home/john/test.txt"
	user := "johndoe"

	if IsReadableByUser(target, user) {
		fmt.Printf("%s is readable for user %s!\n", target, user)
	} else {
		fmt.Printf("%s is NOT readable for user %s!\n", target, user)
	}
}

func ExampleIsWritable() {
	target := "/home/john/test.txt"

	if IsWritable(target) {
		fmt.Printf("%s is writable!\n", target)
	} else {
		fmt.Printf("%s is NOT writable!\n", target)
	}
}

func ExampleIsWritableByUser() {
	target := "/home/john/test.txt"
	user := "johndoe"

	if IsWritableByUser(target, user) {
		fmt.Printf("%s is writable for user %s!\n", target, user)
	} else {
		fmt.Printf("%s is NOT writable for user %s!\n", target, user)
	}
}

func ExampleIsExecutable() {
	target := "/home/john/myapp"

	if IsExecutable(target) {
		fmt.Printf("%s is executable!\n", target)
	} else {
		fmt.Printf("%s is NOT executable!\n", target)
	}
}

func ExampleIsExecutableByUser() {
	target := "/home/john/myapp"
	user := "johndoe"

	if IsExecutableByUser(target, user) {
		fmt.Printf("%s is executable for user %s!\n", target, user)
	} else {
		fmt.Printf("%s is NOT executable for user %s!\n", target, user)
	}
}

func ExampleIsEmpty() {
	target := "/home/john/test.txt"

	if IsEmpty(target) {
		fmt.Printf("%s is an empty file!\n", target)
	} else {
		fmt.Printf("%s is NOT an empty file!\n", target)
	}
}

func ExampleIsNonEmpty() {
	target := "/home/john/test.txt"

	if IsNonEmpty(target) {
		fmt.Printf("%s is NOT an empty file!\n", target)
	} else {
		fmt.Printf("%s is an empty file!\n", target)
	}
}

func ExampleIsEmptyDir() {
	target := "/home/john/myfiles"

	if IsEmptyDir(target) {
		fmt.Printf("%s is an empty directory!\n", target)
	} else {
		fmt.Printf("%s is NOT an empty directory!\n", target)
	}
}

func ExampleGetOwner() {
	target := "/home/john/test.txt"

	uid, gid, err := GetOwner(target)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Owner UID (User ID): %d\n", uid)
	fmt.Printf("Owner GID (Group ID): %d\n", gid)
}

func ExampleGetATime() {
	target := "/home/john/test.txt"
	aTime, err := GetATime(target)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Access timestamp (atime): %v\n", aTime)
}

func ExampleGetCTime() {
	target := "/home/john/test.txt"
	aTime, err := GetCTime(target)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Change timestamp (ctime): %v\n", aTime)
}

func ExampleGetMTime() {
	target := "/home/john/test.txt"
	aTime, err := GetMTime(target)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Modified timestamp (mtime): %v\n", aTime)
}

func ExampleGetSize() {
	target := "/home/john/test.txt"
	size := GetSize(target)

	if size != -1 {
		fmt.Printf("File size: %d bytes\n", size)
	}
}

func ExampleGetMode() {
	target := "/home/john/test.txt"
	mode := GetMode(target)

	if mode != 0 {
		fmt.Printf("File mode: %v\n", mode)
	}
}

func ExampleCopyFile() {
	target := "/home/john/test.txt"

	err := CopyFile(target, "/home/bob/test.txt", 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File %s successfully copied to bob user directory\n", target)
}

func ExampleMoveFile() {
	target := "/home/john/test.txt"

	err := MoveFile(target, "/home/bob/test.txt", 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File %s successfully moved to bob user directory\n", target)
}

func ExampleCopyDir() {
	target := "/home/john/documents"

	err := CopyDir(target, "/home/bob/")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Directory %s successfully copied to bob user directory\n", target)
}

func ExampleTouchFile() {
	err := TouchFile("/srv/myapp/.lock", 0600)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Lock file successfully created!")
}
