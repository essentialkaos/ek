// Package jsonutil provides methods for working with JSON data
package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// GzipLevel is default Gzip compression level
var GzipLevel = gzip.BestSpeed

// ////////////////////////////////////////////////////////////////////////////////// //

// Read reads and decode JSON file
func Read(file string, v any) error {
	return readFile(file, v, false)
}

// ReadGz reads and decode gzipped JSON file
func ReadGz(file string, v any) error {
	return readFile(file, v, true)
}

// Write encodes data to JSON and save it to file
func Write(file string, v any, perms ...os.FileMode) error {
	return writeFile(file, v, perms, false)
}

// Write encodes data to gzipped JSON and save it to file
func WriteGz(file string, v any, perms ...os.FileMode) error {
	return writeFile(file, v, perms, true)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readFile(file string, v any, compress bool) error {
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	defer fd.Close()

	return readData(fd, v, compress)
}

func readData(rw io.ReadWriter, v any, compress bool) error {
	var err error
	var dr io.Reader

	r := bufio.NewReader(rw)

	if compress {
		dr, err = gzip.NewReader(r)

		if err != nil {
			return err
		}
	} else {
		dr = r
	}

	return json.NewDecoder(dr).Decode(v)
}

func writeFile(file string, v any, perms []os.FileMode, compressed bool) error {
	var perm os.FileMode = 0644

	if len(perms) != 0 {
		perm = perms[0]
	}

	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)

	if err != nil {
		return err
	}

	defer fd.Close()

	return writeData(fd, v, compressed)
}

func writeData(rw io.ReadWriter, v any, compressed bool) error {
	var err error
	var je *json.Encoder
	var gw *gzip.Writer

	w := bufio.NewWriter(rw)

	if compressed {
		gw, err = gzip.NewWriterLevel(w, GzipLevel)

		if err != nil {
			return err
		}

		je = json.NewEncoder(gw)
	} else {
		je = json.NewEncoder(w)
	}

	je.SetIndent("", "  ")

	err = je.Encode(v)

	if err != nil {
		return err
	}

	if compressed {
		err = gw.Close()

		if err != nil {
			return err
		}
	}

	return w.Flush()
}
