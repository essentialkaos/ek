package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	BLOCK_SIZE = 0 // Minimal block size
	ALIGN_SIZE = 0 // Align size
)

// ////////////////////////////////////////////////////////////////////////////////// //

func openFile(file string, flag int, perm os.FileMode) (*os.File, error) {
	panic("UNSUPPORTED")
	return nil, nil
}
