package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	BLOCK_SIZE = 4096 // BLOCK_SIZE is the minimum aligned block size required for O_DIRECT I/O
	ALIGN_SIZE = 4096 // ALIGN_SIZE is the memory alignment boundary required for O_DIRECT buffers
)

// ////////////////////////////////////////////////////////////////////////////////// //

func openFile(file string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(file, syscall.O_DIRECT|flag, perm)
}
