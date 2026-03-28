//go:build !linux && !darwin

// Package secstr provides methods and structs for working with protected (secure) strings
package secstr

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ String contains protected data backed by mlock'd, mprotect'd memory.
// Use Bytes() to access the data. Writing to the returned slice will
// cause a SIGSEGV — this is intentional and enforced by the kernel.
type String struct{}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ NewSecureString creates a new secure string from a byte slice, string, or
// string pointer. The source data is zeroed after the secure copy is made.
func NewSecureString(data any) (*String, error) {
	panic("UNSUPPORTED")
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ❗ IsEmpty returns true if the string is nil or contains no data
func (s *String) IsEmpty() bool {
	panic("UNSUPPORTED")
}

// ❗ Destroy zeroes and releases the protected memory region. It is safe to call
// multiple times.
func (s *String) Destroy() error {
	panic("UNSUPPORTED")
}

// ❗ Bytes returns the underlying protected byte slice directly.
// The returned slice is read-only at the OS level (mprotect PROT_READ).
// Any write attempt will panic with a SIGSEGV — do not copy unless
// you accept that the copy loses memory protection guarantees.
func (s *String) Bytes() []byte {
	panic("UNSUPPORTED")
}
