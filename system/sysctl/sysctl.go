// Package sysctl provides methods for reading kernel parameters
package sysctl

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import "strconv"

// ////////////////////////////////////////////////////////////////////////////////// //

// Param represents a single kernel parameter with its name and raw string value
type Param struct {
	Name  string
	Value string
}

// Params is a slice of kernel parameters returned by a sysctl query
type Params []Param

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if the parameter name or value is empty
func (p Param) IsEmpty() bool {
	return p.Name == "" || p.Value == ""
}

// String returns the parameter value as a string
func (p Param) String() string {
	return p.Value
}

// Int returns the parameter value parsed as an int
func (p Param) Int() (int, error) {
	return strconv.Atoi(p.Value)
}

// Int64 returns the parameter value parsed as an int64
func (p Param) Int64() (int64, error) {
	return strconv.ParseInt(p.Value, 10, 64)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns the first parameter with the given name, or an empty Param if not found
func (p Params) Get(name string) Param {
	for _, pp := range p {
		if pp.Name == name {
			return pp
		}
	}

	return Param{}
}
