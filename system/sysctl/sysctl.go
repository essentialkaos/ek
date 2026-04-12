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

// Param represents single kernel parameter
type Param struct {
	Name  string
	Value string
}

// Params contains all kernel parameters
type Params []Param

// ////////////////////////////////////////////////////////////////////////////////// //

// IsEmpty returns true if parameter value is empty
func (p Param) IsEmpty() bool {
	return p.Name == "" || p.Value == ""
}

// String returns parameter value as string
func (p Param) String() string {
	return p.Value
}

// Int returns parameter value as int
func (p Param) Int() (int, error) {
	return strconv.Atoi(p.Value)
}

// Int returns parameter value as int
func (p Param) Int64() (int64, error) {
	return strconv.ParseInt(p.Value, 10, 64)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Get returns kernel parameter with given name
func (p Params) Get(name string) Param {
	for _, pp := range p {
		if pp.Name == name {
			return pp
		}
	}

	return Param{}
}
