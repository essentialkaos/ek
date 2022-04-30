package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"time"
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

	target = "/home/john/test.txt"

	// Permissions and target file name are optional and can be omitted
	err = CopyFile(target, "/home/bob")

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File %s successfully copied to bob user directory\n", target)
}

func ExampleCopyAttr() {
	source := "/home/john/test1.txt"
	target := "/home/john/test2.txt"

	err := CopyAttr(source, target)

	if err != nil {
		panic(err.Error())
	}

	fmt.Print("File attributes successfully copied")
}

func ExampleMoveFile() {
	target := "/home/john/test.txt"
	err := MoveFile(target, "/home/bob/test.txt", 0644)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File %s successfully moved to bob user directory\n", target)

	target = "/home/john/test.txt"

	// Permissions and target file name are optional and can be omitted
	err = MoveFile(target, "/home/bob/")

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

func ExampleCountLines() {
	file := "/home/john/test.txt"

	lineNum, err := CountLines(file)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("File %s contains %d lines of text\n", file, lineNum)
}

func ExampleList() {
	dir := "/home/john/documents"

	// List all objects including hidden (files and directries
	// with the dot at the beginning of the file name)
	objects := List(dir, false)

	if len(objects) == 0 {
		fmt.Printf("Directory %s is empty", dir)
		return
	}

	fmt.Printf("Directory %s contains:\n", dir)

	for _, object := range objects {
		fmt.Printf("  %s\n", object)
	}
}

func ExampleListAll() {
	dir := "/home/john/documents"

	// List all objects excluding hidden (files and directries
	// with the dot at the beginning of the file name)
	objects := ListAll(dir, true)

	if len(objects) == 0 {
		fmt.Printf("Directory %s is empty", dir)
		return
	}

	fmt.Printf("Directory %s contains:\n", dir)

	for _, object := range objects {
		fmt.Printf("  %s\n", object)
	}
}

func ExampleListAllDirs() {
	target := "/home/john/documents"

	// List all directories including hidden (directries
	// with the dot at the beginning of the file name)
	dirs := ListAllDirs(target, true)

	if len(dirs) == 0 {
		fmt.Printf("Directory %s is empty", target)
		return
	}

	fmt.Printf("Directory %s contains:\n", target)

	for _, dir := range dirs {
		fmt.Printf("  %s\n", dir)
	}
}

func ExampleListAllFiles() {
	target := "/home/john/documents"

	// List all files including hidden (files with the dot
	// at the beginning of the file name)
	files := ListAllFiles(target, true)

	if len(files) == 0 {
		fmt.Printf("Directory %s is empty", target)
		return
	}

	fmt.Printf("Directory %s contains:\n", target)

	for _, file := range files {
		fmt.Printf("  %s\n", file)
	}
}

func ExampleListToAbsolute() {
	dir := "/home/john/documents"

	// List all objects including hidden (files and directries
	// with the dot at the beginning of the file name)
	objects := List(dir, false)

	if len(objects) == 0 {
		fmt.Printf("Directory %s is empty", dir)
		return
	}

	// The method adds the path to the beginning of every item in the slice
	ListToAbsolute(dir, objects)
}

func ExampleListingFilter() {
	dir := "/home/john/documents"

	filter := ListingFilter{
		MatchPatterns: []string{"*.doc", "*.docx", "*.pdf"},
		MTimeOlder:    time.Now().Unix() - 3600,
		SizeGreater:   50 * 1024,
		Perms:         "FR",
	}

	docs := List(dir, false, filter)

	if len(docs) == 0 {
		fmt.Printf("No documents found in %s\n", dir)
		return
	}

	fmt.Println("Found documents:")

	for _, doc := range docs {
		fmt.Printf("  %s\n", doc)
	}
}

func ExamplePush() {
	// Current working directory is the directory where binary was executed

	cwd := Push("/home/john/documents")

	// Current working directory set to /home/john/documents

	cwd = Push("/home/john/documents/work")
	fmt.Println(cwd)

	// Current working directory set to /home/john/documents/work

	cwd = Pop()
	fmt.Println(cwd)

	// Current working directory set to /home/john/documents

	cwd = Pop()
	fmt.Println(cwd)

	// Current working directory set to initial working directory
}

func ExamplePop() {
	// Current working directory is the directory where binary was executed

	cwd := Push("/home/john/documents")

	// Current working directory set to /home/john/documents

	cwd = Push("/home/john/documents/work")
	fmt.Println(cwd)

	// Current working directory set to /home/john/documents/work

	fmt.Println(cwd)
	cwd = Pop()

	// Current working directory set to /home/john/documents

	cwd = Pop()
	fmt.Println(cwd)

	// Current working directory set to initial working directory
}
