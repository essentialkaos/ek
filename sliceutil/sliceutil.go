// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2021 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StringToInterface converts slice with strings to slice with interface{}
func StringToInterface(data []string) []interface{} {
	var result []interface{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// IntToInterface converts slice with ints to slice with interface{}
func IntToInterface(data []int) []interface{} {
	var result []interface{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// StringToError converts slice with strings to slice with errors
func StringToError(data []string) []error {
	var result []error

	for _, e := range data {
		result = append(result, errors.New(e))
	}

	return result
}

// ErrorToString converts slice with errors to slice with strings
func ErrorToString(data []error) []string {
	var result []string

	for _, e := range data {
		result = append(result, e.Error())
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

// Exclude returns slice without items in second given slice
func Exclude(slice, items []string) []string {
	var result []string

LOOP:
	for _, i := range slice {
		for _, j := range items {
			if i == j {
				continue LOOP
			}
		}

		result = append(result, i)
	}

	return result
}

// Deduplicate removes duplicates from slice.
// Slice must be sorted before deduplication.
func Deduplicate(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}

	j := 0

	for i := 1; i < len(slice); i++ {
		if slice[j] == slice[i] {
			continue
		}

		j++

		slice[j] = slice[i]
	}

	return slice[:j+1]
}
