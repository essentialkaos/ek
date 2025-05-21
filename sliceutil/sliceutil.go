// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
)

// ////////////////////////////////////////////////////////////////////////////////// //

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
//
// Deprecated: Use method slices.Index instead
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

// Join concatenates the elements of its first argument to create a single string.
// Unlike strings.Join, this method supports slices of any type.
func Join[K any](slice []K, sep string) string {
	var buf bytes.Buffer

	for i, v := range slice {
		fmt.Fprintf(&buf, "%v", v)

		if i+1 != len(slice) {
			buf.WriteString(sep)
		}
	}

	return buf.String()
}

// Diff returns the difference (added, removed) between two slices.
// Note that slices MUST be sorted.
func Diff[K comparable](before, after []K) ([]K, []K) {
	switch {
	case len(before) == 0:
		return after, nil
	case len(after) == 0:
		return nil, before
	}

	var added, deleted []K

L1:
	for _, b := range before {
		for _, a := range after {
			if b == a {
				continue L1
			}
		}

		deleted = append(deleted, b)
	}

L2:
	for _, a := range after {
		for _, b := range before {
			if b == a {
				continue L2
			}
		}

		added = append(added, a)
	}

	return added, deleted
}

// Shuffle shuffles slice in place
func Shuffle[K comparable](slice []K) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Copy creates copy of given slice
//
// Deprecated: Use method slices.Clone instead
func Copy[K comparable](slice []K) []K {
	if len(slice) == 0 {
		return nil
	}

	s := make([]K, len(slice))
	copy(s, slice)

	return s
}

// IsEqual compares two slices and returns true if the slices are equal
//
// Deprecated: Use method slices.Equal instead
func IsEqual[K comparable](s1, s2 []K) bool {
	switch {
	case s1 == nil && s2 == nil:
		return true
	case len(s1) != len(s2):
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

// Contains checks if string slice contains some value
//
// Deprecated: Use method slices.Contains instead
func Contains[K comparable](slice []K, value K) bool {
	return Index(slice, value) != -1
}

// Deduplicate removes duplicates from slice.
// Slice must be sorted before deduplication.
//
// Deprecated: Use method slices.Compact instead
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
