// Package hashutil contains various helper functions for working with hashes
package hashutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"hash"
	"io"
	"os"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrNilHasher is returned when the provided hasher is nil
	ErrNilHasher = errors.New("hasher is nil")

	// ErrNilSource is returned when the provided source reader is nil
	ErrNilSource = errors.New("source is nil")

	// ErrNilDest is returned when the provided destination writer is nil
	ErrNilDest = errors.New("destination is nil")

	// ErrNilReader is returned when the provided reader is nil
	ErrNilReader = errors.New("reader is nil")

	// ErrNilWriter is returned when the provided writer is nil
	ErrNilWriter = errors.New("writer is nil")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Hash is a byte slice holding raw hash digest data
type Hash []byte

// Reader is a transparent hashing wrapper around [io.Reader].
// Every byte read is simultaneously fed into the underlying hasher.
type Reader struct {
	hasher hash.Hash
	r      io.Reader
}

// Writer is a transparent hashing wrapper around [io.Writer].
// Every byte written is simultaneously fed into the underlying hasher.
type Writer struct {
	hasher hash.Hash
	w      io.Writer
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Copy copies data from src to dst and simultaneously computes a hash,
// returning the number of bytes copied, the resulting hash, and any error.
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

// File computes the hash of the named file using the given hasher.
// Returns nil if the file cannot be opened, read, or the hasher is nil.
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

	_, err = io.Copy(hasher, fd)

	if err != nil {
		return nil
	}

	return Sum(hasher)
}

// Bytes computes the hash of the given byte slice using the provided hasher
func Bytes(data []byte, hasher hash.Hash) Hash {
	if len(data) == 0 || hasher == nil {
		return nil
	}

	hasher.Reset()
	hasher.Write(data)

	return Sum(hasher)
}

// String computes the hash of the given string using the provided hasher
func String(data string, hasher hash.Hash) Hash {
	if len(data) == 0 || hasher == nil {
		return nil
	}

	hasher.Reset()
	io.WriteString(hasher, data)

	return Sum(hasher)
}

// Sum returns the current digest of the hasher as a Hash value
func Sum(hasher hash.Hash) Hash {
	if hasher == nil {
		return nil
	}

	return Hash(hasher.Sum(nil))
}

// ////////////////////////////////////////////////////////////////////////////////// //

// String returns the hash encoded as a lowercase hexadecimal string
func (h Hash) String() string {
	if len(h) == 0 {
		return ""
	}

	return fmt.Sprintf("%x", h.Bytes())
}

// Bytes returns the raw hash digest as a plain byte slice
func (h Hash) Bytes() []byte {
	return []byte(h)
}

// Equal reports whether h and hh contain identical digest bytes
func (h Hash) Equal(hh Hash) bool {
	return bytes.Equal(h, hh)
}

// EqualString reports whether the hex string representation of h equals hh
func (h Hash) EqualString(hh string) bool {
	return h.String() == hh
}

// EqualConstantTime returns true if both hashes are equal using a
// constant-time comparison, safe for use with security-sensitive digests.
func (h Hash) EqualConstantTime(hh Hash) bool {
	return subtle.ConstantTimeCompare(h, hh) == 1
}

// IsEmpty reports whether the hash contains no data
func (h Hash) IsEmpty() bool {
	return len(h) == 0
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewReader creates a Reader that wraps r and feeds all read bytes into hasher.
// Returns an error if either argument is nil.
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

// Read reads from the underlying reader and feeds the bytes into the hasher.
// Implements [io.Reader].
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

// Sum returns the hash of all bytes read so far
func (r *Reader) Sum() Hash {
	if r == nil || r.hasher == nil {
		return nil
	}

	return Sum(r.hasher)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewWriter creates a Writer that wraps w and feeds all written bytes into hasher.
// Returns an error if either argument is nil.
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

// Write writes p to the underlying writer and simultaneously updates the hasher.
// Implements [io.Writer].
func (w *Writer) Write(p []byte) (int, error) {
	switch {
	case w == nil || w.w == nil:
		return 0, ErrNilWriter
	case w.hasher == nil:
		return 0, ErrNilHasher
	}

	n, err := w.w.Write(p)

	if n > 0 {
		w.hasher.Write(p[:n])
	}

	return n, err
}

// Sum returns the hash of all bytes written so far
func (w *Writer) Sum() Hash {
	if w == nil || w.hasher == nil {
		return nil
	}

	return Sum(w.hasher)
}

// ////////////////////////////////////////////////////////////////////////////////// //
