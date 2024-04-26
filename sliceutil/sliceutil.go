// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Copy creates copy of given slice
func Copy[K comparable](slice []K) []K {
	if len(slice) == 0 {
		return nil
	}

	s := make([]K, len(slice))
	copy(s, slice)

	return s
}

// StringToInterface converts slice with strings to slice with any
func StringToInterface(data []string) []any {
	if len(data) == 0 {
		return nil
	}

	result := make([]any, len(data))

	for i, r := range data {
		result[i] = r
	}

	return result
}

// IntToInterface converts slice with ints to slice with any
func IntToInterface(data []int) []any {
	if len(data) == 0 {
		return nil
	}

	result := make([]any, len(data))

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
func Index[K comparable](slice []K, item K) int {
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
func Contains[K comparable](slice []K, value K) bool {
	return Index(slice, value) != -1
}

// Exclude removes items from slice
func Exclude[K comparable](slice []K, items ...K) []K {
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
func Deduplicate[K comparable](slice []K) []K {
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
