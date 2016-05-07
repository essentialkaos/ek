// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2016 Essential Kaos                         //
//      Essential Kaos Open Source License <http://essentialkaos.com/ekol?en>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StringToInterface convert slice with strings to slice with interface{}
func StringToInterface(data []string) []interface{} {
	result := make([]interface{}, 0)

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// IntToInterface convert slice with ints to slice with interface{}
func IntToInterface(data []int) []interface{} {
	result := make([]interface{}, 0)

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// StringToError convert slice with strings to slice with errors
func StringToError(data []string) []error {
	result := make([]error, 0)

	for _, e := range data {
		result = append(result, errors.New(e))
	}

	return result
}

// StringToError convert slice with errors to slice with strings
func ErrorToString(data []error) []string {
	result := make([]string, 0)

	for _, e := range data {
		result = append(result, e.Error())
	}

	return result
}

// Contains check if string slice contains some value
func Contains(slice []string, value string) bool {
	if len(slice) == 0 {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

// Exclude return slice without items in second given slice
func Exclude(slice []string, items []string) []string {
	var result = make([]string, 0)

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

// Deduplicate return slice without duplicates. Given slice must be sorted
// before deduplication
func Deduplicate(slice []string) []string {
	sliceLen := len(slice)

	if sliceLen <= 1 {
		return slice
	}

	var result []string
	var lastItem string

	for _, v := range slice {
		if lastItem == v {
			continue
		}

		result = append(result, v)
		lastItem = v
	}

	return result
}
