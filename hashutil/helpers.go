// Package hashutil contains various helper functions for working with hashes
package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"hash"
	"io"
	"os"
	"strconv"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilHasher is returned if given hasher is nil
	ErrNilHasher = fmt.Errorf("Hasher is nil")

	// ErrNilSource is returned if given source is nil
	ErrNilSource = fmt.Errorf("Source is nil")

	// ErrNilDest is returned if given destination is nil
	ErrNilDest = fmt.Errorf("Destination is nil")

	// ErrNilReader is returned if given reader is nil
	ErrNilReader = fmt.Errorf("Reader is nil")

	// ErrNilWriter is returned if given writer is nil
	ErrNilWriter = fmt.Errorf("Writer is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Hash is bytes slice with hash data
type Hash []byte

// Reader is transparent hashing reader
type Reader struct {
	hasher hash.Hash
	r      io.Reader
}

// Writer is transparent hashing writer
type Writer struct {
	hasher hash.Hash
	w      io.Writer
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Copy is io.Copy like function with transparent hash calculation
func Copy(dst io.Writer, src io.Reader, hasher hash.Hash) (int64, Hash, error) {
	switch {
	case src == nil:
		return 0, nil, ErrNilSource
	case dst == nil:
		return 0, nil, ErrNilDest
	case hasher == nil:
		return 0, nil, ErrNilHasher
	}

	hasher.Reset()

	w := io.MultiWriter(dst, hasher)
	n, err := io.Copy(w, src)

	return n, Sum(hasher), err
}

// File calculates hash of the file using given hasher
func File(file string, hasher hash.Hash) Hash {
	if hasher == nil {
		return nil
	}

	fd, err := os.Open(file)

	if err != nil {
		return nil
	}

	defer fd.Close()

	hasher.Reset()
	io.Copy(hasher, fd)

	return Sum(hasher)
}

// Bytes calculates data hash using given hasher
func Bytes(data []byte, hasher hash.Hash) Hash {
	if len(data) == 0 || hasher == nil {
		return nil
	}

	hasher.Reset()
	hasher.Write(data)

	return Sum(hasher)
}

// String calculates string hash using given hasher
func String(data string, hasher hash.Hash) Hash {
	if len(data) == 0 || hasher == nil {
		return nil
	}

	hasher.Reset()
	fmt.Fprint(hasher, data)

	return Sum(hasher)
}

// Sum prints checksum
func Sum(hasher hash.Hash) Hash {
	if hasher == nil {
		return nil
	}

	return Hash(hasher.Sum(nil))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns hash as hex string
func (h Hash) String() string {
	if len(h) == 0 {
		return ""
	}

	return fmt.Sprintf("%0"+strconv.Itoa(len(h)/2)+"x", h.Bytes())
}

// Bytes returns hash as byte slice
func (h Hash) Bytes() []byte {
	return []byte(h)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader creates new reader with transparent hash calculation
func NewReader(r io.Reader, hasher hash.Hash) (*Reader, error) {
	switch {
	case r == nil:
		return nil, ErrNilReader
	case hasher == nil:
		return nil, ErrNilHasher
	}

	hasher.Reset()

	return &Reader{r: r, hasher: hasher}, nil
}

// Read reads data and simultaneously writes it into hasher
func (r *Reader) Read(p []byte) (int, error) {
	switch {
	case r == nil || r.r == nil:
		return 0, ErrNilReader
	case r.hasher == nil:
		return 0, ErrNilHasher
	}

	n, err := r.r.Read(p)

	if n > 0 {
		r.hasher.Write(p[:n])
	}

	return n, err
}

// Sum returns hash of read data
func (r *Reader) Sum() Hash {
	if r == nil || r.r == nil || r.hasher == nil {
		return nil
	}

	return Sum(r.hasher)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewWriter creates new writer with transparent hash calculation
func NewWriter(w io.Writer, hasher hash.Hash) (*Writer, error) {
	switch {
	case w == nil:
		return nil, ErrNilWriter
	case hasher == nil:
		return nil, ErrNilHasher
	}

	hasher.Reset()

	return &Writer{w: w, hasher: hasher}, nil
}

// Write simultaneously writes data to underlying writer and hasher
func (w *Writer) Write(p []byte) (int, error) {
	switch {
	case w == nil || w.w == nil:
		return 0, ErrNilWriter
	case w.hasher == nil:
		return 0, ErrNilHasher
	}

	n, err := w.w.Write(p)

	w.hasher.Write(p)

	return n, err
}

// Sum returns hash of written data
func (w *Writer) Sum() Hash {
	if w == nil || w.w == nil || w.hasher == nil {
		return nil
	}

	return Sum(w.hasher)
}

// ////////////////////////////////////////////////////////////////////////////////// //
