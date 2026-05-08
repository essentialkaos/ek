// Package sliceutil provides utility functions for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"slices"
	"strings"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ToAny converts a typed slice to a []any slice
func ToAny[T any](data []T) []any {
	if len(data) == 0 {
		return nil
	}

	result := make([]any, len(data))

	for i, r := range data {
		result[i] = r
	}

	return result
}

// StringToError converts a slice of strings to a slice of errors,
// wrapping each string with [errors.New]
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

// ErrorToString converts a slice of errors to a slice of strings
// by calling Error() on each element
func ErrorToString(data []error) []string {
	if len(data) == 0 {
		return nil
	}

	result := make([]string, len(data))

	for i, e := range data {
		if e != nil {
			result[i] = e.Error()
		}
	}

	return result
}

// Exclude returns a copy of slice with all occurrences of items removed
func Exclude[K comparable](slice []K, items ...K) []K {
	var n int

	s := slices.Clone(slice)

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

// Join concatenates slice elements into a single string separated by sep.
// Unlike [strings.Join], it supports slices of any type via fmt's %v verb.
func Join[T any](slice []T, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	var buf strings.Builder

	for i, v := range slice {
		fmt.Fprintf(&buf, "%v", v)

		if i+1 != len(slice) {
			buf.WriteString(sep)
		}
	}

	return buf.String()
}

// JoinFunc concatenates elems into a single string separated by sep,
// applying f to each element before joining.
func JoinFunc[T any](slice []T, sep string, f func(v T) string) string {
	if len(slice) == 0 {
		return ""
	}

	var buf strings.Builder

	for i, v := range slice {
		buf.WriteString(f(v))

		if i+1 != len(slice) {
			buf.WriteString(sep)
		}
	}

	return buf.String()
}

// Diff returns elements added to and removed from before to produce after
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

// Shuffle randomizes the order of elements in slice in place
func Shuffle[T any](slice []T) {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
}

// Filter returns a new slice containing only the elements for which
// predicate returns true, preserving the original order
func Filter[T any](slice []T, filter func(v T, index int) bool) []T {
	if len(slice) == 0 || filter == nil {
		return nil
	}

	var result []T

	for index, item := range slice {
		if filter(item, index) {
			result = append(result, item)
		}
	}

	return result
}
