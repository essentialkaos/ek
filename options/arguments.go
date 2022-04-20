package options

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"strconv"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type Arguments []string

// ////////////////////////////////////////////////////////////////////////////////// //

// Has returns true if arguments contains argument with given index
func (a Arguments) Has(index int) bool {
	return index < len(a) && a[index] != ""
}

// Get returns argument with given index
func (a Arguments) Get(index int) string {
	if index >= len(a) {
		return ""
	}

	return a[index]
}

// GetI returns argument with given index as int
func (a Arguments) GetI(index int) (int, error) {
	v := a.Get(index)
	return strconv.Atoi(v)
}

// GetF returns argument with given index as float
func (a Arguments) GetF(index int) (float64, error) {
	v := a.Get(index)
	return strconv.ParseFloat(v, 64)
}

// GetB returns argument with given index as bool
func (a Arguments) GetB(index int) (bool, error) {
	v := a.Get(index)
	switch strings.ToLower(v) {
	case "true", "yes", "y", "1":
		return true, nil
	case "false", "no", "n", "0", "":
		return false, nil
	default:
		return false, fmt.Errorf("Unsupported boolean value %q", v)
	}
}
