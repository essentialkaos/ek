package tmp

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func ExampleNewTemp() {
	tmp, err := NewTemp()

	if err != nil {
		panic(err.Error())
	}

	tmp.DirPerms = 0700
	tmp.FilePerms = 0600
}

func ExampleTemp_MkDir() {
	tmp, err := NewTemp()

	if err != nil {
		panic(err.Error())
	}

	// Create a temporary file with an auto-generated name and suffix "myapp"
	// The suffix is optional and can be omitted.
	tmpDir, err := tmp.MkDir("myapp")

	if err != nil {
		panic(err.Error())
	}

	err = os.WriteFile(tmpDir+"/test.txt", []byte("Test data\n"), 0644)

	if err != nil {
		panic(err.Error())
	}
}

func ExampleTemp_MkFile() {
	tmp, err := NewTemp()

	if err != nil {
		panic(err.Error())
	}

	// Create a temporary file with an auto-generated name and suffix ".txt"
	// The suffix is optional and can be omitted.
	fd, filename, err := tmp.MkFile(".txt")

	if err != nil {
		panic(err.Error())
	}

	defer fd.Close()

	fmt.Fprint(fd, "Test data\n")
	fmt.Printf("Test data written to %s\n", filename)
}

func ExampleTemp_MkName() {
	tmp, err := NewTemp()

	if err != nil {
		panic(err.Error())
	}

	// Create a temporary file with an auto-generated name and suffix ".txt"
	// The suffix is optional and can be omitted.
	filename := tmp.MkName(".txt")

	if filename == "" {
		panic("Can't create name for temporary file")
	}

	err = os.WriteFile(filename, []byte("Test data\n"), 0644)

	if err != nil {
		panic(err.Error())
	}
}

func ExampleTemp_Clean() {
	tmp, err := NewTemp()

	if err != nil {
		panic(err.Error())
	}

	tmp.MkDir()

	// All temporary data (directories and files) will be removed
	tmp.Clean()
}
