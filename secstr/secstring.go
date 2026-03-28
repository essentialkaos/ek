//go:build linux || darwin

// Package secstr provides methods and structs for working with protected (secure) strings
package secstr

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"runtime"

	"github.com/essentialkaos/ek/v13/errors"

	"golang.org/x/sys/unix"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// String contains protected data backed by mlock'd, mprotect'd memory.
// Use Bytes() to access the data. Writing to the returned slice will
// cause a SIGSEGV — this is intentional and enforced by the kernel.
type String struct {
	data []byte
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewSecureString creates a new secure string from a byte slice, string, or
// string pointer. The source data is zeroed after the secure copy is made.
func NewSecureString(data any) (*String, error) {
	switch v := data.(type) {
	case []byte:
		return secureStringFromSlice(v)
	case *string:
		return secureStringFromStringPointer(v)
	case string:
		return secureStringFromString(v)
	default:
		return nil, fmt.Errorf("unsupported data type for secure string: %T", data)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if the string is nil or contains no data
func (s *String) IsEmpty() bool {
	return s == nil || len(s.data) == 0
}

// Bytes returns the underlying protected byte slice directly.
// The returned slice is read-only at the OS level (mprotect PROT_READ).
// Any write attempt will panic with a SIGSEGV — do not copy unless
// you accept that the copy loses memory protection guarantees.
func (s *String) Bytes() []byte {
	if s == nil {
		return nil
	}

	return s.data
}

// String returns protected data as string
func (s *String) String() string {
	if s == nil {
		return ""
	}

	return string(s.Bytes())
}

// Destroy zeroes and releases the protected memory region. It is safe to call multiple
// times.
func (s *String) Destroy() error {
	if s == nil {
		return nil
	}

	return destroySecureString(s)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// secureStringFromSlice allocates a protected memory region, copies data into it,
// and zeroes the source slice
func secureStringFromSlice(data []byte) (*String, error) {
	var err error

	s := &String{}

	// Destroy source data
	defer clear(data)

	s.data, err = unix.Mmap(
		-1, 0, len(data),
		unix.PROT_READ|unix.PROT_WRITE, // Pages may be read and written
		unix.MAP_ANON|unix.MAP_PRIVATE, // The mapping is not backed by any file + private copy-on-write
	)

	if err != nil {
		return nil, err
	}

	// Lock memory with data
	err = unix.Mlock(s.data)

	if err != nil {
		unix.Munmap(s.data)
		return nil, err
	}

	// Copy data
	copy(s.data, data)

	// Protect memory region with data
	err = unix.Mprotect(s.data, unix.PROT_READ) // The memory can be read

	if err != nil {
		unix.Munmap(s.data)
		clear(s.data) // Destroy data if memory cannot be protected
		return nil, err
	}

	// SetFinalizer MUST remain as the last statement after all syscalls succeed
	runtime.SetFinalizer(s, destroySecureString)

	return s, nil
}

// secureStringFromString creates secure string from string
func secureStringFromString(data string) (*String, error) {
	return secureStringFromSlice([]byte(data))
}

// secureStringFromStringPointer creates secure string from string pointer
func secureStringFromStringPointer(data *string) (*String, error) {
	s, err := secureStringFromSlice([]byte(*data))

	// Clear source data
	*data = ""

	return s, err
}

// ////////////////////////////////////////////////////////////////////////////////// //

// destroySecureString destroys secure string data
func destroySecureString(s *String) error {
	if s.data == nil {
		return nil
	}

	var errs errors.Errors

	err := unix.Mprotect(s.data, unix.PROT_READ|unix.PROT_WRITE)

	if err != nil {
		errs = append(errs, fmt.Errorf("mprotect: %w", err))
	}

	clear(s.data) // Clear data

	// Unlock memory
	err = unix.Munlock(s.data)

	if err != nil {
		errs = append(errs, fmt.Errorf("munlock: %w", err))
	}

	err = unix.Munmap(s.data)

	if err != nil {
		errs = append(errs, fmt.Errorf("munmap: %w", err))
	}

	s.data = nil // Mark as nil for GC

	return errs.Join()
}
