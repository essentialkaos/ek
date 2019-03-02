// Package jsonutil provides methods for working with JSON data
package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
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
func Read(file string, v interface{}) error {
	return readFile(file, v, false)
}

// ReadGz reads and decode gzipped JSON file
func ReadGz(file string, v interface{}) error {
	return readFile(file, v, true)
}

// Write encodes data to JSON and save it to file
func Write(file string, v interface{}, perms ...os.FileMode) error {
	return writeFile(file, v, perms, false)
}

// Write encodes data to gzipped JSON and save it to file
func WriteGz(file string, v interface{}, perms ...os.FileMode) error {
	return writeFile(file, v, perms, true)
}

// EncodeToFile encodes data to JSON and save to file
func EncodeToFile(file string, v interface{}, perms ...os.FileMode) error {
	return Write(file, v, perms...)
}

// DecodeFile reads and decode JSON file
func DecodeFile(file string, v interface{}) error {
	return Read(file, v)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func readFile(file string, v interface{}, compress bool) error {
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	defer fd.Close()

	return readData(fd, v, compress)
}

func readData(rw io.ReadWriter, v interface{}, compress bool) error {
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

func writeFile(file string, v interface{}, perms []os.FileMode, compressed bool) error {
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

func writeData(rw io.ReadWriter, v interface{}, compressed bool) error {
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
