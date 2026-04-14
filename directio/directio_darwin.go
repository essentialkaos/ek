package directio

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"os"
	"syscall"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	BLOCK_SIZE = 4096 // BLOCK_SIZE is the minimum block size used for Direct IO reads and writes
	ALIGN_SIZE = 0    // ALIGN_SIZE is zero on macOS because F_NOCACHE does not require aligned buffers
)

// ////////////////////////////////////////////////////////////////////////////////// //

func openFile(file string, flag int, perm os.FileMode) (*os.File, error) {
	fd, err := os.OpenFile(file, flag, perm)

	if err != nil {
		return nil, err
	}

	_, _, e := syscall.Syscall(syscall.SYS_FCNTL, uintptr(fd.Fd()), syscall.F_NOCACHE, 1)

	if e != 0 {
		fd.Close()
		return nil, errors.New("can't set F_NOCACHE for given file")
	}

	return fd, nil
}
