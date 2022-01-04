// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2022 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Copy creates copy of given slice
func Copy(slice []string) []string {
	if len(slice) == 0 {
		return nil
	}

	s := make([]string, len(slice))
	copy(s, slice)

	return s
}

// CopyInts creates copy of given slice
func CopyInts(slice []int) []int {
	if len(slice) == 0 {
		return nil
	}

	s := make([]int, len(slice))
	copy(s, slice)

	return s
}

// CopyFloats creates copy of given slice
func CopyFloats(slice []float64) []float64 {
	if len(slice) == 0 {
		return nil
	}

	s := make([]float64, len(slice))
	copy(s, slice)

	return s
}

// StringToInterface converts slice with strings to slice with interface{}
func StringToInterface(data []string) []interface{} {
	if len(data) == 0 {
		return nil
	}

	result := make([]interface{}, len(data))

	for i, r := range data {
		result[i] = r
	}

	return result
}

// IntToInterface converts slice with ints to slice with interface{}
func IntToInterface(data []int) []interface{} {
	if len(data) == 0 {
		return nil
	}

	result := make([]interface{}, len(data))

	for i, r := range data {
		result[i] = r
	}

	return result
}

// StringToError converts slice with strings to slice with errors
func StringToError(data []string) []error {
	if len(data) == 0 {
		return nil
	}

	result := make([]error, len(data))

	for i, e := range data {
		result[i] = errors.New(e)
	}

	return result
}

// ErrorToString converts slice with errors to slice with strings
func ErrorToString(data []error) []string {
	if len(data) == 0 {
		return nil
	}

	result := make([]string, len(data))

	for i, e := range data {
		result[i] = e.Error()
	}

	return result
}

// Index returns index of given item in a slice or -1 otherwise
func Index(slice []string, item string) int {
	if len(slice) == 0 {
		return -1
	}

	for i, v := range slice {
		if v == item {
			return i
		}
	}

	return -1
}

// Contains checks if string slice contains some value
func Contains(slice []string, value string) bool {
	return Index(slice, value) != -1
}

// Exclude removes items from slice
func Exclude(slice []string, items ...string) []string {
	var n int

	s := Copy(slice)

	if len(slice) == 0 || len(items) == 0 {
		return s
	}

LOOP:
	for _, i := range s {
		for _, j := range items {
			if i == j {
				continue LOOP
			}
		}

		s[n] = i
		n++
	}

	return s[:n]
}

// Deduplicate removes duplicates from slice.
// Slice must be sorted before deduplication.
func Deduplicate(slice []string) []string {
	var n int

	s := Copy(slice)

	if len(slice) <= 1 {
		return s
	}

	for i := 1; i < len(s); i++ {
		if s[n] == s[i] {
			continue
		}

		n++
		s[n] = s[i]
	}

	return s[:n+1]
}
