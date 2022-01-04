package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	BLOCK_SIZE = 4096 // Minimal block size
	ALIGN_SIZE = 4096 // Align size
)

// ////////////////////////////////////////////////////////////////////////////////// //

func openFile(file string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(file, syscall.O_DIRECT|flag, perm)
}
