// Package jsonutil provides methods for working with JSON data
package jsonutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"sync"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	gzipLevel   = gzip.BestSpeed
	gzipLevelMu sync.RWMutex
)

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

// SetGzipLevel sets the compression level (0-9) used by [WriteGz]
func SetGzipLevel(level int) {
	gzipLevelMu.Lock()
	gzipLevel = min(max(0, level), 9)
	gzipLevelMu.Unlock()
}

// GzipLevel returns the current compression level
func GzipLevel() int {
	gzipLevelMu.RLock()
	defer gzipLevelMu.RUnlock()
	return gzipLevel
}

// ////////////////////////////////////////////////////////////////////////////////// //

// readFile reads and decode JSON file
func readFile(file string, v any, compress bool) error {
	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	err = readData(fd, v, compress)

	if err != nil {
		return err
	}

	return fd.Close()
}

// readData reads and decode JSON data from io.ReadWriter
func readData(r io.Reader, v any, compress bool) error {
	var err error
	var dr io.Reader

	rr := bufio.NewReader(r)

	if compress {
		dr, err = gzip.NewReader(rr)

		if err != nil {
			return err
		}

		defer dr.(*gzip.Reader).Close()
	} else {
		dr = rr
	}

	return json.NewDecoder(dr).Decode(v)
}

// writeFile encodes data to JSON and save it to file
func writeFile(file string, v any, perms []os.FileMode, compressed bool) error {
	var perm os.FileMode = 0644

	if len(perms) != 0 {
		perm = perms[0]
	}

	fd, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)

	if err != nil {
		return err
	}

	err = writeData(fd, v, compressed)

	if err != nil {
		return err
	}

	return fd.Close()
}

// writeData encodes data to JSON and writes it to io.ReadWriter
func writeData(w io.Writer, v any, compressed bool) error {
	var err error
	var je *json.Encoder
	var gw *gzip.Writer

	ww := bufio.NewWriter(w)

	defer ww.Flush()

	if compressed {
		gw, err = gzip.NewWriterLevel(ww, gzipLevel)

		if err != nil {
			return err
		}

		je = json.NewEncoder(gw)

		defer gw.Close()
	} else {
		je = json.NewEncoder(ww)
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

	return ww.Flush()
}
