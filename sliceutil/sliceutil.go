// Package sliceutil provides methods for working with slices
package sliceutil

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2019 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"errors"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// StringToInterface convert slice with strings to slice with interface{}
func StringToInterface(data []string) []interface{} {
	var result []interface{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// IntToInterface convert slice with ints to slice with interface{}
func IntToInterface(data []int) []interface{} {
	var result []interface{}

	for _, r := range data {
		result = append(result, r)
	}

	return result
}

// StringToError convert slice with strings to slice with errors
func StringToError(data []string) []error {
	var result []error

	for _, e := range data {
		result = append(result, errors.New(e))
	}

	return result
}

// ErrorToString convert slice with errors to slice with strings
func ErrorToString(data []error) []string {
	var result []string

	for _, e := range data {
		result = append(result, e.Error())
	}

	return result
}

// Index return index of given item in a slice or -1 otherwise
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

// Contains check if string slice contains some value
func Contains(slice []string, value string) bool {
	return Index(slice, value) != -1
}

// Exclude return slice without items in second given slice
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

// Deduplicate return slice without duplicates. Given slice must be sorted
// before deduplication
func Deduplicate(slice []string) []string {
	sliceLen := len(slice)

	if sliceLen <= 1 {
		return slice
	}

	var (
		result   []string
		lastItem string
	)

	for _, v := range slice {
		if lastItem == v {
			continue
		}

		result = append(result, v)
		lastItem = v
	}

	return result
}
