package fsutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// CountLines returns number of lines in given file
func CountLines(file string) (int, error) {
	if file == "" {
		return 0, ErrEmptyPath
	}

	fd, err := os.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return 0, err
	}

	// Use 32k buffer
	buf := make([]byte, 32*1024)
	count, sep := 0, []byte{'\n'}

	for {
		c, err := fd.Read(buf)

		if err != nil && err != io.EOF {
			fd.Close()
			return 0, err
		}

		count += bytes.Count(buf[:c], sep)

		if err == io.EOF {
			break
		}
	}

	return count, fd.Close()
}
