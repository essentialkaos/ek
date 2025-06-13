//go:build linux || darwin
// +build linux darwin

// Package secstr provides methods and structs for working with protected (secure) strings
package secstr

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"runtime"

	"golang.org/x/sys/unix"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// String contains protected data
type String struct {
	Data []byte
}

// ////////////////////////////////////////////////////////////////////////////////// //

// NewSecureString creates new secure string
func NewSecureString(data any) (*String, error) {
	switch v := data.(type) {
	case []byte:
		return secureStringFromSlice(v)
	case *string:
		return secureStringFromStringPointer(v)
	case string:
		return secureStringFromString(v)
	default:
		return nil, fmt.Errorf("Unsupported data type for secure string: %t", data)
	}
}

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns false if string is empty
func (s *String) IsEmpty() bool {
	return s == nil || len(s.Data) == 0
}

// Destroy destroys data
func (s *String) Destroy() error {
	if s == nil {
		return nil
	}

	return destroySecureString(s)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// secureStringFromSlice creates secure string from byte slice
func secureStringFromSlice(data []byte) (*String, error) {
	var err error

	s := &String{}

	// Destroy source data
	defer clearByteSlice(data)

	s.Data, err = unix.Mmap(
		-1, 0, len(data),
		unix.PROT_READ|unix.PROT_WRITE, // Pages may be read and written
		unix.MAP_ANON|unix.MAP_PRIVATE, // The mapping is not backed by any file + private copy-on-write
	)

	if err != nil {
		return nil, err
	}

	// Lock memory with data
	err = unix.Mlock(s.Data)

	if err != nil {
		unix.Munmap(s.Data)
		return nil, err
	}

	// Copy data
	for i := range data {
		s.Data[i] = data[i]
	}

	// Protect memory region with data
	err = unix.Mprotect(s.Data, unix.PROT_READ) // The memory can be read

	if err != nil {
		unix.Munmap(s.Data)
		clearByteSlice(s.Data) // Destroy data if memory cannot be protected
		return nil, err
	}

	// Set finalizer
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
	if s.Data == nil {
		return nil
	}

	err := unix.Mprotect(s.Data, unix.PROT_READ|unix.PROT_WRITE)

	if err != nil {
		return err
	}

	clearByteSlice(s.Data) // Clear data

	// Unlock memory
	err = unix.Munlock(s.Data)

	if err != nil {
		return err
	}

	err = unix.Munmap(s.Data)

	if err != nil {
		return err
	}

	s.Data = nil // Mark as nil for GC

	return nil
}

// clearByteSlice clears byte slice data
func clearByteSlice(s []byte) {
	for i := range s {
		s[i] = 0
	}
}
