package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2020 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// LineCount return number of lines in file
func LineCount(file string) int {
	if file == "" {
		return -1
	}

	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return -1
	}

	defer fd.Close()

	// Use 32k buffer
	buf := make([]byte, 32*1024)
	count := 0
	sep := []byte{'\n'}

	for {
		c, err := fd.Read(buf)

		if err != nil && err != io.EOF {
			return count
		}

		count += bytes.Count(buf[:c], sep)

		if err == io.EOF {
			break
		}
	}

	return count
}
